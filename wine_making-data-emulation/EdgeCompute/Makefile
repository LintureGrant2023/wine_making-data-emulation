#.PHONY告诉编译器，make的对象是假的，只有规则没有依赖
.PHONY: all docker clean
all:
		go build -o data-emulation main.go
docker:
		docker build . -t data-emulation:v1
clean:
		rm -rf ./data-emulation

