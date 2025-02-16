# Building Backend
FROM golang:alpine as pusher-server

WORKDIR /source
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs -o /dist ./pkg/main.go

# Runtime
FROM golang:alpine

COPY --from=pusher-server /dist /pusher/server

COPY locales /locales

EXPOSE 8444

CMD ["/pusher/server"]
