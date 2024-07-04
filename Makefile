
format:
	gofmt -e -l -s -w .

test:
	go test -cover -v ./...

mock:
	mockgen -destination pkg/transport/mock/connection_mock.go -package mockTransport github.com/105-Code/multiplayer-server-sfs/pkg/transport Connection