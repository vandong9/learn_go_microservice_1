build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-cli .

run:
	# docker run consignment-cli
	docker run -e MICRO_REGISTRY=mdns \
      consignment-cli ./consignment.json \
      99f808f3-c225-499b-b898-9ea9b47809ab