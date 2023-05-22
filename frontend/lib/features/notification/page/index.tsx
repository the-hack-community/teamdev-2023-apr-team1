import {
  useFonts,
  NotoSansJP_400Regular
} from '@expo-google-fonts/noto-sans-jp';
import messaging from '@react-native-firebase/messaging';
import { Text, Box, NativeBaseProvider } from 'native-base';
import React, { useEffect } from 'react';
import { Alert, View } from 'react-native';

function NotificationScreen() {
  const [fontsLoaded] = useFonts({
    NotoSansJP_400Regular
  });

  useEffect(() => {
    const unsubscribe = messaging().onMessage(async (remoteMessage) => {
      //受け取ったメッセージを表示
      Alert.alert(
        '通知メッセージを受け取りました!',
        JSON.stringify(remoteMessage)
      );
    });

    return unsubscribe;
  }, []);

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
        <View>
          <Text>通知テスト</Text>
          <Box py='4' width='250'></Box>
        </View>
      </View>
    </NativeBaseProvider>
  );
}

export default NotificationScreen;
