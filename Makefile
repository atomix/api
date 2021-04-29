.PHONY: build proto

build:
	docker run -it -v `pwd`:/go/src/github.com/atomix/atomix-api \
		-w /go/src/github.com/atomix/atomix-api \
		--entrypoint build/bin/compile_protos.sh \
		onosproject/protoc-go:stable
