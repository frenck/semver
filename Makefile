clean:
	rm -rf ./bin
build:
	mkdir -p ./bin
	go build -o ./bin/ semver.go 