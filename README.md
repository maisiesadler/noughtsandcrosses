# noughtsandcrosses

## How to run

Make sure you have docker installed.

In a new directory, create the following file

```
package main

import "github.com/maisiesadler/noughtsandcrosses/server"

func main() {
	server.StartServer()
}
```

Create a dockerfile

```
FROM golang:latest
RUN go get -u github.com/maisiesadler/noughtsandcrosses/game
RUN go get -u github.com/maisiesadler/noughtsandcrosses/server
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

Example response:
```
O Turn
   |   |   
---+---+---
   |   |   
---+---+---
 X |   |   
 ```
