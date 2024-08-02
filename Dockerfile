# syntax=docker/dockerfile:1

FROM golang:1.22.3 AS build-stage

WORKDIR /go/src/github.com/justjack1521/mevway
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mevway cmd/mevway/main.go

FROM build-stage AS run-test-stage

RUN go test -v ./...

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /go/src/github.com/justjack1521/mevway/mevway mevway

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./mevway"]