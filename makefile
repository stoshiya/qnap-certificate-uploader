client = bin/qnap-certificate-uploader
clientFiles = main.go cli.go qnap.go
opt = -ldflags "-X main.gitDescribe=$(shell git describe --always)"

all: $(client)

$(client): $(clientFiles)
	go build $(opt) -o $(client) $(clientFiles)

test: $(clientFiles) $(testFiles)
	gotestcover -v -coverprofile=coverage.out

coverage: test
	go tool cover -html=coverage.out

clean:
	rm -f $(client) *.go~ coverage.out
