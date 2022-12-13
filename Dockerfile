FROM docker.io/library/golang:1.19-alpine AS builder

WORKDIR /app
COPY go.mod go.sum .
RUN go mod download && apk -U add upx
COPY . .
RUN CGO_ENABLED=0 go build -ldflags '-s -w' -trimpath -o dist/qrcode . && upx dist/qrcode

FROM scratch

WORKDIR /app
COPY --from=builder /app/dist/qrcode .

CMD ["/app/qrcode"]
