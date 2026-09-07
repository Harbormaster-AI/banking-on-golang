FROM golang:1.22-alpine AS builder

LABEL org.opencontainers.image.vendor="Harbormaster"
LABEL org.opencontainers.image.title="banking-on-golang"
LABEL org.opencontainers.image.version="0.0.1"
LABEL com.harbormaster.blueprint="Golang"
LABEL com.harbormaster.model="Banking Industry Domain Model"
LABEL com.harbormaster.generated="2026-09-07"
#LABEL com.harbormaster.certification="eeb624a6-acc4-46b2-827b-a5b914e1d728"

WORKDIR /app

COPY bin/banking-on-golang .

RUN chmod +x banking-on-golang

EXPOSE 8080

ENTRYPOINT ["./banking-on-golang"]