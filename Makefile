BINARY_NAME=oc-prompt

build:
	go build -o ${BINARY_NAME} main.go

run:
	./${BINARY_NAME}

build_and_run: build run

build_prod:
	go build -o ${BINARY_NAME} main.go
	sudo cp ${BINARY_NAME} /usr/local/bin/

clean:
	go clean
	rm ${BINARY_NAME}