# My Gin Service App
Routes: GET /health, GET /hello, GET /items/{id}, POST /echo (JSON {"message": "..."}).
go.mod and go.sum are included and complete, so `go build ./...` and `go test ./...` work with no extra steps.
Run: go run .   ->   http://localhost:8080/hello
