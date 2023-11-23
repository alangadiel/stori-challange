# Stori Challenge

## How to run the server
First, create a file named *.env* with the following variables, adding the corresponding values:
```sh
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=
EMAIL_PASSWORD=
```

Then, execute the following command:
```sh
docker-compose up -d
```

## API example

```sh
curl --location 'http://localhost:8080/balance' \
--header 'Content-Type: application/json' \
--data-raw '{
    "file_name": "/transactions.csv",
    "email": "alangadiel@gmail.com"
}'
```