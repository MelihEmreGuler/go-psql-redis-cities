# go-psql-redis-cities
### A simple web app that uses Go, Postgres as database and Redis as cache. 
### Display a list of cities and their traffic codes.

**Docker command to run postgresql container:**
```bash
$ docker run -it -d -p 5432:5432 --name cities-postgre -e POSTGRES_PASSWORD=cities-pass -d postgres:alpine3.18
```

**SQL command to create table:**
```sql
Create Table cities (
id SERIAL PRIMARY KEY,
name VARCHAR(25) NOT NULL,
code NUMERIC(2,0) NULL
);
```