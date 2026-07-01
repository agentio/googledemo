FROM golang:1.26.4 AS builder
WORKDIR /app
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -ldflags="-s -w" -trimpath -o googledemo .

FROM gcr.io/distroless/static-debian13
COPY --from=builder /app/googledemo /usr/local/bin/googledemo
ENTRYPOINT ["/usr/local/bin/googledemo"]
CMD ["googledemo"]
