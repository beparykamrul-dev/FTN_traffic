FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 go build -trimpath -ldflags='-s -w' -o /out/ftn-traffic ./cmd/ftn-traffic

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/ftn-traffic /ftn-traffic
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/ftn-traffic"]
