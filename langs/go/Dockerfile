FROM odanado/os-command-injection:base


RUN apt update
RUN apt install -y wget tar


RUN wget https://dl.google.com/go/go1.10.linux-amd64.tar.gz -O go.tar.gz
RUN tar zxf go.tar.gz
RUN mv go/bin/* /usr/bin/
RUN rm -rf go go.tar.gz

RUN rm -rf /var/lib/apt/lists/* 
