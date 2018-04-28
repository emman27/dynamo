test:
	ginkgo

cover:
	go test -coverprofile c.out ./...

install:
	dep ensure