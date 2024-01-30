.PHONY:	build
build:
	go build -trimpath -ldflags "-w -s" -o app

.PHONY:	saveKeyFile
saveKeyFile:
	envsubst < secretKeyFile.json > keyFile.json

.PHONY:	setup
setup:
	gcloud version && \
	docker pull gcr.io/google.com/cloudsdktool/google-cloud-cli:455.0.0 