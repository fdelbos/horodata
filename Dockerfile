FROM alpine

RUN apk --update upgrade && \
    apk add \
        curl \
        ca-certificates \
        tzdata \
        && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

ADD horo /
EXPOSE 3000
CMD ["/horo"]
