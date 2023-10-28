# Alfred

Alfred is a gin rest api service that handles all crud operations in

## Local Development

API Framework: [gin](https://github.com/gin-gonic/gin)

### Run local server

```
go run main.go
```

## Testing

Testing framework: [goblin](https://github.com/franela/goblin)

### Running tests

```
go test ./tests/...
```

## Swagger

### Install he [swag](https://github.com/swaggo/swag) cli

```
go install github.com/swaggo/swag/cmd/swag@latest
```

### Updating Swagger documentation

```
swag init
```

### Formatting Swagger Documentation

```
swag fmt
```
