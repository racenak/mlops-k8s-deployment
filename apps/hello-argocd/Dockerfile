FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 go build -o /hello-argocd .

FROM alpine:3.21
RUN apk add --no-cache ca-certificates
COPY --from=build /hello-argocd /hello-argocd
EXPOSE 8080
CMD ["/hello-argocd"]
