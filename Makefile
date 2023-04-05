##
# Project Title
#
# @file
# @version 0.1

build:
	go build .

run:
	go run ./main.go 

install:
	go install ./...
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


build-docker:
	docker build -t kaowao:latest .


testrun-docker:
	sh -c "docker run --rm -it --env KAOWAO_SALT=12345  -v $(PWD):/data kaowao:latest scan test2.yaml"
# end
