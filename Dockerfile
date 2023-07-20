FROM golang:alpine
WORKDIR /go/src/be-service-saksi
COPY . .

RUN cp api-specification/openapi.yaml saksi/delivery/http/openapi
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o be-service-saksi app/main.go

FROM alpine:latest
EXPOSE 8222

RUN apk --no-cache add ca-certificates
RUN apk add --no-cache tzdata
WORKDIR /app/
COPY --from=0 /go/src/be-service-saksi/be-service-saksi .
COPY --from=0 /go/src/be-service-saksi/db/migration ./db/migration
CMD ["./be-service-saksi"]
