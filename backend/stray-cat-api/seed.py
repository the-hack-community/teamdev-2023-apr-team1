import os
import requests
from datetime import datetime
import random
from io import BytesIO

# ノラネコの名前、特徴、状態のリスト
names = ['タマ', 'ミケ', 'シロ', 'チビ', 'クロ', 'ハチ', 'サバ', 'ポチ', 'トラ', 'ベン']
features = ['黒い毛', '白い毛', '三毛', '長毛', '短毛', '大きな尾', '小さな尾', '黄色い目', '青い目', '緑色の目']
conditions = ['健康', '少し痩せている', 'けがをしている', '病気', '元気がない', '活発', '寝ている', '食事をしている', '遊んでいる', '恐怖']


SERVER_NAME = os.environ.get('SERVER_NAME', 'localhost')  # If 'SERVER_NAME' is not set, it defaults to 'localhost'

def post_cat_info():
    # 猫画像を取得
    response = requests.get('https://api.thecatapi.com/v1/images/search')
    data = response.json()

    # 画像をバイナリ形式で取得
    img_data = requests.get(data[0]['url']).content

    # 日本のランダムな座標を生成
    # lat = random.uniform(24.0, 46.0)  # 緯度
    # long = random.uniform(123.0, 154.0)  # 経度

    # 東京駅周辺にランダムな座標を生成
    lat = random.uniform(35.660000, 35.690000)  # 緯度
    long = random.uniform(139.735000, 139.760000)  # 経度

    # データを準備
    post_data = {
        'name': random.choice(names),  # ランダムな名前
        'features': random.choice(features),  # ランダムな特徴
        'condition': random.choice(conditions),  # ランダムな状態
        'captureDateTime': datetime.now().isoformat(),  # 撮影日時
        'lat': lat,  # 緯度
        'long': long,  # 経度
    }

    files = {
        'photo': ('photo.jpg', BytesIO(img_data), 'image/jpeg'),
    }

    # APIサーバにデータを送信
    response = requests.post(f'http://{SERVER_NAME}:8080/stray-cats', data=post_data, files=files)

    return response.status_code

# 10回関数を実行
for i in range(10):
    status_code = post_cat_info()
    print(f'Attempt {i+1}: Status code {status_code}')
