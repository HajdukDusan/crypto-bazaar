FROM golang:1.19-alpine as dev

# we download the lib here before copy so it can be cached

WORKDIR /app
COPY . /app/

RUN go mod download

RUN CGO_ENABLED=0 go build ./cmd/develop/main.go

FROM gcr.io/distroless/static-debian11 as prod

COPY --from=dev /app/main /
CMD ["/main"]

# builds main.exe in root folder
# RUN go build ./cmd/develop/main.go 

# ENV PORT=8080

# EXPOSE 8080

# CMD [ "./main" ]