CC = go
MAIN = ./src/cli-tutor/...

all:
	${CC} run ${MAIN}

build:
	${CC} build -o ./bin ${MAIN}
	docker build -t cli-tutor-test .

dockerbuild:
	docker build -t cli-tutor-test .

docker:
	docker run -it cli-tutor-test

test:
	gotest -v ./src/...
