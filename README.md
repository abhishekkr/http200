## http200

> your friendly http server to use as placeholder for integration points of your service under development


#### it provides:

* listens default at port `:9000`, allows to change it using environment variable like `LISTEN_AT=:8080`

* shows this wiki at `/wiki`

* a simple placeholder http server providing `/200`,`/400`,`/404`,`/500` for respective HTTP response codes

* returns `404` response code for any non-default or non-customized route


#### todo:

* add custom route and request method with response code and body if required

---



