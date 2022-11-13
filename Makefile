cobraboot: go.mod go.sum $(shell find . -name "*.go")
	go build -o $@ github.com/tgf9/cobraboot/cmd/cobraboot
