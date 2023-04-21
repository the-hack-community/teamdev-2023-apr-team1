import { NativeBaseProvider, ScrollView } from 'native-base';
import React, { useEffect, useState } from 'react';
import { Image, View } from 'react-native';

type jsonData = {
  id: string;
  url: string;
  width: string;
  height: string;
};

// 画像表示のデモ
function PostsScreen() {
  const [catImages, setCatImages] = useState<jsonData[]>([]);
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
  }, []);

  return (
    <NativeBaseProvider>
      <ScrollView>
        {catImages &&
          catImages.map((item: jsonData, index: number) => {
            return (
              <View key={index}>
                <Image
                  style={{ width: 400, height: 400 }}
                  source={{ uri: item.url }}
                />
              </View>
            );
          })}
      </ScrollView>
    </NativeBaseProvider>
  );
}

export default PostsScreen;
