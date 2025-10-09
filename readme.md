# Introduction

This is a simple API code for calling to a database and receiving information on. Written in go.

## User Endpoints
Currently only one user endpoint at /api/
Takes in *api_key* and *country* and returns a json currently formatted as such:
```
{
    "GDP": int,
    "Population": int,
    "CapitolCity": String,
    "Continent": String,
    "SizeInSqMiles": int,
    "Country": String
}
```

## Admin Endpoints
Not yet implemented, plans to implement

## How to run
Install go's mysql driver
``github.com/go-sql-driver/mysql``


go run main.go
or
go build main.go and run exectuable

## Current plans
Updates for this project
 * Containerize Database and server functionality
 * Add admin functions