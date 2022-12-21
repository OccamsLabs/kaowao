##
# Project Title
#
# @file
# @version 0.1

build:
	go build .

run:
	go run ./main.go 

test:
	go test ./...

clean:
	rm -rf kaowao

lint:
	docker run --rm -it -v $(PWD):/app -w /app golangci/golangci-lint:v1.50.1 golangci-lint run -v

gosec:
	docker run --rm -it -v $(PWD):/app -w /app securego/gosec:latest ./...

fmt:
	gofmt -w -s .

ci: lint gosec test build

# end
