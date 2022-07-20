FROM ubuntu

RUN yes | unminimize
RUN apt -y install less
RUN apt -y install man
RUN apt -y install curl

RUN useradd -ms /bin/bash cli-student
ADD ./bin/cli-tutor /usr/bin

WORKDIR /home/cli-student
USER cli-student
ENV TERM xterm-256color

# CMD ./cli-tutor
