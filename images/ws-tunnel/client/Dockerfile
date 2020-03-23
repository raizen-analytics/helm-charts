FROM golang:1.14.1 as build

COPY . /app/
WORKDIR /app
RUN apt update && apt install ca-certificates && update-ca-certificates
RUN CGO_ENABLED=0 GOOS=linux go build app.go

FROM scratch
COPY --from=build /app/app /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT [ "/app" ]