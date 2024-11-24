build:
	go build -o gitana
install:
	sudo cp gitana /usr/bin/gitana
	sudo chmod +x /usr/bin/gitana