# noughtsandcrosses

## How to run

Make sure you have docker installed.

Create a dockerfile

```
FROM golang:latest 
ENV GOPATH=/usr/local/go/src/github.com/maisiesadler
RUN go get -u github.com/maisiesadler/noughtsandcrosses 
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN go build -o main . 
CMD ["/app/main"]
```

Run 

```
docker build . -t noughtsandcrosses
docker run -d -p 50:8080 noughtsandcrosses
```

## How to play

### View the board
GET http://localhost:50

Example response:
```
X Turn
   |   |   
---+---+---
   |   |   
---+---+---
   |   |   
```

### Place counter
POST http://localhost:50/place
body: {
	"counter": "X",
	"X": 0,
	"Y": 2
}
```
O Turn
   |   |   
---+---+---
   |   |   
---+---+---
 X |   |   
 ```