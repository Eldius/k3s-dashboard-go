
clean:
	-rm *.log*
	-rm **/*.log*

start:
	DASHBOARD_PROMETHEUS_ENDPOINT=http://127.0.0.1:9090 \
		go run main.go start

test: clean
	go test ./... -cover
