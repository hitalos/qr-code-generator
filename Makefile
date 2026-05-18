build:
	CGO_ENABLED=0 go build -ldflags '-s -w' -trimpath -o dist/qrcode .

IMAGE_BUILDER=$(shell [ -e /usr/bin/podman ] && echo podman || echo docker)
container-image:
	$(IMAGE_BUILDER) image build -t hitalos/qr-code-generator .

sec:
# go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...

# go install github.com/securego/gosec/v2/cmd/gosec@latest
	gosec ./...

# go install github.com/anchore/grype@latest
	grype .

# go install github.com/aquasecurity/trivy/cmd/trivy@latest
	trivy fs .

lint:
	go fix -diff ./...

# go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
	golangci-lint run ./...


.PHONY: build container-image
