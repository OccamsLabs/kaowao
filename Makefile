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

# end
