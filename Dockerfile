FROM golang:alpine

WORKDIR /build

COPY go.*  .

RUN go mod download

COPY . .

RUN go build -o /main cmd/main.go

EXPOSE 8000

ENV CONFIG_PATH './configs/prod.yml'

CMD ["/main"]