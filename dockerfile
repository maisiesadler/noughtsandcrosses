FROM golang:latest 
RUN go get -u github.com/maisiesadler/noughtsandcrosses 
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN go build -o main . 
CMD ["/app/main"]