import requests
import json
import os
from guessit import api

def scan(path):
    videos = [
		os.path.join(dir, file) for dir, _, files in os.walk(path) for file in files
	]

    to_create = []
    for path in videos:
        print('reading...', path)
        raw = api.guessit(path, 
		{
			"single_value": True
		})

        if "mimetype" not in raw or not raw["mimetype"].startswith("video"):
            continue

        if raw.get("type") == "movie":
            info = {
                "name": raw["title"],
                "year":str(raw.get("year")),
                "path": path
            }
        elif raw.get("type") == "episode":
            info = {
                "showName": raw["title"],
                "season":str(raw.get("season")),
                "episodeNumber":str(raw.get("episode")) if raw.get("episode") is not None  else raw.get("episode_title"),
                "name":raw.get("episode_title"),
                "path": path
            }
        to_create.append(info)
    
    return to_create
        
if __name__ == '__main__':
    print('starting...')
    to_create = scan("/media/rafael/Expansion/Series/The IT Crowd")

    for x in to_create:
        print(x)
        response = requests.post("http://localhost:3000/api/episodes", data=json.dumps(x), headers={'Content-Type': 'application/json'})
        print(response.text)