FROM golang:1.13-alpine AS build
WORKDIR /var/www
COPY . .

RUN go build -o server .
# RUN go run main.go

FROM alpine:3.12
EXPOSE 9999
COPY --from=build /var/www/server /server
CMD ["/server"]