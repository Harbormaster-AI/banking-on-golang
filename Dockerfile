FROM golang:1.22-alpine AS builder

LABEL org.opencontainers.image.vendor="Harbormaster"
LABEL org.opencontainers.image.title="banking-on-golang"
LABEL org.opencontainers.image.version="0.0.1"
LABEL com.harbormaster.blueprint="Golang"
LABEL com.harbormaster.model="Banking Industry Domain Model"
LABEL com.harbormaster.generated="2026-09-07"
#LABEL com.harbormaster.certification="24ab8573-87e6-43cd-b9eb-c93d36e49eb1"

WORKDIR /app

COPY bin/banking-on-golang .

RUN chmod +x banking-on-golang

EXPOSE 8080

ENTRYPOINT ["./banking-on-golang"]