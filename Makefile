
clean:
	-rm *.log*
	-rm **/*.log*

start:
	DASHBOARD_PROMETHEUS_ENDPOINT=http://127.0.0.1:9090 \
		go run main.go start

test: clean
	go test ./... -cover

dockerbuild:
	docker buildx build \
		--push \
		--platform linux/arm/v7,linux/arm64/v8,linux/amd64 \
		--tag eldius/k3s-dashboard-go:latest \
		--tag eldius/k3s-dashboard-go:$(shell git rev-parse --short HEAD) \
		.

testdockerbuild:
	docker buildx build \
		--platform linux/arm/v7 \
		--tag eldius/k3s-dashboard-go:latest \
		--tag eldius/k3s-dashboard-go:$(shell git rev-parse --short HEAD) \
		.

dockerrun: dockerbuild
	docker run -it --rm --name mocky -p 8080:8080 -p 8081:8081 eldius/k3s-dashboard-go:latest

dockertest:
	docker build \
		-t eldius/mock-server-go \
		.
