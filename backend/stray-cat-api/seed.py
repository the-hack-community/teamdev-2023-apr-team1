import requests
from datetime import datetime
import base64
import random

# 猫画像を取得
response = requests.get('https://api.thecatapi.com/v1/images/search')
data = response.json()

# 画像をバイナリ形式で取得
img_data = requests.get(data[0]['url']).content
# バイナリデータをbase64形式に変換
img_b64 = base64.b64encode(img_data).decode('utf-8')

# 日本のランダムな座標を生成
lat = random.uniform(24.0, 46.0)  # 緯度
long = random.uniform(123.0, 154.0)  # 経度

# データを準備
post_data = {
    'name': 'ノラネコ',  # 名前
    'features': '特徴',  # 特徴
    'condition': '状態',  # 状態
    'captureDateTime': datetime.now().isoformat(),  # 撮影日時
    'lat': lat,  # 緯度
    'long': long,  # 経度
    'photo': img_b64,  # 写真
}

# APIサーバにデータを送信
response = requests.post('http://localhost:8080/stray-cats', data=post_data)

print(response.status_code)
