package schema

import (
	"github.com/graphql-go/graphql"
)

var (
	astronautData      map[int]Astronaut
	astronautType      *graphql.Object
	mutationType       *graphql.Object
	astronautInputType *graphql.Object
	AstronautSchema    graphql.Schema
)

// Astronaut ...
type Astronaut struct {
	Id   int
	Name string
	Age  int
}

// init initializes the dummy data
func init() {
	benny := Astronaut{
		Id:   1,
		Name: "Benny",
		Age:  20,
	}
	alice := Astronaut{
		Id:   2,
		Name: "Alice",
		Age:  20,
	}
	astronautData = map[int]Astronaut{
		1: benny,
		2: alice,
	}

	// astronautType is a gql query object made up of resolvers, which
	// return scalar values from the fields: id, name, age
	// declare a graphql object that will be queried by the FE
	// this defines what fields are exposed in the graphql query
	astronautType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Astronaut",
		Description: "A set of astronauts",
		Fields: graphql.Fields{
			// each field on each type is backed by a function called a resolver
			// when a field is executed, the corresponding resolver is called to produce the value
			"id": &graphql.Field{
				//Non-null types enforce that their values are never null and can ensure
				// an error is raised if this ever occurs during a request.
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The id of the astronaut",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if a, ok := p.Source.(Astronaut); ok {
						return a.Id, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the astronaut",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if a, ok := p.Source.(Astronaut); ok {
						return a.Name, nil
					}
					return nil, nil
				},
			},
			"age": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The age of the astronaut",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if a, ok := p.Source.(Astronaut); ok {
						return a.Age, nil
					}
					return nil, nil
				},
			},
		},
	},
	)
	// queryType can be found at the top level of every graphql server - it represents
	// all possible entries into the graphql api
	// we call it query by convention, and allows only read operations
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"astronauts": &graphql.Field{
				Type: graphql.NewList(astronautType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetAllAstronauts(), nil
				},
			},
		},
	})

	// mutation to accept input type
	// Input types can't have fields that are other objects, only basic scalar types,
	// list types, and other input types.
	astronautInputType := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "AstronautInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The id of the astronaut",
			},
			"name": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the astronaut",
			},
			"age": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The age of the astronaut",
			},
		},
	})
	// mutationType can also be found at the top level, and is used for read-write operations
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "NotMutation",
		Fields: graphql.Fields{
			"addCrewMember": &graphql.Field{
				Type: graphql.NewList(astronautType),
				// the crewmember is added via argument
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Description: "An input with details about the new astronauts to add to the crew",
						Type:        graphql.NewNonNull(astronautInputType),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// from the argument parse from the interface (inputs) and pass into struct Astronaut
					var input = p.Args["input"].(map[string]interface{})
					// assert not nill and valued stored is Type
					a := Astronaut{
						Id:   input["id"].(int),
						Name: input["name"].(string),
						Age:  input["age"].(int),
					}
					return AddAstronaut(a), nil
				},
			},
		},
	})
	// schema definition
	// has query and mutation object types
	AstronautSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
}

//GetAllAstronauts return all the dummy data - it will be called by the query internally
func GetAllAstronauts() []Astronaut {
	// create slice
	astronauts := []Astronaut{}
	for _, a := range astronautData {
		astronauts = append(astronauts, a)
	}
	return astronauts
}

// AddAstronaut is called when addCrewMember is requested
// append the astronaut to the slice
func AddAstronaut(a Astronaut) []Astronaut {
	astronautData[len(astronautData)+1] = a
	return GetAllAstronauts()
}
