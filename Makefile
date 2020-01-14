.PHONY: start


OUTFILE=pionus
PIDFILE=pionus.pid


build_plugin:
	go build -buildmode=plugin -o storages/file.so storages/file.go

dev: build_plugin
	go run *.go

build: build_plugin
	go build -o $(OUTFILE) *.go

stop:
ifeq (,$(wildcard $PIDFILE))
	kill -9 $$(cat $(PIDFILE))
	rm -f $(PIDFILE)
endif

start: build
	./$(OUTFILE) & echo $$! > $(PIDFILE)

restart: stop
	./$(OUTFILE) & echo $$! > $(PIDFILE)