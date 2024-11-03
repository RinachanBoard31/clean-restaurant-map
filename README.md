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

### Login user
- dbにlogin_user_api_example.jsonに記載されているemailがある => {}
- dbにlogin_user_api_example.jsonに記載されているemailがない => エラーメッセージとなる
```
$ curl -H "Content-Type: application/json" -X POST -d "@example/login_user_api_example.json" http://localhost:8080/login
```
