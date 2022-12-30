all:
	go build -o brush ./cmd/cli/main.go
clean:
	rm -fr brush
install:
	cp ./brush /usr/local/bin

install-cfg:
	mkdir -p /etc/brush
	cp ./cmd/brush/config.ini /etc/brush