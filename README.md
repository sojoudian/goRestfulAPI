# goRestfulAPI

test health check via curl command:
`curl http://localhost:8080/api/health`

run pgSQL database with docker:
`docker run --name comment-api-db -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres`

test APIs with cur;/Postman :
`localhost:8080/api/comment/1`

test, add new comment for a new post using post method and postman:
endpoint : `localhost:8080/api/comment`
type: `application/json`
content:
`{"slug": "/post2", "author": "Maziar"}`
