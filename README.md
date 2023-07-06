# go-psql-redis-cities
### A simple web app that uses Go, Postgres as database and Redis as cache. 
### Display a list of cities and their traffic codes.

**Docker command to run postgresql container:**
`$ docker run -it -d -p 5432:5432 --name cities-postgre -e POSTGRES_PASSWORD=cities-pass -d postgres:alpine3.14`