import moment from 'moment';
import { Center } from 'native-base';
import React from 'react';
import { Image, Text, View } from 'react-native';
import Dialog from 'react-native-dialog';

import { IStrayCat } from '../../../types';

type Props = {
  m: IStrayCat | null;
  isShow: boolean;
  setIsShow: (isShow: boolean) => void;
};

export default function DetailDialog(props: Props) {
  const { m, isShow, setIsShow } = props;

  const date = moment(m?.captureDateTime).format('YYYY年MM月DD日 HH時mm分');

  return (
    <Dialog.Container visible={isShow}>
      <Center>
        <View>
          <Image
            style={{ width: 250, height: 250 }}
            source={{ uri: `data:image/jpeg;base64,${m?.photoData}` }}
            resizeMode='cover'
          />
        </View>
      </Center>
      <View style={{ marginTop: 8 }}>
        <Text>名前: {m?.name}</Text>
        <Text>状態: {m?.condition}</Text>
        <Text>特徴: {m?.features}</Text>
        <Text>日時: {date}</Text>
      </View>
      <Dialog.Button label='閉じる' onPress={() => setIsShow(false)} />
    </Dialog.Container>
  );
}
