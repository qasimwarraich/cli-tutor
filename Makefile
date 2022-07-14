CC = go
MAIN = ./src/cli-tutor/...

all:
	${CC} run ${MAIN}

build:
	${CC} build -o ./bin ${MAIN}
	docker build -t qasimwarraich/cli-tutor .

docker:
	docker run -it qasimwarraich/cli-tutor

test:
	gotest -v ./src/...
