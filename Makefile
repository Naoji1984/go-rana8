.PHONY:	build
build:
	go build -trimpath -ldflags "-w -s" -o app

.PHONY:	saveKeyFile
saveKeyFile:
	envsubst < secretKeyFile.json > keyFile.json

.PHONY:	deploy
deploy:
	docker pull gcr.io/google.com/cloudsdktool/google-cloud-cli:455.0.0 && \
	docker run -v "$(PWD):/go-rana8" --name gcloud-config gcr.io/google.com/cloudsdktool/google-cloud-cli:455.0.0 gcloud auth activate-service-account --key-file=/go-rana8/keyFile.json && \
	docker run --rm -v "$(PWD):/go-rana8" --volumes-from gcloud-config gcr.io/google.com/cloudsdktool/google-cloud-cli:455.0.0 gcloud run deploy $(SERVICE_ID) --project rana8-poc --region asia-northeast1 --platform managed --source /go-rana8/src --set-env-vars RANA8_DB_HOST=$(RANA8_DB_HOST),RANA8_DB_PORT=$(RANA8_DB_PORT),RANA8_DB_USER=$(RANA8_DB_USER),RANA8_DB_PASSWORD=$(RANA8_DB_PASSWORD),RANA8_DB_NAME=$(RANA8_DB_NAME)
