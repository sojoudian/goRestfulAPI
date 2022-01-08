# goRestfulAPI

Use Docker PostgresQL for this project, using below commands:
`docker run --name comment-api-db -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres`

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

to validate the post request use the endpoint : `localhost:8080/api/comment` again but this time with `GET` method

Update comment, use the endpoint : `localhost:8080/api/comment/$ID` change the &ID with comment ID you want to update, for example `localhost:8080/api/comment/3` and change the request body with same method:
type: `application/json`
content:
`{"slug": "/post2", "author": "Maziar"}`

to :
type: `application/json`
content:
`{"slug": "/post2", "author": "Maziar Sojoudian"}`

beside other fields, `UpdatedAt` will also update
`"UpdatedAt": "2022-01-07T14:19:11.341906-05:00",`

Delete a comment:
hit `localhost:8080/api/comment/$ID` endpoint with `DELETE` method, which `$ID` will be deleted, e.g. :
`localhost:8080/api/comment/7`

You have to get the:
`{ "Message": "Comment successfully deleted" }`
response.

to validate the delete request hit endpoint : `localhost:8080/api/comment` again with `GET` method, and you shouldn't be able to see comment with ID=7

## Todo

create test for the rest of the APIs
