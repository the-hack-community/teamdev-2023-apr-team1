import { Alert } from 'react-native';
import { Marker, MarkerPressEvent } from 'react-native-maps';

import { IStrayCat } from '../../../types';

type Props = {
  m: IStrayCat;
  handleOnPress: (e: MarkerPressEvent) => void;
};

export default function MapMarker(props: Props) {
  const { m, handleOnPress } = props;

  return (
    <>
      <Marker
        title={m?.name}
        pinColor='red'
        coordinate={{
          latitude: Number(m?.location?.lat),
          longitude: Number(m?.location?.long)
        }}
        description={m?.features}
        onPress={handleOnPress}
        // onPress={(e: MarkerPressEvent) => {
        // e.stopPropagation();
        // Alert.alert(`特徴:${m?.condition}\n 発見日時: ${m?.captureDateTime}`);
        // }}
      />
    </>
  );
}
