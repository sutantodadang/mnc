# mnc

## Installation

To run this project need:

1. [Taskfile](https://taskfile.dev/installation/)
2. [Go](https://go.dev/doc/install) 1.23
3. [Postman](https://www.postman.com/downloads/)
4. [Docker](https://docs.docker.com/engine/install/)

run this command on terminal it will download dependencies

```bash
task setup
```

```bash
go mod tidy
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`GOOSE_DRIVER`

`GOOSE_DBSTRING`

`GOOSE_MIGRATION_DIR`

`PORT`

`JWT_SECRET`

`KAFKA_ADDR`

`KAFKA_GROUP_ID`

`KAFKA_TOPIC`

## Run Locally

Run with docker

```bash
docker-compose up -d
```

Run without docker

```bash
task run
```

## Documentation

import using postman this json

```json
{
  "info": {
    "_postman_id": "20b2eea8-f671-4897-8279-c8385f9cbc11",
    "name": "mnc",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
    "_exporter_id": "14623263"
  },
  "item": [
    {
      "name": "user",
      "item": [
        {
          "name": "register",
          "request": {
            "method": "POST",
            "header": [],
            "body": {
              "mode": "raw",
              "raw": "{\r\n    \"first_name\": \"dadang\",\r\n    \"last_name\": \"sutanto\",\r\n    \"phone_number\": \"0000000000\",\r\n    \"address\": \"jalan-jalan\",\r\n    \"pin\": \"123456\"\r\n}",
              "options": {
                "raw": {
                  "language": "json"
                }
              }
            },
            "url": {
              "raw": "http://localhost:7575/api/v1/register",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "register"]
            }
          },
          "response": []
        },
        {
          "name": "login",
          "request": {
            "method": "POST",
            "header": [],
            "body": {
              "mode": "raw",
              "raw": "{\r\n    \"phone_number\":\"0000000000\",\r\n    \"pin\":\"123456\"\r\n}",
              "options": {
                "raw": {
                  "language": "json"
                }
              }
            },
            "url": {
              "raw": "http://localhost:7575/api/v1/login",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "login"]
            }
          },
          "response": []
        },
        {
          "name": "update profile",
          "request": {
            "auth": {
              "type": "bearer",
              "bearer": [
                {
                  "key": "token",
                  "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk3MDIzMDMsInN1YiI6ImExNzk4YTJhLWViODAtMTFlZi04ZTA1LTAyNDJhYzFhMDAwMiJ9.Mh5nI828n8gZzAxoPPM3prr1OG3oy8iFs_EbUUebgB8",
                  "type": "string"
                }
              ]
            },
            "method": "PUT",
            "header": [],
            "body": {
              "mode": "raw",
              "raw": "{\r\n    \"first_name\": \"dadang edit\",\r\n    \"last_name\": \"sutanto edit\",\r\n    \"address\": \"jalan desa\"\r\n}",
              "options": {
                "raw": {
                  "language": "json"
                }
              }
            },
            "url": {
              "raw": "http://localhost:7575/api/v1/profile",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "profile"]
            }
          },
          "response": []
        }
      ]
    },
    {
      "name": "topup",
      "item": [
        {
          "name": "add balance topup",
          "request": {
            "auth": {
              "type": "bearer",
              "bearer": [
                {
                  "key": "token",
                  "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk2OTg3MTksInN1YiI6ImExNzk4YTJhLWViODAtMTFlZi04ZTA1LTAyNDJhYzFhMDAwMiJ9.4ihtWHwLF4yKD5TAnnSYQQzPiwqlkjlT0fGhEXb9e58",
                  "type": "string"
                }
              ]
            },
            "method": "POST",
            "header": [],
            "body": {
              "mode": "raw",
              "raw": "{\r\n    \"amount\": 1000000\r\n}",
              "options": {
                "raw": {
                  "language": "json"
                }
              }
            },
            "url": {
              "raw": "http://localhost:7575/api/v1/topup",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "topup"]
            }
          },
          "response": []
        }
      ]
    },
    {
      "name": "payment",
      "item": [
        {
          "name": "payment",
          "request": {
            "auth": {
              "type": "bearer",
              "bearer": [
                {
                  "key": "token",
                  "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk2OTg3MTksInN1YiI6ImExNzk4YTJhLWViODAtMTFlZi04ZTA1LTAyNDJhYzFhMDAwMiJ9.4ihtWHwLF4yKD5TAnnSYQQzPiwqlkjlT0fGhEXb9e58",
                  "type": "string"
                }
              ]
            },
            "method": "POST",
            "header": [],
            "body": {
              "mode": "raw",
              "raw": "{\r\n    \"amount\": 25000,\r\n    \"remarks\": \"beli pulsa\"\r\n}",
              "options": {
                "raw": {
                  "language": "json"
                }
              }
            },
            "url": {
              "raw": "http://localhost:7575/api/v1/pay",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "pay"]
            }
          },
          "response": []
        }
      ]
    },
    {
      "name": "transfer",
      "item": [
        {
          "name": "transfer",
          "request": {
            "auth": {
              "type": "bearer",
              "bearer": [
                {
                  "key": "token",
                  "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk3MDU4MDMsInN1YiI6ImExNzk4YTJhLWViODAtMTFlZi04ZTA1LTAyNDJhYzFhMDAwMiJ9.timS9zk-1-f7OIcRVSmha5KXgJHlGhA1sFWkno785DQ",
                  "type": "string"
                }
              ]
            },
            "method": "POST",
            "header": [],
            "body": {
              "mode": "raw",
              "raw": "{\r\n    \"target_user\": \"9c8474e4-eb80-11ef-8e05-0242ac1a0002\",\r\n    \"amount\": 50000,\r\n    \"remarks\": \"uang bulanan 3\"\r\n}",
              "options": {
                "raw": {
                  "language": "json"
                }
              }
            },
            "url": {
              "raw": "http://localhost:7575/api/v1/transfer",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "transfer"]
            }
          },
          "response": []
        }
      ]
    },
    {
      "name": "transaction",
      "item": [
        {
          "name": "transaction report",
          "request": {
            "auth": {
              "type": "bearer",
              "bearer": [
                {
                  "key": "token",
                  "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk2OTg3MTksInN1YiI6ImExNzk4YTJhLWViODAtMTFlZi04ZTA1LTAyNDJhYzFhMDAwMiJ9.4ihtWHwLF4yKD5TAnnSYQQzPiwqlkjlT0fGhEXb9e58",
                  "type": "string"
                }
              ]
            },
            "method": "GET",
            "header": [],
            "url": {
              "raw": "http://localhost:7575/api/v1/transactions",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "transactions"]
            }
          },
          "response": []
        }
      ]
    }
  ]
}
```

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Authors

- [@sutantodadang](https://www.github.com/sutantodadang)
