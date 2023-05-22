import {
  useFonts,
  NotoSansJP_400Regular
} from '@expo-google-fonts/noto-sans-jp';
import messaging from '@react-native-firebase/messaging';
import { Text, Box, NativeBaseProvider, Center } from 'native-base';
import React, { useEffect } from 'react';
import { Alert, View } from 'react-native';

function NotificationScreen() {
  const [fontsLoaded] = useFonts({
    NotoSansJP_400Regular
  });

  useEffect(() => {
    const unsubscribe = messaging().onMessage(async (remoteMessage) => {
      const { title, body } = remoteMessage?.notification || {};
      //受け取ったメッセージを表示
      Alert.alert(
        '通知メッセージを受け取りました!',
        `タイトル: ${title}\n本文: ${body}`
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
          <Center>
            <Text>通知テスト</Text>
            <Text>メッセージが届くまで待ちましょう！</Text>
          </Center>
        </View>
      </View>
    </NativeBaseProvider>
  );
}

export default NotificationScreen;
