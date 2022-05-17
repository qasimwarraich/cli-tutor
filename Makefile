CC = go
MAIN = ./src/cli-tutor/...

all:
	${CC} run ${MAIN}

build:
	${CC} build -o ./bin ${MAIN}

