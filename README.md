# BEST_PRICE
* This is demo project to find best price of shops.
* Required GO 1.9 for [dep](https://github.com/golang/dep) (dependency tool)

## How to run this project
```
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

GET /search?productName={productName}

> for example: (To find best pen price.)
```shell
curl http://localhost:3000/search?productName=pen
```

## Run test
```shell
# test all
GOCACHE=off go test -v ./...

# test Carrefour shop
GOCACHE=off go test -v ./... -run Test*/Carrefour
```