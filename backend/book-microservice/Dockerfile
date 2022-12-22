# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.19
LABEL maintainer="crypto-bazaar"

WORKDIR /app

COPY ./ ./
RUN go mod download

RUN go build ./cmd/develop/main.go

ENV PORT=8080

EXPOSE 8080

CMD [ "./main" ]