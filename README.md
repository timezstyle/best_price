# BEST_PRICE
* This is demo project to find best price of shops.
* Required GO 1.9 for [dep](https://github.com/golang/dep) (dependency tool).
* Required docker and docker-compose installed.

## How to run this project
```
# prepare selenium (make sure docker and docker-compose are already istalled)
docker-compose up -d

# set your GOPATH (optional)
export GOPATH=~/go

# clone project
go get github.com/timezstyle/best_price

# cd to project directory
cd $GOPATH/src/github.com/timezstyle/best_price

# listen at :3000
go run main.go

# to see more detail (optional)
go run main.go -h
```

## Find Best Price

GET /search?product_name={productName}&offset=0&limit=30

> for example: (To find best price of pen.)
```shell
curl 'http://localhost:3000/search?product_name=pen&offset=0&limit=30'
```

## Run test
```shell
# test all
GOCACHE=off go test -v ./...

# test Carrefour shop
GOCACHE=off go test -v ./... -run Test*/Carrefour
```