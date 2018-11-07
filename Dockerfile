FROM alpine:latest

ADD ./simpleapi /

EXPOSE 8080

CMD ["/simpleapi"]