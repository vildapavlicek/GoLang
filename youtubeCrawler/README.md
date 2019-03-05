"# youtubeCrawler" 

Endpoint: localhost:8080/api/v1/link
Adds link to crawl. Method POST only
Payload example (only 1 link per request): 
/watch?v=DT61L8hbbJ4
/watch?v=wOGu2j3PnFg
/watch?v=MH9FyLsfDzw
/watch?v=HZa1iFO0Juk


Endpoint: localhost:8080/api/v1/stop
Stops all go routines and closes all channels
