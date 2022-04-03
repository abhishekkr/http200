package persist

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/abhishekkr/gol/golenv"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db           *sqlx.DB
	sqliteDbPath = golenv.OverrideIfEnv("HTTP200_SQLITE_DBPATH", "/tmp/http200.db")
)

type Request struct {
	UUID       string `db:"uuid"`
	Client     string `db:"client"`
	Method     string `db:"method"`
	Proto      string `db:"proto"`
	Host       string `db:"host"`
	RequestURI string `db:"request_uri"`
	Fragment   string `db:"fragment"`
	UserInfo   string `db:"userinfo"`
	Params     []byte `db:"params"`
	Headers    []byte `db:"headers"`
	Body       []byte `db:"body"`
	CreatedAt  string `db:"created_at"`
}

func init() {
	var err error
	db, err = sqlx.Connect("sqlite3", sqliteDbPath)
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(webhookSchemaSQL)
}

func InsertRequest(req *http.Request) {
	id := uuid.NewString()
	params := mapToBytes(req.URL.Query())
	headers := mapToBytes(req.Header)
	body := requestBody(req)
	db.MustExec(
		insertRequestSQL,
		id,
		req.RemoteAddr,
		req.Method, req.Proto, req.Host, req.RequestURI, req.URL.Fragment, req.URL.User.String(),
		params,
		headers,
		body,
	)
}

func mapToBytes(someMap interface{}) []byte {
	result, err := json.Marshal(someMap)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return result
}

func requestBody(req *http.Request) []byte {
	result, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return result
}
