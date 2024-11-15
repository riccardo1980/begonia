FROM golang:latest AS builder
WORKDIR /build
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 go build


FROM scratch
WORKDIR /app
COPY --from=builder /build/begonia ./begonia
EXPOSE 8080
ENTRYPOINT ["./begonia"]