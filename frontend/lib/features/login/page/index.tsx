import {
  useFonts,
  NotoSansJP_400Regular
} from '@expo-google-fonts/noto-sans-jp';
import { Button, Text, Box, NativeBaseProvider } from 'native-base';
import React from 'react';
import { View, Alert } from 'react-native';

function LoginScreen() {
  const [fontsLoaded] = useFonts({
    NotoSansJP_400Regular
  });

  const handleGoogleLogin = () => {
    Alert.alert('Googleアカウントでログイン');
  };

  const handleFacebookLogin = () => {
    Alert.alert('Facebookアカウントでログイン');
  };

  if (!fontsLoaded) {
    return <View></View>;
  }

  return (
    <NativeBaseProvider>
      <View
        style={{
          flex: 1,
          justifyContent: 'center',
          alignItems: 'center',
          backgroundColor: '#eeeeee'
        }}>
        <Box py='4' width='250'>
          <Button
            alignItems='center'
            onPress={handleGoogleLogin}
            style={{ backgroundColor: 'white', width: '100%' }}>
            <Text
              fontSize='sm'
              color='black'
              fontFamily='NotoSansJP_400Regular'>
              Googleアカウントでログイン
            </Text>
          </Button>
        </Box>
        <Box py='4' width='250'>
          <Button
            alignItems='center'
            onPress={handleFacebookLogin}
            style={{ backgroundColor: '#385490', width: '100%' }}>
            <Text
              fontSize='sm'
              color='white'
              fontFamily='NotoSansJP_400Regular'>
              Facebookアカウントでログイン
            </Text>
          </Button>
        </Box>
      </View>
    </NativeBaseProvider>
  );
}

export default LoginScreen;
