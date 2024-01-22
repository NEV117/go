ARG GO_VERSION=1.21.4

FROM golang:${GO_VERSION}-alpine as builder

RUN go env -w GOPROXY=direct
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates && update-ca-certificates

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./

RUN CGO_ENABLED=0 go build \
    -installsuffix "static" \
    -o /nev-rest-ws

FROM scratch AS runner
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs

COPY .env ./

COPY --from=builder /nev-rest-ws /nev-rest-ws 

EXPOSE 5050

ENTRYPOINT [ "/nev-rest-ws" ]