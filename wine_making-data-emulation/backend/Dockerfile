FROM ubuntu:18.04
#FROM golang:1.20.6
LABEL description="test how to build backend image"
#RUN apt update
#RUN apt install -y vim
RUN mkdir /app
COPY backend /app
COPY settings.yaml /app
WORKDIR /app
#RUN go build -o backend main.go
EXPOSE 8089
ENTRYPOINT ["./backend"]
#run exe should be backend, not ./backend
CMD ["./backend"]
#CMD ["go","run","backend"]
