import {
  useFonts,
  NotoSansJP_400Regular
} from '@expo-google-fonts/noto-sans-jp';
import { GoogleSigninButton } from '@react-native-google-signin/google-signin';
import { Button, Text, Box, NativeBaseProvider } from 'native-base';
import React from 'react';
import { View } from 'react-native';

import { useAuth } from '../parts/authContext';

function LoginScreen() {
  const [fontsLoaded] = useFonts({
    NotoSansJP_400Regular
  });

  const { userInfo, isSigninInProgress, onGoogleLogin, onSignOut } = useAuth();

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
        {!userInfo ? (
          <Box py='4' width='250'>
            <GoogleSigninButton
              disabled={isSigninInProgress}
              onPress={() => onGoogleLogin()}></GoogleSigninButton>
          </Box>
        ) : (
          <View>
            <Text> ログインユーザのメールアドレス:</Text>
            <Text>{userInfo && userInfo?.user?.email}</Text>
            <Box py='4' width='250'>
              <Button
                onPress={() => onSignOut()}
                style={{ backgroundColor: '#fff' }}>
                <Text>サインアウト</Text>
              </Button>
            </Box>
          </View>
        )}
      </View>
    </NativeBaseProvider>
  );
}

export default LoginScreen;
