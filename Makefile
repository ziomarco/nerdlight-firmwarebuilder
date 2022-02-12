.DEFAULT_GOAL := build

build:
	make clean
	@echo "Building..."
	go build
	GOOS=windows GOARCH=amd64 go build -o dist/nerdlight-fb-amd64.exe main.go
	GOOS=windows GOARCH=386 go build -o dist/nerdlight-fb-386.exe main.go
	GOOS=darwin GOARCH=amd64 go build -o dist/nerdlight-fb-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o dist/nerdlight-fb-darwin-arm64 main.go
	GOOS=linux GOARCH=386 go build -o dist/nerdlight-fb-linux-386 main.go
	GOOS=linux GOARCH=amd64 go build -o dist/nerdlight-fb-linux-amd64 main.go
	GOOS=linux GOARCH=arm64 go build -o dist/nerdlight-fb-linux-arm64 main.go
	rm -rf nerdlight-firmwarebuilder
clean:
	@echo "Cleaning up"
	rm -rfv dist
	rm -rf nerdlight-firmwarebuilder
release:
	@echo "Releasing tag $VERSION"
	git tag -a $VERSION -m "Release $VERSION"
	git push origin $VERSION