## http200

> [get binary of latest release here](https://github.com/abhishekkr/http200/releases/latest)

> your friendly http server to use as placeholder for integration points of your service under development
>
> docker: `docker run -it docker.io/abhishekkr/http:0.4.1`

#### it provides:

* listens default at port `:9000`, allows to change it using environment variable like `HTTP200_LISTEN_AT=:8080`

* shows this wiki at `/wiki`

* a simple placeholder http server providing `/200`,`/400`,`/404`,`/500` for respective HTTP response codes

* returns `404` response code for any non-default or non-customized route

* un-handled route's response status code could be customized via env `HTTP200_DEFAULT_ROUTE` with values `Route200`, `Route400`, `Route404`, `Route500`

* add custom route and request method with response code and body if required


#### todo:

* custom routes from persistent VCR-cassettes, OpenAPI specs, RAML

---
---
