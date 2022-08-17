FROM ubuntu

RUN yes | unminimize
RUN apt -y install less
RUN apt -y install man
RUN apt -y install tree
RUN apt -y install curl

RUN useradd -ms /bin/bash cli-student
ADD ./bin/cli-tutor /usr/bin

USER cli-student
RUN mkdir /home/cli-student/tutor/
WORKDIR /home/cli-student/tutor/

ENV TERM xterm-256color

CMD cli-tutor
