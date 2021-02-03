all: build

build:
#	../swagbins/swag init --parseDependency --parseInternal
	../swagbins/swag init
	go build -o ../bin/

clean:
	rm -f ../bin/bookshelf*
	rm -f ../bin/*log
	rm -f ../cover.out
	rm -f ./docs/docs.go
	rm -f ./docs/go.sum
	rm -f ./docs/swagger*

test:
	go test
	go test commons
	go test books
# go test -v -coverprofile="../cover.cov"
# go tool cover -func ../cover.cov
# go tool cover -html ../cover.cov
