dev:
	go run main.go -config dev

production:
	go build -o dist/app main.go

production-run:
	./dist/app -config production

testing:
	go test -v ./test/...