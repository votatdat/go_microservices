# for faster development
FROM alpine:latest

RUN mkdir /app

COPY loggerApp /app

CMD ["/app/loggerApp"]