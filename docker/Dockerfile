FROM golang:1.17

WORKDIR /go/src

COPY go.mod .
COPY go.sum .
COPY load-test-fake-backend.go ./load-test-fake-backend.go

RUN go build -a -installsuffix cgo -o load-test-fake-backend ./

EXPOSE 4000

CMD ["./load-test-fake-backend"]
