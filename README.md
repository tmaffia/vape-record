# vape-record
Web based vaping tracker. Tracks your vaping habits

## Requirements

* Go
* PostgreSQL
* Bower
* Docker (Optional)

## Install

#### Install using `go get`
`go get github.com/tmaffia/vape-record`


#### Requires a postgres db, configure using conf.json

`cp conf.json.template conf.json`



#### Clone the [front-end](https://github.com/tmaffia/vape-record-front-end) into a directory named "public"

```
git clone https://github.com/tmaffia/vape-record-front-end.git public
cd public
bower install
```

## Deploy

Deploy using [Docker](https://www.docker.com/) with the included dockerfile, or run from the command line with the go tool

```
go install github.com/tmaffia/vape-record
vape-record
```
