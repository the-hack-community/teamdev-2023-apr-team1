import { useIsFocused } from '@react-navigation/native';
import { Button } from 'native-base';
import { useState, useEffect } from 'react';
import { Alert, StyleSheet, View, TextInput } from 'react-native';
import { Dimensions } from 'react-native';
import MapView, {
  MapPressEvent,
  MarkerPressEvent,
  PROVIDER_GOOGLE
} from 'react-native-maps';

import { BASE_URL, GOOGLE_MAPS_API_KEY } from '../../../../environment';
import { IStrayCat } from '../../../types';

import DetailDialog from './detailDialog';
import MapMarker from './marker';

//init position: Tokyo
const INIT_LAT = 35.681236;
const INIT_LNG = 139.767125;

const { width, height } = Dimensions.get('window');
const ASPECT_RATIO = width / height;
const LATITUDE_DELTA = 0.0922;
const LONGITUDE_DELTA = LATITUDE_DELTA * ASPECT_RATIO;
const INITIAL_POSITION = {
  latitude: INIT_LAT,
  longitude: INIT_LNG,
  latitudeDelta: LATITUDE_DELTA,
  longitudeDelta: LONGITUDE_DELTA
};

const initRegion = {
  latitude: INIT_LAT,
  longitude: INIT_LNG,
  latitudeDelta: LATITUDE_DELTA,
  longitudeDelta: LONGITUDE_DELTA
};

interface IGeolocation {
  latitude: number;
  longitude: number;
}

interface IRegion {
  latitude: number;
  longitude: number;
  latitudeDelta: number;
  longitudeDelta: number;
}

export default function Map() {
  const [markers, setMarkers] = useState<IStrayCat[]>([]);
  const [region, setRegion] = useState<IRegion>(initRegion);
  const [address, setAddress] = useState<string>('');

  const handleLocationSearch = async () => {
    if (!address) return;

    try {
      const res: any = await fetch(
        `https://maps.googleapis.com/maps/api/geocode/json?address=${address}&key=${GOOGLE_MAPS_API_KEY}`
      );
      const data = await res.json();
      if (data.status === 'OK') {
        const location = data.results[0].geometry.location;

        setRegion({
          latitude: location.lat,
          longitude: location.lng,
          latitudeDelta: LATITUDE_DELTA,
          longitudeDelta: LONGITUDE_DELTA
        });
      } else {
        console.log('Geocoding API request failed.');
      }
    } catch (e) {
      console.error(e);
    }
  };

  useEffect(() => {
    //ステネコのデータを取得
    const fetchStrayCatsData = async () => {
      const url = `${BASE_URL}/stray-cats/search`;
      try {
        const res = await fetch(url);
        const data = await res.json();
        setMarkers(data);
      } catch (e) {
        console.error(e);
      }
    };

    fetchStrayCatsData();
  }, []);

  const [selectedMarker, setSelectedMarker] = useState<IStrayCat | null>(null);
  const [isShowDialog, setIsShowDialog] = useState<boolean>(false);

  const handleOnPressMarker = (e: MarkerPressEvent, m: IStrayCat) => {
    e.stopPropagation();

    setIsShowDialog(true);
    setSelectedMarker(m);
  };

  return (
    <>
      {region && <></>}
      <View style={styles.searchContainer}>
        <TextInput
          placeholder='住所を入力してください'
          value={address}
          style={{
            borderColor: 'gray',
            borderWidth: 1,
            borderRadius: 4,
            width: '100%',
            padding: 8
          }}
          // w='100%'
          onChangeText={(text) => setAddress(text)}
          onSubmitEditing={handleLocationSearch}
        />
        {/* <Button onPress={handleLocationSearch}>検索</Button> */}
      </View>
      <MapView
        style={styles.map}
        // initialRegion={INITIAL_POSITION}
        region={region}
        provider={PROVIDER_GOOGLE}
        onRegionChange={(region) => {
          console.log(region);
        }}
        onRegionChangeComplete={(region) => {
          console.log(region);
        }}
        onPress={(e: MapPressEvent) => {
          e.stopPropagation();
          // const marker = {
          //   latitude: e.nativeEvent.coordinate.latitude,
          //   longitude: e.nativeEvent.coordinate.longitude
          // };
        }}>
        {markers &&
          markers.map((m, index) => {
            return (
              <MapMarker
                m={m}
                key={index}
                handleOnPress={(e) => handleOnPressMarker(e, m)}
              />
            );
          })}
      </MapView>
      <View>
        <DetailDialog
          m={selectedMarker}
          isShow={isShowDialog}
          setIsShow={setIsShowDialog}
        />
      </View>
    </>
  );
}

const styles = StyleSheet.create({
  map: {
    // ...StyleSheet.absoluteFillObject,
    width: Dimensions.get('window').width,
    height: Dimensions.get('window').height
  },
  searchContainer: {
    position: 'absolute',
    top: 10,
    backgroundColor: 'white',
    width: 320,
    flex: 1,
    flexDirection: 'row',
    justifyContent: 'space-between',
    padding: 8,
    borderRadius: 8,
    left: 40,
    zIndex: 1
  }
});
