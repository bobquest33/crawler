package crawler

var (
	queryCreateTableTpl = `
CREATE TABLE IF NOT EXISTS "%s" (
    url     TEXT      PRIMARY KEY,
    type    INT       NOT NULL,
    status  INT       NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT statement_timestamp(),
    updated TIMESTAMP NOT NULL DEFAULT statement_timestamp()
);
`
	queryClearTable = `
TRUNCATE TABLE "%s";
`
	querySetStatus = `
INSERT INTO "%s" (
    url,
    type,
    status
) VALUES ($1, $2, $3) ON CONFLICT(url) DO UPDATE SET status = $3;
`
	queryGetAll = `
SELECT url, type, status FROM "%s";
`
)
