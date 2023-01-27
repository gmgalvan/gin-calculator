## Gin-Calculator
Simple Calculator using Gin framework

## Setup and run

# Executing with Docker
```
docker build --tag gin-calculator .
docker run -p 8080:8080 gin-calculator 
```

# Execututing with go command
```
go get
go run main.go
```

## Examples

- Sum numbers
```
curl -H "Content-Type: application/json" -X POST -d '{"quantities":[2,4,45,45,1]}' http://localhost:8080/calculator/v1/sum
```

- Return
```
{"result": 97}
```

## Updating documentation

- Install swaggo [https://github.com/swaggo/swag](repo)
```
go install github.com/swaggo/swag/cmd/swag@v1.8.10
```

- generate docs
```
swag init --parseDependency
```

- Open link to show docs
http://localhost:8080/swagger/index.html