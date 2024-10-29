# clean-restaurant-map

## Setup

### DB

```
$ docker compose -f environments/docker-compose.yml up -d
```

## Run

```
$ go build cmd/main.go
$ ./main
```

## Usage

### Get store

```
$ curl http://localhost:8080
```

### Save user

```
$ curl -H "Content-Type: application/json" -X POST -d "@example/create_user_api_example.json" http://localhost:8080/user
```

### Check user
- dbにcheck_user_api_example.jsonに記載されているemailがある => {}
- dbにcheck_user_api_example.jsonに記載されているemailがない => エラーメッセージとなる
```
$ curl -H "Content-Type: application/json" -X POST -d "@example/check_user_api_example.json" http://localhost:8080/user-check
```
