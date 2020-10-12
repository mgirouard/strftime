export GOFLAGS=-mod=vendor

all: bin/strftime

test:
	go test -v ...

clean:
	$(RM) -r bin/

bin/strftime:
	go build -o $@ cmd/strftime/main.go
