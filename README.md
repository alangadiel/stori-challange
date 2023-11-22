# Stori Challenge

## How to run the server
Execute the following commands, adding the email password:
```sh
docker-compose up -d
export EMAIL_PASSWORD=...
go run cmd/server.go
```

## API example

```sh
curl --location 'http://localhost:8080/balance' \
--header 'Content-Type: application/json' \
--data-raw '{
    "file_name": "transactions.csv",
    "email": "alangadiel@gmail.com"
}'
```