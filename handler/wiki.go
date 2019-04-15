package handler

import "fmt"

var (
	wiki = `
	<html>
	<head>
	  <title>http200: your placeholder service</title>
	</head>
	<body>
	  <h3>http200</h3>
	  <div>it provides:
	    <ul>
		  <li>listens default at port <code>:9000</code>, allows to change it using environment variable like <code>LISTEN_AT=:8080</code></li>
		  <li>shows this wiki at <code>/wiki</code></li>
		  <li>a simple placeholder http server providing <code>/200</code>,<code>/400</code>,<code>/404</code>,<code>/500</code> for respective HTTP response codes</li>
		  <li>returns 404 response code for any non-default or non-customized route</li>
		</ul>
	  </div>
	  <div>todo:
	    <ul>
		  <li>add custom route and request method with response code and body if required</li>
		</ul>
	  </div>
	  <div></div>
	</body>
	</html>
	`
)

func RouteWiki(ctx *Context) {
	fmt.Fprintf(ctx.ResponseWriter, wiki)
}
