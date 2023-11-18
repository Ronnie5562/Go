import requests

response = requests.get("https://prettyprinted.com")
# response = requests.get("https://thesuperkidsschool.org/")

with open("schools.html", "w", encoding='utf-8') as f:
    f.write(response.text)

if response:
    print("Response OK")
else:
    print("Response FAILED")

