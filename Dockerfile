from golang:latest

RUN mkdir /build
WORKDIR /build
COPY . .

RUN go build main.go

RUN cp main /usr/local/bin

CMD ["main", "0.0.0.0", "$PORT|4000", "serverconfig.env"]