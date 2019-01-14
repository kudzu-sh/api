# Copyright © 2018 the Kudzu contributors.
# Licensed under the Apache License, Version 2.0; see the NOTICE file.

#
# Until Go has generics, implementing Kubernetes types in Go will involve lots of codegen.
# This Makefile does the codegen steps for Kudzu.
#

PACKAGE := kudzu.sh/api
GROUPS := kudzu/v1alpha1
BOILERPLATE := hack/header.go.txt
GEN := zz_generated

.PHONY: all build generate deepcopy clientset lister informer

all: build

build: deps generate
	go build ./...

deps:
	dep ensure -v

generate: deepcopy defaulter register openapi clientset lister informer

deepcopy: generators/deepcopy
	generators/deepcopy -i $(PACKAGE)/$(GROUPS) -O $(GEN).deepcopy -h $(BOILERPLATE)

defaulter: generators/defaulter
	generators/defaulter -i $(PACKAGE)/$(GROUPS) -O $(GEN).defaults -h $(BOILERPLATE)

register: generators/register defaulter
	generators/register -i $(PACKAGE)/$(GROUPS) -O $(GEN).register -h $(BOILERPLATE)

openapi: generators/openapi
	generators/openapi -i $(PACKAGE)/$(GROUPS) -O openapi -p $(PACKAGE)/openapi -h $(BOILERPLATE)

clientset: generators/client
	generators/client -p kudzu.sh/api --input-base 'kudzu.sh/api' --input $(GROUPS) -n client -h $(BOILERPLATE)
	sed -i '' -e 's/aPIs/apis/g' client/typed/kudzu/v1alpha1/api.go

lister: generators/lister
	generators/lister -i $(PACKAGE)/$(GROUPS) -p $(PACKAGE)/client/listers -h $(BOILERPLATE)

informer: generators/informer
	generators/informer -i $(PACKAGE)/$(GROUPS) -p $(PACKAGE)/client/informers --versioned-clientset-package $(PACKAGE)/client --listers-package $(PACKAGE)/client/listers -h $(BOILERPLATE)

generators/%: Gopkg.lock
	@mkdir -p generators
	go build -o generators/$* ./vendor/k8s.io/code-generator/cmd/$*-gen