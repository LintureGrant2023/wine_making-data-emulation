#base on ubuntu
FROM ubuntu:18.04
LABEL description="test how to docker build images"

RUN mkdir /app
WORKDIR /app
#copy from build dir
COPY data-emulation /app

#define default command
ENTRYPOINT ["/data-emulation"]

#run the exe
CMD ["data-emulation"]
