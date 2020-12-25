This is a web application written in go using http templates from the standard library (https://golang.org/pkg/html/template/)

## Application
A notes/snippet app with basic CRUD operations. I have used mysql as the database for now, but any db can be used as the db methods are abstracted in the `pkg` directory.

The main routes are :
- `/snippets/` Returns the 10 latest snippets
- `/snippets?id=x` Returns a snippet with id `x`

Since I am using html templates, the endpoints return the final html instead of json data.

## Dependencies
- [Go MySQL driver](https://github.com/go-sql-driver/mysql)
- [Alice](https://github.com/justinas/alice) for clean middleware chaining

## Preview
- Clone the repo and from the root directory run
```
go run ./cmd/web
```

Server will start by default on port `4000`, but the flag `-addr` can be specified to run on a specific port.
Currently, it is expected that MySQL is running on the same host as that of the application, and the env variables `MYSQLUSER` and `MYSQLPASS` must be present to establish the connection.


## WIP
This app is a work in progress and a lot of things are missing, mainly -
- DB Migrations whenever the app starts (currently I ran the queries manually)
- Proper POST and DELETE endpoints
- Proper routing (does not support path params for now)
- Integration tests for the db
- Add users and login
- Host this somewhere?
