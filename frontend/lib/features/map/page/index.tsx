import { NativeBaseProvider } from 'native-base';
import React from 'react';
import { View } from 'react-native';

import Map from '../parts/map';

function MapScreen() {
  return (
    <NativeBaseProvider>
      <View>
        <Map />
      </View>
    </NativeBaseProvider>
  );
}

export default MapScreen;
