package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

// func readFile(t *testing.T, path string) []byte {
// 	t.Helper()
// 	script, err := ioutil.ReadFile(path)

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	return script
// }

// func execSQLScript(t *testing.T, db *sql.DB, script []byte) {
// 	t.Helper()
// 	_, err := db.Exec(string(script))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

func newTestDB() (*sql.DB, func()) {
	mysqlUsername := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASS")

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/", mysqlUsername, mysqlPassword)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dbName := "snippetbox"
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE " + dbName)
	if err != nil {
		panic(err)
	}

	dsn = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?multiStatements=true&&parseTime=true", mysqlUsername, mysqlPassword, dbName)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 4)

	m := doMigrations(db)
	return db, func() {
		err := m.Steps(-1)
		if err != nil {
			fmt.Println("Error Migrating DOWN")
			panic(err)
		}
		db.Close()
	}
}

func doMigrations(db *sql.DB) *migrate.Migrate {
	driver, _ := mysql.WithInstance(db, &mysql.Config{})

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	migrationsPath := "file:///" + pwd + "/migrations"
	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}
	err = m.Steps(-1)
	err = m.Steps(1)
	if err != nil {
		fmt.Println("Error Migrating UP")
		panic(err)
	}
	return m
}
