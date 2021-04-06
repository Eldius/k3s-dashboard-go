
clean:
	-rm *.log*
	-rm **/*.log*

start:
	go run main.go start

test: clean
	go test ./... -cover
