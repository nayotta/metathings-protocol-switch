PROTOC := $(shell which protoc)

GRPC_FLAGS := -Iproto -I/usr/local/include \
	-I $(GOPATH)/src

PROTO_FILES := $(shell find proto -name *.proto)

.PHONY: build

build-go:
	$(PROTOC) $(GRPC_FLAGS) \
	--descriptor_set_out=proto/service.protoset \
	--go_out=plugins=grpc:go/proto \
	$(PROTO_FILES)

build: build-go
