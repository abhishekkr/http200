## http200

> [get binary of latest release here](https://github.com/abhishekkr/http200/releases/latest)

> your friendly http server to use as placeholder for integration points of your service under development
>
> docker: `docker run -p 9000:9000 -it docker.io/abhishekkr/http:0.4.2`

[![HTTP200 to dynamically configure HTTP Requests for Testing with Fake Response | 2 Minutes HowTo](https://i9.ytimg.com/vi_webp/tYPUDdi0o9c/mqdefault.webp?time=1608803700000&sqp=CPTKkf8F&rs=AOn4CLA0EicxIX0G-PcX7846ntOD-auf-A)](https://www.youtube.com/watch?v=tYPUDdi0o9c "Video: HTTP200 to dynamically configure HTTP Requests for Testing with Fake Response | 2 Minutes HowTo")

#### it provides:

* listens default at port `:9000`, allows to change it using environment variable like `HTTP200_LISTEN_AT=:8080`

* enable printing request body using environment variable `HTTP200_BODY=true`

* shows this wiki at `/wiki`

* a simple placeholder http server providing `/200`,`/400`,`/404`,`/500` for respective HTTP response codes

* returns `404` response code for any non-default or non-customized route

* un-handled route's response status code could be customized via env `HTTP200_DEFAULT_ROUTE` with values `Route200`, `Route400`, `Route404`, `Route500`

* add custom route and request method with response code and body if required


#### todo:

* custom routes from persistent VCR-cassettes, OpenAPI specs, RAML

---
---
