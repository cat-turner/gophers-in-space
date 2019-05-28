## Example 1 

Adapted from: https://medium.com/@maumribeiro/my-first-graphql-server-in-go-responding-to-apollo-bd1c11426572

## What's this

An example that shows a few Graphql primary concepts:

## What does a graphql api look like to the client?

A single endpoint is used for all datasoures.

`router.Methods("POST").Path("/").Name("graphql").Handler(handler)`

The apis capabilities are defined in the Schema.

You use queries select fields from an entity.

```
query {
  astronauts {
    name

  }
}
```

You use mutations to modfiy data.

```
mutation{
  addCrewMember(input: {
    name: "Billy",
    id: 101,
    age:10
  }){
    age
    id
    name
  }
}
```

Types are building blocks that make up graphql Schemas and define:
- what data can be queried
- how the data can be manipulated
- relationships between types (ie Mission and Astronaut)
- the apis capabilities, via introspection

Schema Definition Language (SDL) is used to express what types are available within a schema and how they related to eachother.

## Types
Instead of endpoints, Graphql uses the Type System to describe what data can be queried and how it can be changed.

GraphQL type system categorizes several custom types, that can be defined by the code including:

Objects
Interfaces
Unions
Enums
Scalars
InputObjects

## This code

AstronautSchema: defines the schema. Has two fields, query (queryType) and mutation (mutationType).

Query is the root type, which means it represents all possible entry points into the api.

## Resolvers: let's talk about them

A resolver is a function.

Each field is backed by a function called a resolver, which produces the next value. If the field produces a scalar value (such as a string or int) then the execution of the query is complete. However if the field contains another object value, then it will call that objects' resolve function, until a scalar value is produced. 

From the docs - Resolvers have four aguments:
- obj: the previous object, which for root query field is not used
- args: arguments provided to the field in the GraphQL query
- context: holds contextual info such as which user is logged in (not used here)
- info: infomation about the execution state, which I have never used

implementation in graphql-go:

all of this in the code is defined in `graphql.ResolveParams`

```
type ResolveParams struct {
    // Source is the source value
    Source interface{}

    // Args is a map of arguments for current GraphQL request
    Args map[string]interface{}

    // Info is a collection of information about the current execution state.
    Info ResolveInfo

    // Context argument is a context value that is provided to every resolve function within an execution.
    // It is commonly
    // used to represent an authenticated user, or request-specific caches.
    Context context.Context
}
```

For example: Query: queryType -> astronautType

Mutation works the same way.

Object Types can refer to one another.


What have we learned here, from this code?
- How the client talks to the api
- Types, and how they are used to define what the api is able to do
- Queries, which are object types that fetch data
- Mutations, which are expected to mutate date
- Resolvers, functions that define how the data is retrieved and how it is manupulated in the query and mutations, respectively

Next: Queries and Mutations, in less code.