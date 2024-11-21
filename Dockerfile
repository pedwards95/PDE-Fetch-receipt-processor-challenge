FROM public.ecr.aws/docker/library/golang:1.23.3-alpine3.20 AS base

WORKDIR /app

ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

FROM public.ecr.aws/c2u4x1k5/golangci-lint:v1.59.0-alpine AS golint

RUN golangci-lint run --timeout 5m

FROM base AS builder

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/PDE-Fetch-receipt-processor-challenge ./cmd/PDE-Fetch-receipt-processor-challenge

FROM public.ecr.aws/docker/library/alpine:3.20.3

COPY --from=builder /app/bin/PDE-Fetch-receipt-processor-challenge /bin/PDE-Fetch-receipt-processor-challenge

ENV docker="true"

EXPOSE 8080

ENTRYPOINT ["/bin/PDE-Fetch-receipt-processor-challenge"]