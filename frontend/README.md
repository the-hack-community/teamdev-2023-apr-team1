# Frontend
React Nativeを使用した地図アプリケーションです。

## 使用技術
- https://reactnative.dev/
- https://expo.dev/

## セットアップ
- `$ npm install -g expo-cli`
- `$ git clone`
- `$ cd frontend`
- `$ npm install`

- `environment_sample.ts`を`environment.ts`にリネームしてください。
- `app.sample.json`を`app.json`にリネームしてください。
- GoogleMapApiKeyが必要です。
`app.json`と`environment.ts`にGoogleMapApiKeyを設定してください。
- Firebaseコンソールから`google-services.json`をダウンロードし、frontendのルートフォルダに配置してください。

## ローカルでのサーバー起動方法

`$ npm run start`

a(Android)またはi(iOS)を押すと、シミュレーターが起動します。
