#base on ubuntu
FROM ubuntu:18.04
LABEL description="build the data emulation image"

RUN mkdir /app
WORKDIR /app
#copy from build dir
COPY . /app
COPY data-emulation /app

#define default command
#ENTRYPOINT ["/data-emulation"]

#run the exe
CMD ["./data-emulation"]
