# API Server

Simple Rest API using gin(framework) & gorm(orm)
https://github.com/wantedly/apig

## Start it
- go build main.js
- ./main

## Endpoint list

### Features Resource

```
GET    /features
GET    /features/:id
POST   /features
PUT    /features/:id
DELETE /features/:id
```

### Mockups Resource

```
GET    /mockups
GET    /mockups/:id
POST   /mockups
PUT    /mockups/:id
DELETE /mockups/:id
```

### Projects Resource

```
GET    /projects
GET    /projects/:id
POST   /projects
PUT    /projects/:id
DELETE /projects/:id
```

### Users Resource

```
GET    /users
GET    /users/:id
POST   /users
PUT    /users/:id
DELETE /users/:id
```

server runs at http://localhost:8080
