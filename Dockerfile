FROM ubuntu

RUN yes | unminimize
RUN apt -y install less
RUN apt -y install man
RUN apt -y install curl

RUN useradd -ms /bin/bash cli-pro
USER cli-pro

WORKDIR /home/cli-pro/
ADD ./bin/cli-tutor ./
ENV TERM xterm-256color

# CMD ./cli-tutor
