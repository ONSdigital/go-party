FROM golang:alpine AS build
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .

FROM alpine:latest
COPY --from=build /app/main .
EXPOSE 8081
CMD ["./main"]