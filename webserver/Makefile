REGISTRY_SERVICE ?= quay.io
REGISTRY_ORGANISATION ?= your_organisation
REGISTRY_IMAGE ?= simple_webserver
RELEASE_TAG ?= 0.0.0

.PHONY: build
build:
	# Tag your Image for the registry you are using
	docker build -t $(REGISTRY_SERVICE)/$(REGISTRY_ORGANISATION)/$(REGISTRY_IMAGE):$(RELEASE_TAG) .
	docker push $(REGISTRY_SERVICE)/$(REGISTRY_ORGANISATION)/$(REGISTRY_IMAGE)
