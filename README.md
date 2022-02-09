# Simple Go CRUD App with PostgreSQL Database

## Types

| Field    | Type   | Reference       |
|----------|--------|-----------------|
| ID       | int64  | json:"id"       |
| Name     | string | json:"name"     |
| Location | string | json:"location" |
| Age      | int64  | json:"age"      |

## Methods

#### GET "/api/user/{id}"
#### GET "/api/user"
#### POST "/api/newuser"
#### PUT "/api/user/{id}"
#### DELETE "/api/deleteuser/{id}"

## Create Registry
```
curl -d '{"name":"paulo", "location":"rio", "age":38}'  -H "Content-Type: application/json" -X POST http://localhost:8080/api/newuser
```

## Update Registry
```
curl -d '{"name":"pedro", "location":"rio", "age":38}'  -H "Content-Type: application/json" -X PUT http://localhost:8080/api/user/{id}
```

## List Users
```
curl -H "Content-Type: application/json" -X GET http://localhost:8080/api/user
```

## User Details
```
curl -H "Content-Type: application/json" -X GET http://localhost:8080/api/user/{id}
```

## Delete User
```
curl -H "Content-Type: application/json" -X DELETE http://localhost:8080/api/deleteuser/{id}
```

## Database Connection
### URL Variable Example
```
POSTGRES_URL=postgres://user:pass@127.0.0.1:5432/databasename?sslmode=disable
```

## Docker Reference
### Example of PostgreSQL Database Backend
```
mkdir -p /postgresql
chown -R 1001:1001 /postgresql

docker run \
   --name postgresql \
   -e POSTGRESQL_USERNAME=goappuser \
   -e POSTGRESQL_PASSWORD=goapppass \
   -e POSTGRESQL_DATABASE=goappdb \
   -p 5432:5432 \
   -v /postgresql:/bitnami/postgresql/data \
   quay.io/bitnami/postgresql:13
```
### Create Tables
```
docker run \
   --name=migrate \
   -e POSTGRES_URL=postgres://goappuser:goapppass@127.0.0.1:5432/goappdb?sslmode=disable \
   goapp \
   /app/migrate
```
### Start Goapp
```
docker run -d \
   --name=goapp \
   -p 8080:8080 \
   -e POSTGRES_URL=postgres://goappuser:goapppass@127.0.0.1:5432/goappdb?sslmode=disable \
   goapp
```