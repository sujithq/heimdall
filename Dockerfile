ARG BUILDER_IMAGE=golang:alpine
ARG PROC_IMAGE=scratch

FROM ${BUILDER_IMAGE} AS builder

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

ENV USER=moducate
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR $GOPATH/src/github.com/moducate/heimdall
COPY go.mod .

ENV GO111MODULE=on
RUN go mod download
RUN go mod verify

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -ldflags='-w -s -extldflags "-static"' -a -o /usr/bin/heimdall main.go

FROM ${PROC_IMAGE}

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /usr/bin/heimdall /usr/bin/heimdall

EXPOSE 1470

USER moducate:moducate

ENTRYPOINT ["heimdall"]
CMD ["serve"]
