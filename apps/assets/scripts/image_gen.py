from openai import OpenAI
import os
import time
import base64
from dotenv import load_dotenv


load_dotenv()
client = OpenAI()
client.api_key = os.getenv("OPENAI_API_KEY")  # API keyのセット
response = client.images.generate(
  model="dall-e-3",  # モデル
  prompt="""
  私はwebアプリの開発者です。私は、あなたに私のアプリのためにホームページのトップに置くアプリを理解できるような画像を生成してほしいです。以下の要件に従ってください。
  - アプリの名前: odachin
  - アプリの説明: odachinは、子供のお小遣いを親と一緒に管理し、月のお小遣いやお手伝いによるお小遣いをポイントとして親から子供に渡せて、使い道を話し合えるようなアプリです。
  - アプリのテーマ: 子供の発育を促すためのアプリ。
  - アプリの雰囲気: 明るい雰囲気。
  - 画像のスタイル: おちついた雰囲気のデザイン。
  子供が親の家事の手伝いをしているイラストやお小遣いをもらっているイラストなど、ユースケースを複数含めてください。
  
  """
  ,
  n=1,  # 生成数
  size="1792x1024",  # 解像度 dall-e-3では1024x1024、1792x1024、1024x1792
  response_format="b64_json",  # レスポンスフォーマット url or b64_json
  quality="hd",  # 品質 standard or hd
  style="vivid"  # スタイル vivid or natural
)

# 画像保存
# ファイル名にはタイムスタンプと通番を含めています
for i, d in enumerate(response.data):
    with open(f"images/dall-e-3_{int(time.time())}_{i}.png", "wb") as f:
        f.write(base64.b64decode(d.b64_json))