build_dir = .build
app = ppp-grpc

build:
	rm -rf $(build_dir)/
	mkdir -p $(build_dir)
	GOOS=linux GOARCH=amd64 go build -v -o $(build_dir)/$(app)

copy:
	cp Dockerfile $(build_dir)/

docker_build:
	docker build $(build_dir)/ -t $(app)

docker_run: build copy docker_build
	docker rm -fv $(app) || echo "no container"
	docker run -p 5001:5000 \
		--name $(app) \
		$(app)