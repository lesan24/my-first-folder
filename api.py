import requests

print("--- GITHUB SAYS ---")
print(requests.get("https://api.github.com/zen").text)

print("\n--- YOUR INTERNET LOCATION ---")
data = requests.get("https://ipinfo.io/json").json()
print(f"City: {data.get('city')}, {data.get('region')}")
print(f"IP: {data.get('ip')}")
print(f"Org: {data.get('org')}")
