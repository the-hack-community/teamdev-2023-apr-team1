import { NativeBaseProvider, ScrollView } from 'native-base';
import React, { useEffect, useState, useCallback, useMemo } from 'react';
import { Image, View, Button } from 'react-native'; // Buttonを追加
import ImagePicker from 'react-native-image-picker';

// 画像データの型定義
type CatImageData = {
  id: string;
  url: string;
  width: number;
  height: number;
};

function PostsScreen() {
  const [catImages, setCatImages] = useState<CatImageData[]>([]);
  const baseURL = 'https://api.thecatapi.com/v1/images/search';

  // APIから画像データを取得する関数
  const fetchData = useCallback(async () => {
    const json = await fetch(`${baseURL}?limit=10`);
    const data = await json.json();
    if (data) {
      setCatImages(data);
    }
  }, []);

  // APIから画像データを取得
  useEffect(() => {
    fetchData();
  }, [fetchData]);

  // 画像選択機能
  const handleImagePicker = useCallback(() => {
    const options = {
      title: 'Select Avatar',
      storageOptions: {
        skipBackup: true,
        path: 'images',
      },
    };

    // 画像選択ダイアログを表示
    ImagePicker.showImagePicker(options, response => {
      if (response.didCancel) {
        console.log('User cancelled image picker');
      } else if (response.error) {
        console.log('ImagePicker Error: ', response.error);
      } else {
        const source = { uri: response.uri };
        const resizedSource = {
          uri: response.uri,
          width: response.width,
          height: response.height,
        };
        setCatImages(prevState => [ ...prevState, resizedSource ]);
      }
    });
  }, []);

  // 画像データを表示する
  const catImageData = useMemo(() => catImages.map((item, index) => (
    <View key={index}>
      <Image
        style={{ width: item.width, height: item.height }} 
        source={{ uri: item.url }}
        alt="cat image"
      />
    </View>
  )), [catImages]);

  return (
    <NativeBaseProvider>
      <ScrollView>
        {catImageData}
        <View>
          <Button onPress={handleImagePicker} title="Pick an image" />
        </View>
      </ScrollView>
    </NativeBaseProvider>
  );
}

export default PostsScreen;