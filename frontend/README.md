# Frontend
ステネコスタンプ、フロントエンドサイドの実装です。

## 使用技術
- https://reactnative.dev/
- https://expo.dev/

## セットアップ
- `$ npm install -g expo-cli`
- `$ git clone https://github.com/the-hack-community/teamdev-2023-apr-team1.git`
- `$ cd frontend`
- `$ npm install`

- `environment_sample.ts`を`environment.ts`にリネームしてください。
- `app.sample.json`を`app.json`にリネームしてください。
- GoogleMapApiKeyが必要です。
`app.json`と`environment.ts`にGoogleMapApiKeyを設定してください。
- Firebaseコンソールから`google-services.json`をダウンロードし、frontendのルートフォルダに配置してください。 

## ローカルでのサーバー起動方法

`$ npm run start`
or
`$ npx expo start`

コマンド実行後、インタラクティブモードで、`a`(Android)または`i`(iOS)を押すと、シミュレーターが起動します。

## React Native Firebaseライブラリをプロジェクトにリンク、再構築を行う場合

以下のコマンドを実行する。

`$ npx react-native run-android`

## ネイティブコードへの反映

以下のコマンドを実行する。

`$ npx expo run:android`

## APKファイルの生成

以下のコマンドを実行する。

`cd android`
`./gradlew assembleDebug`

実行後、`android/app/build/outputs/apk/debug/app-debug.apk`ファイルが生成される。

ストアにリリースする場合は、以下コマンドを実行する。
`./gradlew bundleRelease`

詳しくは、
https://reactnative.dev/docs/signed-apk-android