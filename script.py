import requests

response = requests.get("https://prettyprinted.com")

with open("indeex.html", "w") as f:
    f.write(response.text)

if response:
    print("Response OK")
else:
    print("Response FAILED")

