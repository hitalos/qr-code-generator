build:
	CGO_ENABLED=0 go build -ldflags '-s -w' -trimpath -o dist/qrcode .

IMAGE_BUILDER=$(shell [ -e /usr/bin/podman ] && echo podman || echo docker)
container-image:
	$(IMAGE_BUILDER) image build -t hitalos/qr-code-generator .

.PHONY: build container-image
