FROM debian 

WORKDIR /

ADD config.yml /
ADD deathstar /

EXPOSE 80/tcp

STOPSIGNAL SIGINT

ENTRYPOINT ["/deathstar"]