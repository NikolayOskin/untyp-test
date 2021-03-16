build:
	docker build -t untyp-test .
start:
	docker run --env PORT=$(port) -p $(port):$(port) -it untyp-test