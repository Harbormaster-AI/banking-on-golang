
FROM golang:1.22-alpine AS builder

LABEL org.opencontainers.image.vendor="Harbormaster"
LABEL org.opencontainers.image.title="bankingOnGolang"
LABEL org.opencontainers.image.version="0.0.1"
LABEL com.harbormaster.blueprint="Golang"
LABEL com.harbormaster.model="Banking Industry Domain Model"
LABEL com.harbormaster.generated="2026-09-20"
#LABEL com.harbormaster.certification="845b0106-b422-4f00-9759-090ffd24660e"

WORKDIR /app

COPY bin/bankingOnGolang .

RUN chmod +x bankingOnGolang

EXPOSE 8080

ENTRYPOINT ["./bankingOnGolang"]