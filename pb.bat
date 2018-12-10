IF NOT EXIST message mkdir message
protoc -I proto/ proto/message.proto --proto_path=%GOPATH%/src --go_out=plugins=grpc:message --govalidators_out=./message