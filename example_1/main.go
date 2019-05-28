package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/cat-turner/gophers-in-space/example_1/schema"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
)

// GetAstronauts ... Function to resolve the graphql query
func GetAstronauts(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln("Error GetAstronauts", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error GetAstronauts", err)
	}

	var clientQuery map[string]interface{}

	fmt.Println("Received request GetAstronauts")
	if err := json.Unmarshal(body, &clientQuery); err != nil { // unmarshall body contents as a type query
		fmt.Println(err)
		fmt.Println("Error on Unmarshalling!!!")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error GetAstronauts unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	query := clientQuery["query"]
	variables := clientQuery["variables"]
	result := graphql.Do(graphql.Params{
		Schema:         schema.AstronautSchema,
		RequestString:  query.(string),
		VariableValues: variables.(map[string]interface{}),
	})
	json.NewEncoder(w).Encode(result)
	return
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	var handler http.Handler
	handler = http.HandlerFunc(GetAstronauts)
	// only one entry point defines all capabilities of a graphql server
	router.Methods("POST").Path("/").Name("graphql").Handler(handler)
	fmt.Println("Now server is running on port 8090")
	// launch server
	log.Fatal(http.ListenAndServe(":8090", router))

}
