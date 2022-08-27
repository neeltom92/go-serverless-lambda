.PHONY: build package light deploy clean

build:
	dep ensure -v
	env GOARCH=amd64 GOOS=linux go build -mod=mod -ldflags="-s -w" -o main main.go

package: build
	zip deployment.zip main

light: package
	sls deploy -f lambdacode

deploy: package
	sls deploy --stage production

clean:
	@rm -f deployment.zip
	@rm -f main
