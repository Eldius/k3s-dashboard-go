
clean:
	-rm *.log*
	-rm **/*.log*

start:
	DASHBOARD_PROMETHEUS_ENDPOINT=http://127.0.0.1:9090 \
		go run main.go start

test: clean
	go test ./... -cover

dockerbuild:
	docker build \
		-t eldius/mock-server-go \
		.
	docker tag eldius/mock-server-go eldius/mock-server-go:$(shell git rev-parse --short HEAD)

dockerrun: dockerbuild
	docker run -it --rm --name mocky -p 8080:8080 -p 8081:8081 eldius/mock-server-go:latest
