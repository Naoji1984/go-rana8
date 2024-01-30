.PHONY:	build
build:
	go build -trimpath -ldflags "-w -s" -o app

.PHONY:	setup
setup:
	gcloud version && \
	gcloud components update --version 455.0.0 && \
	gcloud version
