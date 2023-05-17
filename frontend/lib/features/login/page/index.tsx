import {
  useFonts,
  NotoSansJP_400Regular
} from '@expo-google-fonts/noto-sans-jp';
import {
  GoogleSignin,
  GoogleSigninButton,
  User,
  statusCodes
} from '@react-native-google-signin/google-signin';
import { Button, Text, Box, NativeBaseProvider } from 'native-base';
import React, { useState } from 'react';
import { View } from 'react-native';

import { WEB_CLIENT_ID } from '../../../../environment';

GoogleSignin.configure({
  webClientId: WEB_CLIENT_ID
});

function LoginScreen() {
  const [fontsLoaded] = useFonts({
    NotoSansJP_400Regular
  });
  const [userInfo, setUserInfo] = useState<User | null>();
  const [isSigninInProgress, setIsSigninInProgress] = useState<boolean>(false);

  //Googleログイン
  const onGoogleLogin = async () => {
    setIsSigninInProgress(true);
    try {
      await GoogleSignin.hasPlayServices({
        showPlayServicesUpdateDialog: true
      });
      const userInfo = await GoogleSignin.signIn();
      setUserInfo(userInfo);
    } catch (error: any) {
      if (error.code === statusCodes.SIGN_IN_CANCELLED) {
        // console.log('SIGN_IN_CANCELLED');
      } else if (error.code === statusCodes.IN_PROGRESS) {
        // console.log('IN_PROGRESS');
      } else if (error.code === statusCodes.PLAY_SERVICES_NOT_AVAILABLE) {
        // console.log('PLAY_SERVICES_NOT_AVAILABLE');
      } else {
        console.log(error);
      }
    }
    setIsSigninInProgress(false);
  };

  //サインアウト
  const onSignOut = async () => {
    try {
      await GoogleSignin.signOut();
      console.log('signOut');
      setUserInfo(null);
    } catch (error: any) {
      console.error(error);
    }
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
