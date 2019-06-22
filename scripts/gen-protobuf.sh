#!/bin/bash

protoc --proto_path=. --js_out=./lib/protobuf --go_out=. ./lib/protobuf/admin.proto
protoc --proto_path=. --js_out=./lib/protobuf --go_out=. ./lib/protobuf/blog.proto
protoc --proto_path=. --js_out=./lib/protobuf --go_out=. ./lib/protobuf/category.proto