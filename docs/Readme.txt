migrate create -ext sql -dir migrations create_users

migrate -path migrations -database "postgres://login:pass@localhost:5432/httprest_api?sslmode=disable" up
Common example: "dbdriver://username:password@host:port/dbname?param1=true&param2=false"