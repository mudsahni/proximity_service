import json
import random
import requests
import h3

base_payload = {
    "name": "The big chill",
    "city": "New Delhi",
    "country": "India",
    "type": "Restaurant",
    "latitude": 77.2088,
    "longitude": 28.6139,
}
class LatLng(object):
    def __init__(self, lat, lng):
        self.lat = lat
        self.lng = lng

places = ["The Big Chill", "M Block Market, GK 1", "Juggernaut", "ISKON Temple", "Moolchand Hospital"]
cities = ["New Delhi", "New Delhi", "New Delhi", "New Delhi", "New Delhi"]
countries = ["India", "India", "India", "India", "India"]
types = ["Restaurant", "Other", "Restaurant", "Religious", "Hospital"]
latlngs = [
LatLng(28.556765858466832, 77.24085192593027),
LatLng(28.54506337089723, 77.21650260907923),
LatLng(28.553394436597504, 77.24074977705357),
LatLng(28.556786989423074, 77.25379604088401),
LatLng(28.566587083464594, 77.23534244401858)
]



def send_post_request(i):
    url = "http://localhost:8080/v1/business"
    base_payload["name"] = places[i]
    base_payload["city"] = cities[i]
    base_payload["country"] = countries[i]
    base_payload["type"] = types[i]
    base_payload["latitude"] = latlngs[i].lat
    base_payload["longitude"] = latlngs[i].lng
    payload = base_payload
    headers = {"Content-Type": "application/json"}
    response = requests.post(url, data=json.dumps(payload), headers=headers)
    print(f"POST request sent to {url}")
    print(f"Payload: {payload}")
    print(f"Response status code: {response.status_code}")
    print(f"Response text: {response.text}")
    print("---")

# Send 5 POST requests with random variations
for i in range(0,5):
    send_post_request(i)
