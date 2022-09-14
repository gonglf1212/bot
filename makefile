# Go parameters
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build -v
GOXBUILD=GOOS=linux GOARCH=amd64 $(GOCMD) build -v
GOARMBUILD=GOOS=linux GOARM=7 GOARCH=arm $(GOCMD) build -v
GOTEST=$(GOCMD) test

flags="-X 'main.buildTime=$(shell date '+%Y-%m-%d %H:%M:%S')' -X 'main.goVersion=$(shell go version)'"


bot:
	rm -rf target/bot/bot
	$(GOBUILD) -ldflags ${flags} -o target/bot/bot cmd/bot/main.go
	./target/bot/bot -configfile="./config/config.yaml" -env="local" -port=8880 -httpport=8880 -pprof=18880 -name="QQ频道机器人"