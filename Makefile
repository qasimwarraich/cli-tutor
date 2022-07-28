CC = go
MAIN = main.go

all:
	${CC} run ${MAIN} --no-upload-log -x --no-welcome

log:
	${CC} run ${MAIN}

build:
	mkdir -p bin
	${CC} build -o ./bin

docker:
	docker build -t qasimwarraich/cli-tutor .

dockerrun:
	docker run -it qasimwarraich/cli-tutor

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

test:
	gotest -v ./pkg/...
