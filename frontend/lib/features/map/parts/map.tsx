import { Button, Input } from 'native-base';
import { useState, useEffect } from 'react';
import { Alert, StyleSheet, View } from 'react-native';
import { Dimensions } from 'react-native';
import MapView, {
  MapPressEvent,
  MarkerPressEvent,
  PROVIDER_GOOGLE
} from 'react-native-maps';

import { BASE_URL, GOOGLE_MAPS_API_KEY } from '../../../../environment';
import { IStrayCat } from '../../../types';

import Marker from './marker';
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

const markerList: IStrayCat[] = [];

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
      const location = data.results[0].geometry.location;
      setRegion({
        latitude: location.lat,
        longitude: location.lng,
        latitudeDelta: LATITUDE_DELTA,
        longitudeDelta: LONGITUDE_DELTA
      });
    } catch (e) {
      console.error(e);
    }
  };

  useEffect(() => {
    //ステネコのデータを取得
    const fetchStrayCatsData = async () => {
      const url = `${BASE_URL}/stray-cats/search`;
      const res = await fetch(url);
      const data = await res.json();
      setMarkers(data);
    };

    fetchStrayCatsData();
  }, []);

  return (
    <>
      <View style={styles.searchContainer}>
        <Input
          placeholder='住所を入力してください'
          value={address}
          w='80%'
          onChangeText={(text) => setAddress(text)}
        />
        <Button onPress={handleLocationSearch}>検索</Button>
      </View>
      <MapView
        style={styles.map}
        initialRegion={INITIAL_POSITION}
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
          const marker = {
            latitude: e.nativeEvent.coordinate.latitude,
            longitude: e.nativeEvent.coordinate.longitude
          };
        }}>
        {markers &&
          markers.map((m, index) => {
            return <MapMarker m={m} key={index} />;
          })}
      </MapView>
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
