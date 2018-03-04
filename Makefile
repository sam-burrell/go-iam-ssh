name = iam-ssh
profile = $(or $(AWS_PROFILE), signal-rnd)

build:
	go build -o bin/$(name) src/$(name).go

run:
	make build
	AWS_PROFILE=$(profile) bin/$(name)

build-static:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/$(name)-static src/$(name).go

run-static:
	make build-static
	docker build -t $(name) -f Dockerfile . --no-cache
	docker run -it $(name)
