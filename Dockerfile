FROM golang:1.22-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY ["./go.mod", "/go.sum", "./"]
RUN go mod download

#BUILD
COPY . .
RUN go build -o ./bin/portfolio-go cmd/portfolio-go/main.go


FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/portfolio-go /
COPY --from=builder /usr/local/src/.env /
COPY --from=builder /usr/local/src/migrations /migrations
COPY --from=builder /usr/local/src/web /web


CMD ["/portfolio-go"]
EXPOSE 8000