FROM golang:1.23 as builder

WORKDIR /src

COPY go.mod go.mod
COPY go.sum go.sum
COPY ./pkg ./pkg
COPY chats-service/cmd/messages-service-mock/ ./
RUN go mod tidy && go mod download

RUN CGO_ENABLED=0 go build -o /bin/server main.go

FROM alpine:latest as final

COPY --from=builder /bin/server /bin/
ENTRYPOINT [ "/bin/server" ]