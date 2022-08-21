# simple-user
Simple API using Go, Postgresql with sqlc.

## Running Locally
Enter to the root folder: `$ cd simple-user`
### 1. Setup the `.env`
### 2. Tidy the go modules up
```
$ go mod tidy
```
### 3. Prepare the Postgres db
```
$ make postgres
$ make createdb
$ make migrateup
```
### 4. Run the server
- Using go run:
```
$ make server
```

- Using docker-compose:
```
$ docker compose up
```

