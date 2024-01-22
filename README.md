

> [!WARNING]  
> When running a Go file in your terminal, you may encounter antivirus alerts. This is typically due to the nature of Go's compilation process, which involves the creation of standalone executables.

##  GO REST API + WEBSOCKETS

The following is a project in which it is implemented an HTTP REST server using Go's standard library.

It includes:

- Basic CRUD operations
- Authentication with a JSON Web Token (JWT)
- Implementation of a WebSocket to notify the creation of a new post
- Pagination of posts
- Protected endpoints for insert, update, and delete operations
- A Docker container for a PostgreSQL database


### Endpoints:

login 
```bash
http://localhost:5050/login
```
register:

```bash
http://localhost:5050/signup
```
json:
```json
{
    "email":"test2@mail.com",
    "password":"mypassword"
}
```
---

Protected endpoints (POST, PUT, DELETE)

they need jwt bearer token

Posts

```bash
localhost:5050/api/v1/posts
```
```bash
localhost:5050/api/v1/posts/:id
```
json:
```json
{
    "post_content": "brrrp this is a post"
}
```

GetAll paged
```bash
localhost:5050/posts?page=0
```

GetById

```bash
localhost:5050/posts/:id
```

WebSocket
```bash
localhost:5050/ws
```


## Clone and Run
Clone the project branch:

```bash
git clone https://github.com/NEV117/go/rest-websockets
```

Install dependencies:
```bash
go mod download
```

Build and run the database in Docker:
```bash
docker build . -t nev-ws-rest-db
```

```bash
docker run -p 54321:5432 nev-ws-rest-db
```

Run the project in the terminal:
```bash
go run main.go
```