build:
	protoc -I. --go_out=plugins=micro:. proto/*.proto