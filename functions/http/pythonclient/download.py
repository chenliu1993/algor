import requests


url = 'https://www.facebook.com/favicon.ico'
r = requests.get(url, allow_redirects=True)

open('facebook.ico', 'wb').write(r.content)

# >>> with open("WDI_CSV.zip", mode="wb") as file:
# ...     for chunk in response.iter_content(chunk_size=10 * 1024):
# ...         file.write(chunk)
# ...