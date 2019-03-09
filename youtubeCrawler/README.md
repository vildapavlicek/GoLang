"# youtubeCrawler" <br>
Change all settings in .env file<br>
Endpoint: localhost:8080/api/v1/link <br>
Adds link to crawl. Method POST only<br>
Payload example (only 1 link per request): <br>
/watch?v=DT61L8hbbJ4<br>
/watch?v=wOGu2j3PnFg<br>
/watch?v=MH9FyLsfDzw<br>
/watch?v=HZa1iFO0Juk<br>
<br>

Endpoint: localhost:8080/api/v1/stop<br>
Stops all go routines and closes all channels<br>
