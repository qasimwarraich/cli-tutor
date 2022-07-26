CC = go
MAIN = main.go

all:
	${CC} run ${MAIN} --no-upload-log

log:
	${CC} run ${MAIN}

build:
	mkdir -p bin
	${CC} build -o ./bin
	docker build -t qasimwarraich/cli-tutor .

install:
	mkdir -p bin
	${CC} build -o ./bin
	cp ./bin/cli-tutor /usr/local/bin

uninstall:
	rm -f /usr/local/bin/cli-tutor

clean:
	rm -rf bin

goinstall:
	${CC} install

gouninstall:
	rm ~/go/bin/cli-tutor

docker:
	docker run -it qasimwarraich/cli-tutor

test:
	gotest -v ./pkg/...
