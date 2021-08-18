package main

import (
	"github.com/jihee102/explorer/cli"
	"github.com/jihee102/explorer/db"
)

// const port string = ":4000"

// type URLDescription struct {
// 	URL         URL    `json:"url"`
// 	Method      string `json:"method"`
// 	Description string `json:"description"`
// 	Payload     string `json:"payload,omitempty"`
// }

// type AddBlockDescription struct {
// 	Description string
// }

// type URL string

// func (u URL) MarshalText() ([]byte, error) {
// 	url := fmt.Sprintf("http://localhost%s%s", port, u)
// 	return []byte(url), nil
// }

// func (u URLDescription) String() string {
// 	return "i'm URL description"
// }

// func documentation(rw http.ResponseWriter, r *http.Request) {
// 	data := []URLDescription{
// 		{URL: URL("/"), Method: "GET", Description: "See Documentation"},
// 		{URL: URL("/blocks"), Method: "POST", Description: "Add a Block", Payload: "data:string"},
// 		{URL: URL("/blocks/{id}"), Method: "GET", Description: "See a Block", Payload: "id:string"},
// 	}

// 	rw.Header().Add("Content-Type", "application/json")

// 	// byte, err := json.Marshal(data)
// 	// utils.HandleErr(err)
// 	// fmt.Fprintf(rw, "%s", byte)

// 	json.NewEncoder(rw).Encode(data)
// }

// func blocks(rw http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "GET":
// 		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())

// 	case "POST":
// 		var addBlockDescription AddBlockDescription
// 		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockDescription))
// 		blockchain.GetBlockchain().AddBlock(addBlockDescription.Description)
// 		rw.WriteHeader(http.StatusCreated)

// 	}
// }

func main() {
	// go rest.Start(4000)
	// explorer.Start(5000)

	// cli.Start()

	// http.HandleFunc("/", documentation)
	// http.HandleFunc("/blocks", blocks)
	// fmt.Printf("Listening on http://localhost%s", port)
	// log.Fatal(http.ListenAndServe(port, nil))

	// blockchain.Blockchain().AddBlock("seconde")
	// blockchain.Blockchain().AddBlock("third")

	defer db.Close()
	cli.Start()

}
