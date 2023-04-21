import { Button, Text, Box, NativeBaseProvider } from 'native-base';
import React from 'react';
import { View, Alert } from 'react-native';

function LoginScreen() {
  const handleGoogleLogin = () => {
    Alert.alert('Googleアカウントでログイン');
  };

  const handleFacebookLogin = () => {
    Alert.alert('Facebookアカウントでログイン');
  };

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
            <Text fontSize='sm' color='black'>
              Googleアカウントでログイン
            </Text>
          </Button>
        </Box>
        <Box py='4' width='250'>
          <Button
            alignItems='center'
            onPress={handleFacebookLogin}
            style={{ backgroundColor: '#385490', width: '100%' }}>
            <Text fontSize='sm' color='white'>
              Facebookアカウントでログイン
            </Text>
          </Button>
        </Box>
      </View>
    </NativeBaseProvider>
  );
}

export default LoginScreen;
