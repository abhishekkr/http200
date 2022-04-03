package persist

var webhookSchemaSQL = `
BEGIN;

CREATE TABLE IF NOT EXISTS request (
	uuid		TEXT PRIMARY KEY,
	client		TEXT NOT NULL,
	method		TEXT NOT NULL,
	proto		TEXT NOT NULL,
	host		TEXT NOT NULL,
	request_uri	TEXT NOT NULL,
	fragment	TEXT,
	userinfo	TEXT,
	params		BLOB,
	headers		BLOB,
	body		BLOB,
	created_at	TEXT DEFAULT (STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW', 'localtime'))
);

COMMIT;
`

var insertRequestSQL = `
BEGIN;

INSERT INTO request (
	uuid,
	client,
	method, proto, host, request_uri, fragment, userinfo,
	params,
	headers,
	body
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);

COMMIT;
`
