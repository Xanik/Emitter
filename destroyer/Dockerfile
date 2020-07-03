FROM debian 

WORKDIR /

ADD config.yml /
ADD destroyer /

EXPOSE 80/tcp

STOPSIGNAL SIGINT

ENTRYPOINT ["/destroyer"]