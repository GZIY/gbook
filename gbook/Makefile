.PHONY: docker
docker:
	@rm gbook || true
	@GOOS=linux GOARCH=arm go build -o gbook .
	@docker rmi -f flycash/gbook:v0.0.1
	@docker build -t flycash/gbook-live:v0.0.1 .