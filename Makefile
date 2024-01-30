.PHONY:	build
build:
	go build -trimpath -ldflags "-w -s" -o app

.PHONY:	saveKeyFile
saveKeyFile:
	echo $(KEY_FILE)

.PHONY:	setup
setup:
	gcloud version && \
	docker pull gcr.io/google.com/cloudsdktool/google-cloud-cli:455.0.0 && \
	docker run -ti --name gcloud-config gcr.io/google.com/cloudsdktool/google-cloud-cli:455.0.0 gcloud auth activate-service-account --key-file=XXXXXXX@rana8-poc.iam.gserviceaccount.com.json

