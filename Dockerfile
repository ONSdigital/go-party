FROM golang:alpine AS build
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .

FROM alpine:latest
RUN addgroup -g 988 partysvc && \
    adduser -u 988 -S partysvc -G partysvc
USER partysvc
COPY --from=build /app/main .
EXPOSE 8081
ENTRYPOINT ["./main"]