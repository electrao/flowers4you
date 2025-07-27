# flowers4you
Daily Florist App

# Run App

`go run main.go`

# Run docker app

`docker-compose up`


# API

- GET /messages
    - `curl http://localhost:8080/messages`
- POST /messages
    - `curl -X POST http://localhost:8080/messages -H "Content-Type: application/json" -d '{"text":"Hello API"}'
`