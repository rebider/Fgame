#!/bin/sh
protoc --proto_path=../../../../../../../../  --proto_path=./  --go_out=./ chat.proto 