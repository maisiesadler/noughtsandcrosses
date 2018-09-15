FROM golang:latest 
ENV GOPATH=/usr/local/go/src/github.com/maisiesadler
RUN go get -u github.com/maisiesadler/noughtsandcrosses 
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN go build -o main . 
CMD ["/app/main"]