FROM golang:1.19 AS BuildStage

WORKDIR /app

COPY . .

RUN go build

FROM alpine:latest

WORKDIR /

COPY --from=BuildStage /app/saml-idp /saml-idp

EXPOSE 8080

user nonroot:nonroot

ENTRYPOINT ["/saml-idp"]
