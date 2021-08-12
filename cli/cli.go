package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/jihee102/explorer/explorer"
	"github.com/jihee102/explorer/rest"
)

func usage() {
	fmt.Printf("Welcome to cocoin\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("-port:		Set the PORT of the server\n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest' ")
	// os.Exit(0)
	runtime.Goexit()
}
func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest' ")
	htmlPort := flag.Int("htmlPort", 4000, "Set port of the html server")
	restPort := flag.Int("restPort", 5000, "Set port of the REST server")

	flag.Parse()

	switch *mode {
	case "rest":
		//start rest api
		rest.Start(*port)

	case "html":
		// start explorer
		explorer.Start(*port)
	case "both":
		go rest.Start(*restPort)
		explorer.Start(*htmlPort)
	default:
		usage()
	}

	fmt.Println(*port, *mode)
}
