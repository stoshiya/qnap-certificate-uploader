client = bin/qnap-certificate-uploader
clientFiles = main.go cli.go qnap.go
opt = -ldflags "-X main.gitDescribe=$(shell git describe --always)"

all: $(client)

$(client): $(clientFiles)
	go build $(opt) -o $(client) $(clientFiles)

test: $(clientFiles) $(testFiles)
	echo 'mode: atomic' > coverage.txt && go list ./... | xargs -n1 -I{} sh -c 'go test -covermode=atomic -coverprofile=coverage.tmp {} && tail -n +2 coverage.tmp >> coverage.txt' && rm coverage.tmp

coverage: test
	go tool cover -html=coverage.txt

clean:
	rm -f $(client) *~ coverage.txt
