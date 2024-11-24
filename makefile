build:
	go build -o gitana
	sudo make install
install:
	sudo cp -f gitana /usr/bin/gitana
	sudo chmod +x /usr/bin/gitana