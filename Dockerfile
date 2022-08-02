FROM alpine

RUN apk update
RUN apk add curl

RUN mkdir /app
WORKDIR /app
COPY mcdc-ws mcdc-ws
RUN chmod 555 *

# Healthcheck
COPY docker/healthcheck.sh /usr/healthcheck/healthcheck.sh
RUN chmod 555 /usr/healthcheck/healthcheck.sh
HEALTHCHECK --timeout=60s CMD sh /usr/healthcheck/healthcheck.sh

USER nobody
CMD ["./mcdc-ws"]
