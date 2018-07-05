FROM webhippie/alpine:latest

LABEL maintainer="Thomas Boerger <thomas@webhippie.de>" \
  org.label-schema.name="Templater" \
  org.label-schema.vendor="Thomas Boerger" \
  org.label-schema.schema-version="1.0"

RUN apk add --no-cache ca-certificates mailcap bash && \
  addgroup -g 1000 templater && \
  adduser -D -h /var/lib/templater -s /bin/bash -G templater -u 1000 templater

USER templater
ENTRYPOINT ["/usr/bin/templater"]
CMD ["help"]

COPY dist/binaries/templater-*-linux-amd64 /usr/bin/
