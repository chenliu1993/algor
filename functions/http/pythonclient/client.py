import requests
import json
import re

if __name__ == "__main__":
    jsonfile = open("request.json", "r")
    data = json.load(jsonfile)
    headers = {'Name': 'Jliu'}
    resp = requests.post("http://127.0.0.1:8090/hello", headers=headers, data=data)
    print(re.search(r'Hi [a-zA-Z]+', resp.text).group(0))