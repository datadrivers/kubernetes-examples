REGISTRY_SERVICE ?= quay.io
REGISTRY_ORGANISATION ?= your_organisation
REGISTRY_IMAGE ?= simple_webserver

.PHONY: build
build:
	# Tag your Image for the registry you are using
	docker build -t $(REGISTRY_SERVICE)/$(REGISTRY_ORGANISATION)/$(REGISTRY_IMAGE) .
	docker push $(REGISTRY_SERVICE)/$(REGISTRY_ORGANISATION)/$(REGISTRY_IMAGE)
