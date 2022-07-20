CC = go
MAIN = main.go

all:
	${CC} run ${MAIN} --no-upload-log

log:
	${CC} run ${MAIN} 

build:
	${CC} build -o ./bin 
	docker build -t qasimwarraich/cli-tutor .

docker:
	docker run -it qasimwarraich/cli-tutor

test:
	gotest -v ./src/...
