import os
from guessit import api

def scan(path):
    videos = [
		os.path.join(dir, file) for dir, _, files in os.walk(path) for file in files
	]

    to_create = []
    for path in videos:
        raw = api.guessit(path, 
		{
			"single_value": True
		})

        if "mimetype" not in raw or not raw["mimetype"].startswith("video"):
            continue

        if raw.get("type") == "movie":
            info = {
                "title": raw["title"],
                "year":raw.get("year"),
                "path": path
            }
        elif raw.get("type") == "episode":
            info = {
                "title": raw["title"],
                "season":raw.get("season"),
                "episode_nbr":raw.get("episode") if raw.get("episode") is not None  else raw.get("episode_title"),
                "episode_title":raw.get("episode_title"),
                "absolute":raw.get("episode") if "season" not in raw else None,
                "year":raw.get("year"),
                "path": path
            }
        to_create.append(info)
    
    return to_create
        
if __name__ == '__main__':
    to_create = scan("<path>")

    print(to_create)