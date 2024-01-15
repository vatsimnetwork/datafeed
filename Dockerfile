FROM golang:1.21.5 as builder

WORKDIR /workspace
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /workspace/bin/server /workspace/cmd/server/main.go

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /workspace/bin/server /server

USER 65532:65532
EXPOSE 8080

ENTRYPOINT ["/server"]
