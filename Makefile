it:
	go mod tidy

build: bin/semaphore
bin/semaphore:
	go build -o bin/semaphore


clean:
	rm -rf bin/
