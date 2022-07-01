FROM ubuntu

WORKDIR /cli-tutor

RUN yes | unminimize

RUN apt -y install less
RUN apt -y install man

ENV TERM xterm-256color

RUN useradd -ms /bin/bash cli-pro
USER cli-pro
ADD ./bin/cli-tutor ./

# CMD ./cli-tutor
