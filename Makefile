APP = deployer

golint:
	golangci-lint run -v

easyjson:
	easyjson -all pkg/github/ 