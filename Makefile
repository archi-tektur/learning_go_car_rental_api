build: clear vendor
	go build -o bin/food-app cmd/*.go

swagger:
	swag init -d "./cmd/server,./internal/food/application" -o ./api --ot json

test: vendor
	go test -v ./...

clear:
	rm -rf bin/

IMAGE_NAME?=food-api
container: build/package/Dockerfile
	docker build -t $(IMAGE_NAME) -f build/package/Dockerfile .

vendor: go.mod go.sum
	GOPRIVATE="github.com/archi-tektur/car-rental-api" go mod vendor

.DEFAULT_GOAL=build
.PHONY=build swagger test clear container
