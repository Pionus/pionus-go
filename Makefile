
build_plugin:
	go build -buildmode=plugin -o storages/file.so storages/file.go

dev: build_plugin
	go run *.go
