go get -u github.com/swaggo/swag/cmd/swag
swag init
go run ./pkg/gorm/gen/generate.go
