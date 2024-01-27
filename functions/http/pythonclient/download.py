#import os,requests
#def download(url):
#    get_response = requests.get(url,stream=True)
#    file_name  = url.split("/")[-1]
#    with open(file_name, 'wb') as f:
#        for chunk in get_response.iter_content(chunk_size=1024):
#            if chunk: # filter out keep-alive new chunks
#                f.write(chunk)


import requests

URL = "https://td-cdn.pw/api.php?download=tikdown.org-42500282235.mp4"
FILE_TO_SAVE_AS = "myvideo.mp4" # the name you want to save file as


resp = requests.get(URL) # making requests to server

with open(FILE_TO_SAVE_AS, "wb") as f: # opening a file handler to create new file 
    f.write(resp.content) # writing content to file