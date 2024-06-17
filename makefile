# VERSION=`git describe --tags`
VERSION=1.0.0
BUILD=`date +%FT%T%z`
BINARY="bin/go-split-pdf"

.PHONY: run
run:
	@echo "Running the program..."
	go run cmd/api/main.go

.PHONY: build
build:
	@echo "Building the program..."
	go build -ldflags="-X 'main.Version=${VERSION}' -X main.Build=${BUILD}" -o ${BINARY} cmd/api/main.go 

# Cleans our project: deletes binaries
.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: version
version:
	@echo $(VERSION)
	@echo $(BUILD)