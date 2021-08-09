package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jihee102/explorer/blockchain"
	"github.com/jihee102/explorer/utils"
)

var port string

type uRLDescription struct {
	URL         uRL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type addBlockDescription struct {
	Description string
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

type uRL string

func (u uRL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

func (u uRLDescription) String() string {
	return "i'm URL description"
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []uRLDescription{
		{URL: uRL("/"), Method: "GET", Description: "See Documentation"},
		{URL: uRL("/blocks"), Method: "POST", Description: "Add a Block", Payload: "data:string"},
		{URL: uRL("/blocks/{height}"), Method: "GET", Description: "See a Block", Payload: "id:string"},
	}

	rw.Header().Add("Content-Type", "application/json")

	// byte, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s", byte)

	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())

	case "POST":
		var addBlockDescription addBlockDescription
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockDescription))
		blockchain.GetBlockchain().AddBlock(addBlockDescription.Description)
		rw.WriteHeader(http.StatusCreated)

	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["height"])
	utils.HandleErr(err)
	block, err := blockchain.GetBlockchain().GetBlock(id)

	if err == blockchain.ErrNotFound {
		json.NewEncoder(rw).Encode(errorResponse{fmt.Sprint(err)})

	} else {
		json.NewEncoder(rw).Encode(block)
	}

}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func Start(aPort int) {
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)
	port = fmt.Sprintf(":%d", aPort)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET")
	fmt.Printf("Listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
