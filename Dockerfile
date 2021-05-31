FROM ubuntu:18.04

LABEL MAINTAINER="lucups@gmail.com"

ARG ZIP_NAME=bgm-linux-amd64-0.1.1.zip

RUN  apt update \
  && apt install -y wget \
  && rm -rf /var/lib/apt/lists/*

RUN cd / \
    && mkdir -p /opt/bgm \
    && /bin/wget https://github.com/lucups/bgm/releases/download/v0.1.1/$ZIP_NAME \
    && tar -xf $ZIP_NAME - d /opt/bgm\
    && rm $ZIP_NAME

WORKDIR /opt/bgm
CMD ["nohup ./bgm > logs/bgm.log &"]
EXPOSE 9600
