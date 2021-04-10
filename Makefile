
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
		-t eldius/k3s-dashboard-go \
		.
	docker tag eldius/k3s-dashboard-go eldius/k3s-dashboard-go:$(shell git rev-parse --short HEAD)

dockerrun: dockerbuild
	docker run -it --rm --name mocky -p 8080:8080 -p 8081:8081 eldius/k3s-dashboard-go:latest
