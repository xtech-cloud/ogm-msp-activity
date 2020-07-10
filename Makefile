.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/activity/shared.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/activity/channel.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/activity/record.proto
