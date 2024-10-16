FROM golang:1.22.6-alpine3.20 AS BUILD

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -o main cmd/api/main.go

FROM scratch
WORKDIR /service
COPY --from=BUILD /build/main .

ENTRYPOINT [ "./main" ]
