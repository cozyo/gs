.PHONY: all build test clean run

GO_EXEC_BIN_NAME = gs.exe

all: build

build: clean
	go build -o E:\go\workspace\bin\${GO_EXEC_BIN_NAME} .

clean:
	del E:\go\workspace\bin\${GO_EXEC_BIN_NAME}