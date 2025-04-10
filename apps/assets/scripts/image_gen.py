from openai import OpenAI
import os
import time
import base64
from dotenv import load_dotenv


load_dotenv()
client = OpenAI()
client.api_key = os.getenv("OPENAI_API_KEY")
f = open('prompt.txt', 'r')
prompt = f.read()
f.close()
response = client.images.generate(
  model="dall-e-3", 
  prompt=prompt, 
  n=1,  # 生成数
  size="1792x1024",  # 解像度 dall-e-3では1024x1024、1792x1024、1024x1792
  response_format="b64_json",  # レスポンスフォーマット url or b64_json
  quality="hd",  # 品質 standard or hd
  style="vivid"  # スタイル vivid or natural
)

for i, d in enumerate(response.data):
    with open(f"images/dall-e-3_{int(time.time())}_{i}.png", "wb") as f:
        f.write(base64.b64decode(d.b64_json))