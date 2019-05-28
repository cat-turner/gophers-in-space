
## Example 2 - Use a lib that will write boilerplate code

0. Read example

    https://gqlgen.com/getting-started/

1. Define the schema in schema.graphql

2. Run command to create boilerplate code

    `go run github.com/99designs/gqlgen init`

3. Implement the resolvers

4. Ran the server
    `go run server/server.go`


How it works:
Libs like gqlgen and graph-gophers/graphql-go use go generate
