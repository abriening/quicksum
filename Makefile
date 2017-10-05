build: fmt
	go build

install:
	go install

fmt:
	go fmt

test:
	go test -cover

test.coverage:
	go test -coverprofile=coverage.out && go tool cover -html=coverage.out

clean:
	rm -f quicksum
