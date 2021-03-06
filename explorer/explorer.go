package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/jihee102/explorer/blockchain"
)

var port string

var templates *template.Template

const (
	templateDir string = "explorer/templates/"
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(rw, "Hello from home!")
	// tmpl := template.Must(template.ParseFiles("templates/pages/home.gohtml"))

	data := homeData{"cocoin home", blockchain.Blockchain().Blocks()}
	// tmpl.Execute(rw, data)
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()

		blockchain.Blockchain().AddBlock()
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)

	}

}

func Start(aPort int) {
	handler := http.NewServeMux()
	port = fmt.Sprintf(":%d", aPort)
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)

	fmt.Printf("Listening on http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, handler))
}
