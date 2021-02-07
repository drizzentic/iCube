# The iCube Code Challenge



The project contains 4 different challenges stored in different folders:

  - Darts
  - Knapsack
  - Game of thrones
  - Rest API

# Running the  applications

## Darts (Question 1)
Darts requires [Golang](https://golang.org/)
The programs runs on the cli using the installed go

```sh
$ cd dart
$  go build .
$ ./Darts 0 10 (0 and 10 represent the scores x, y)
```
## Knapsack (Question 2)
Knapsack requires [Golang](https://golang.org/)
The programs runs on the cli using the installed go
The program used the shared json to test the algorithm
``Items: [ { "weight": 5, "value": 10 }, { "weight": 4, "value": 40 }, { "weight": 6, "value": 30 }, { "weight": 4, "value": 50 }``

```sh
$ cd knapsack
$  go build .
$ ./knapsack
```
## Game of Thrones (Question 3)
Game of thrones requires [Node.js](https://nodejs.org/) v4+ to run.
```sh
$ cd gameofthrones
$ npm install
$ yarn start
```

## Rest API (Question 4)
Knapsack requires latest version of [Golang](https://golang.org/)
```sh
$ cd restApi
$ go build .
$ ./restAPi
```
This will start an http server running on port 8000. Ensure the port is not being utilized by another application. This is the documentation for the RESTAPI https://www.getpostman.com/collections/99ef3d5b4c627faa82cb
