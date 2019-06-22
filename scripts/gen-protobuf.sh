#!/bin/bash

protoc --proto_path=. --go_out=. ./lib/protobuf/admin/admin.proto
protoc --proto_path=. --go_out=. ./lib/protobuf/blog/blog.proto