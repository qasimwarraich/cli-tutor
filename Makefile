CC = go
MAIN = ./src/cli-tutor/...

all:
	${CC} run ${MAIN}

build:
	${CC} build -o ./bin ${MAIN}
	docker build -t cli-tutor .

dockerbuild:
	docker build -t cli-tutor .

docker:
	docker run -it cli-tutor

test:
	gotest -v ./src/...
