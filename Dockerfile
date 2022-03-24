FROM golang:1.17.5 AS Builder
ENV CGO_ENABLED 0
COPY . /app
WORKDIR /app
RUN go build -o webapi

FROM alpine:3.15
WORKDIR /app
COPY --from=Builder /app/webapi /app/webapi
EXPOSE 9090
CMD ["./webapi"]