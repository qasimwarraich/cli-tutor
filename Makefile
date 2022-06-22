CC = go
MAIN = ./src/cli-tutor/...

all:
	${CC} run ${MAIN}

build:
	${CC} build -o ./bin ${MAIN}
	sudo docker build -t cli-tutor-test .

dockerbuild:
	sudo docker build -t cli-tutor-test .

docker:
	sudo docker run -it cli-tutor-test

test:
	gotest -v ./src/...
