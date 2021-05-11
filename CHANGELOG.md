
## Changelog

### Unreleased

> :empty

---

### [v0.5](https://github.com/abhishekkr/http200/tree/v0.5) add env var to enable Request Body display

* enable Request Body display by configuring `HTTP200_BODY=true` as environment variable; disabled by default

---

### [v0.4.2](https://github.com/abhishekkr/http200/tree/v0.4.2) fix /wiki Route

* fix `/wiki` Route and details in it

* fix README for `HTTP200_LISTEN_AT` doc

---

### [v0.4.1](https://github.com/abhishekkr/http200/tree/v0.4.1) methods POST, PUT, PATCH, DELETE, HEAD, OPTIONS for /200

* methods POST, PUT, PATCH, DELETE, HEAD, OPTIONS for /200

---

### [v0.4](https://github.com/abhishekkr/http200/tree/v0.4) custom routes for test via POST

* docker: docker run -it docker.io/abhishekkr/http:0.4

* listens default at port :9000, allows to change it using environment variable like LISTEN_AT=:8080

* shows this wiki at /wiki

* a simple placeholder http server providing /200,/400,/404,/500 for respective HTTP response codes

* returns 404 response code for any non-default or non-customized route

* un-handled route's response status code could be customized via env `HTTP200_DEFAULT_ROUTE` with values Route200, Route400, Route404, Route500

* add custom route and request method with response code and body if required

---
