.PHONY:	build
build:
	go build -trimpath -ldflags "-w -s" -o app

.PHONY:	saveKeyFile
saveKeyFile:
	envsubst < secretKeyFile.json > keyFile.json

.PHONY:	deploy
deploy:
	docker pull gcr.io/google.com/cloudsdktool/google-cloud-cli:455.0.0 && \
	docker run -v "$(PWD):/go-rana8" --name gcloud-config gcr.io/google.com/cloudsdktool/google-cloud-cli:455.0.0 gcloud auth activate-service-account --key-file=/go-rana8/keyFile.json


