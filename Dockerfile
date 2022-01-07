from golang:latest

RUN mkdir /build
WORKDIR /build
COPY . .

RUN go build main.go

RUN cp main /usr/local/bin

EXPOSE 4000

CMD ["main", "0.0.0.0", "4000", "serverconfig.env"]