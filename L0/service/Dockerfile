FROM golang

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o service ./service.go

ENTRYPOINT ["./consumer"]