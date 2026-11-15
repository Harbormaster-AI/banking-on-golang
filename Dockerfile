
FROM golang:1.27.4-bookworm AS builder

LABEL org.opencontainers.image.vendor="Harbormaster"
LABEL org.opencontainers.image.title="bankingOnGolang"
LABEL org.opencontainers.image.version="0.0.1"
LABEL com.harbormaster.blueprint="Golang"
LABEL com.harbormaster.model="Banking Industry Domain Model"
LABEL com.harbormaster.generated="2026-11-14"
#LABEL com.harbormaster.certification="4ae343e9-f2cc-4af3-ad5b-0b1792843559"

WORKDIR /app

COPY bin/bankingOnGolang .

RUN chmod +x bankingOnGolang

EXPOSE 8080

ENTRYPOINT ["./bankingOnGolang"]