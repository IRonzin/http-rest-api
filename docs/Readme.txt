migrate create -ext sql -dir migrations create_users

migrate -path migrations -database "postgres://login:pass@localhost:5432/httprest_api?sslmode=disable" up
Common example: "dbdriver://username:password@host:port/dbname?param1=true&param2=false"


Запуск бенчмарков:
go test -bench . -benchmem -cpuprofile=profile_cpu // CPU профиль

go test -bench . -benchmem -memprofile=profile_mem // Memory профиль
go test -bench . -benchmem -count=8 > profile_1 // профиль для утилиты benchstat


go tool pprof profile_cpu // top web list <имя функции>
go tool pprof -http=:7272 'http://localhost:3366/debug/pprof/heap?seconds=10' // в интерактивном режиме

http://localhost:3366/debug/pprof // можно подгружать профили в корректном формате



go tool trace name.trace // просмотр трейса CPU