# Mockidi

### Prerequisites
- Go1.*
- Redis

```
go mod tidy
go run server/main.go
```

### Create Mock URL
```
GET /create HTTP/1.1
Host: localhost:3000
Content-Type: application/json

{
    "status": 200,
    "body": "{\"ID\":\"123456789\"}",
    "headers": {
        "Access-Control-Allow-Origin": "example.com"
    }
}
```
_Response_
```
{
    "hash": "6e4724f5-952a-43d5-8e34-c85973ae7cee"
}
```

### Call Mock URL
```
GET /6e4724f5-952a-43d5-8e34-c85973ae7cee HTTP/1.1
Host: localhost:3000
```
_Response_
```
{
    "ID": "123456789"
}
```