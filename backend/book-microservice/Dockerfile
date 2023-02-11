FROM golang:1.19-alpine as dev

WORKDIR /goapp
COPY . .
RUN go mod download

EXPOSE 8080

RUN CGO_ENABLED=0 go build ./src/cmd/develop/main.go

FROM gcr.io/distroless/static-debian11 as prod
COPY --from=dev /goapp/main / 
COPY --from=dev /goapp/.env /
COPY --from=dev /goapp/src/data/images /src/data/images

EXPOSE 8080

CMD ["./main"]