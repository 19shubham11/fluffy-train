package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"testing"
)

func readFile(t *testing.T, path string) []byte {
	t.Helper()
	script, err := ioutil.ReadFile(path)

	if err != nil {
		t.Fatal(err)
	}
	return script
}

func execSQLScript(t *testing.T, db *sql.DB, script []byte) {
	t.Helper()
	_, err := db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}
}

func newTestDB(t *testing.T) (*sql.DB, func()) {
	db, err := sql.Open("mysql", "test_web:pass@/test_snippetbox?parseTime=true&multiStatements=true")
	if err != nil {
		t.Fatal(err)
	}

	setupScript := readFile(t, "./testdata/setup.sql")
	execSQLScript(t, db, setupScript)

	return db, func() {
		tearDownScript := readFile(t, "./testdata/teardown.sql")
		execSQLScript(t, db, tearDownScript)
		db.Close()
	}
}
