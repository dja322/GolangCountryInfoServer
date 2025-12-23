# Introduction

This is a simple API code for calling to a database and receiving information on countries. Written in go.

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
There are admin endpoints where registered admins can see users, add user, remove user, and update user

## How to run
This project uses docker to run builds
Run docker compose up --build

this is start the containers on your machine, api calls can be made after 
it is finished initializing

## Current plans
Updates for this project
 * Containerize Database and server functionality
 * Add admin functions
