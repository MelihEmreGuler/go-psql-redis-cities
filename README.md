# go-psql-redis-cities
### A simple web app that uses Go, Postgres as database and Redis as cache. 
### Display a list of cities and their traffic codes.

**Docker command to run postgresql container:**
```bash
$ docker run -it -d -p 5432:5432 --name cities-postgre -e POSTGRES_PASSWORD=cities-pass -d postgres:alpine3.18
```

**Docker command to run rabbitMQ container:**
```bash
$ docker run -d -p 15672:15672 -p 5672:5672 --name my-rabbit rabbitmq:3.12-management-alpine
```

**Docker command to run redis container:**
```bash
$ docker run -p 6379:6379 --name cities-redis -d redis
```

**SQL query to create city table:**
```postgresql
Create Table cities (
id SERIAL PRIMARY KEY,
name VARCHAR(20) NOT NULL,
code NUMERIC(2,0) NULL
);
```

**RabbitMQ URL:**
```http request
http://localhost:15672/
```

**API URL:**
```http request
http://localhost:8080/
```

**API Endpoints:**
```http request
GET /cities
GET /cities/{id}
GET /cities/{name}
POST /cities {name, code}
PUT /cities/{id} {name, code}
DELETE /cities/{id}
```

**Sample Requests:**
```curl
CREATE ->   curl -X POST localhost:8080/city -d '{"name":"İstanbul", "code":34}' -v
READ ->     curl localhost:8080/city\?id=1  or  curl localhost:8080/city\?name=İstanbul 
READ ALL -> curl localhost:8080/city -v
UPDATE ->   curl -X PUT localhost:8080/city -d '{"id":1, "name":"İzmir", "code":35}' -v
DELETE ->   curl -X DELETE localhost:8080/city?\id=1
```

**Sample Response:**
```json
[
  {
    "Id": 1,
    "Name": "İzmir",
    "Code": 35
  },
  {
    "Id": 2,
    "Name": "Adana",
    "Code": 1
  },
  {
    "Id": 3,
    "Name": "İstanbul",
    "Code": 34
  },
  {
    "Id": 4,
    "Name": "Düzce",
    "Code": 81
  }
]
```


## Credits:
#### This project is based on the following [Go Eğitim Kampı - 302 & 303 - Database & External Communication](https://www.youtube.com/watch?v=Yf7Uu5fzAYA&t=2019s)

#### Thanks to [Emre Savcı](https://github.com/mstrYoda) for his great presentation and contribution to the [community](https://github.com/GoTurkiye).