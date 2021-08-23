package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

type balanceResponse struct {
	Address string `json:"address"`
	Balance int    `json:"balance"`
}

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
		{URL: uRL("/blocks"), Method: "GET", Description: "See Blocks"},
		{URL: uRL("/blocks/{hash}"), Method: "GET", Description: "See a Block", Payload: "hash:string"},
		{URL: uRL("/status"), Method: "GET", Description: "See the status of the blockchain"},
		{URL: uRL("/balance/{address}"), Method: "GET", Description: "Get TxOuts for an address"},
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
		json.NewEncoder(rw).Encode(blockchain.Blockchain().Blocks())

	case "POST":
		var addBlockDescription addBlockDescription
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockDescription))
		blockchain.Blockchain().AddBlock()
		rw.WriteHeader(http.StatusCreated)

	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]

	block, err := blockchain.FindBlock(hash)
	if err == blockchain.ErrNotFound {
		json.NewEncoder(rw).Encode(errorResponse{fmt.Sprint(err)})

	} else {
		json.NewEncoder(rw).Encode(block)
	}

}

func status(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(blockchain.Blockchain())
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func balance(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	total := r.URL.Query().Get("total")
	address := vars["address"]
	switch total {
	case "true":
		amount := blockchain.Blockchain().BalanceByAddress(address)
		utils.HandleErr(json.NewEncoder(rw).Encode(balanceResponse{address, amount}))

	default:
		utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.Blockchain().TxOutsByAddress(address)))
	}

}

func Start(aPort int) {
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)
	port = fmt.Sprintf(":%d", aPort)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{hash:[a-f0-9]+}", block).Methods("GET")
	router.HandleFunc("/status", status).Methods("GET")
	router.HandleFunc("/balance/{address}", balance).Methods("GET")
	fmt.Printf("Listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
