import { NativeBaseProvider, ScrollView, FlatList } from 'native-base';
import React, { useEffect, useState } from 'react';
import { Dimensions, Image, View } from 'react-native';

type jsonData = {
  id: string;
  url: string;
  width: string;
  height: string;
};

// 画像表示のデモ
function PostsScreen() {
  const [catImages, setCatImages] = useState<jsonData[]>([]);
  const [imageWidth, setImageWidth] = useState<number>(0);
  const [imageHeight, setImageHeight] = useState<number>(0);
  const baseURL = 'https://api.thecatapi.com/v1/images/search';

  const fetchData = async () => {
    // 10件のデータを取得
    const json = await fetch(`${baseURL}?limit=10`);
    const data = await json.json();
    if (data) {
      setCatImages(data);
    }
  };

  useEffect(() => {
    fetchData();
    // スクリーンサイズの取得
    const { width, height } = Dimensions.get('screen');
    // 画像サイズの設定
    setImageWidth(width / 2);
    setImageHeight((width / 2) * 1.5);
  }, []);

  return (
    <NativeBaseProvider>
        <View style={{ flex: 1, alignItems: 'center' }}>
          <FlatList
            data={catImages}
            renderItem={({ item, index }) => {
              return (
                <View style={{ width: imageWidth }}>
                  <Image
                    style={{ width: '100%', height: imageHeight }}
                    resizeMode='cover'
                    source={{ uri: item.url }}
                  />
                </View>
              );
            }}
            numColumns={2}
            contentContainerStyle={{ flexGrow: 1, justifyContent: 'center' }}
          />
        </View>
    </NativeBaseProvider>
  );
}

export default PostsScreen;