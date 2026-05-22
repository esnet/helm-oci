VERSION := $(shell sed -n 's/version:.*"\(.*\)"/\1/p' plugin.yaml)
LDFLAGS := -X main.version=$(VERSION)

HELM_PLUGINS := $(shell helm env HELM_PLUGINS 2>/dev/null)

GO ?= go

.PHONY: build
build:
	mkdir -p bin/
	$(GO) build -v -o bin/helm-oci -ldflags="$(LDFLAGS)" .

.PHONY: test
test:
	$(GO) test -v ./... -race

.PHONY: install
install: build
	mkdir -p $(HELM_PLUGINS)/helm-oci/bin
	cp bin/helm-oci $(HELM_PLUGINS)/helm-oci/bin/
	cp plugin.yaml $(HELM_PLUGINS)/helm-oci/

.PHONY: uninstall
uninstall:
	rm -rf $(HELM_PLUGINS)/helm-oci

.PHONY: clean
clean:
	rm -rf bin/

.PHONY: snapshot
snapshot:
	goreleaser release --snapshot --clean --skip=sign,publish

.PHONY: dist
dist: export CGO_ENABLED=0
dist:
	rm -rf build/ release/
	mkdir -p build/helm-oci/bin release/
	cp plugin.yaml install-binary.sh build/helm-oci/
	GOOS=linux GOARCH=amd64 $(GO) build -o build/helm-oci/bin/helm-oci -trimpath -ldflags="$(LDFLAGS)" . && \
		tar -C build/ -zcvf release/helm-oci-linux-amd64.tgz helm-oci/
	GOOS=linux GOARCH=arm64 $(GO) build -o build/helm-oci/bin/helm-oci -trimpath -ldflags="$(LDFLAGS)" . && \
		tar -C build/ -zcvf release/helm-oci-linux-arm64.tgz helm-oci/
	GOOS=darwin GOARCH=amd64 $(GO) build -o build/helm-oci/bin/helm-oci -trimpath -ldflags="$(LDFLAGS)" . && \
		tar -C build/ -zcvf release/helm-oci-darwin-amd64.tgz helm-oci/
	GOOS=darwin GOARCH=arm64 $(GO) build -o build/helm-oci/bin/helm-oci -trimpath -ldflags="$(LDFLAGS)" . && \
		tar -C build/ -zcvf release/helm-oci-darwin-arm64.tgz helm-oci/
	GOOS=windows GOARCH=amd64 $(GO) build -o build/helm-oci/bin/helm-oci.exe -trimpath -ldflags="$(LDFLAGS)" . && \
		tar -C build/ -zcvf release/helm-oci-windows-amd64.tgz helm-oci/
