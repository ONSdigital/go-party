FROM alpine:latest
RUN apk update && apk upgrade
COPY build/linux-amd64/bin/partysvc /usr/local/bin/
EXPOSE 8081
RUN addgroup -g 988 partysvc && \
    adduser -u 988 -S partysvc -G partysvc
USER partysvc
ENTRYPOINT ["/usr/local/bin/partysvc"]
