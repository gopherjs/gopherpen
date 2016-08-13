//go:generate go run assets_gen.go assets.go

// Gopherpen is a project template that will let you easily get started with GopherJS
// for building a web app. It includes some simple HTML, CSS, and Go code for the frontend.
// Make some changes, and refresh in browser to see results. When there are errors in your
// frontend Go code, they will show up in the browser console.
//
// Once you're done making changes, you can easily create a fully self-contained static
// production binary.
package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/shurcooL/httpfs/html/vfstemplate"
	"github.com/shurcooL/httpgzip"
)

var httpFlag = flag.String("http", ":8080", "Listen for HTTP connections on this address.")

func loadTemplates() (*template.Template, error) {
	t := template.New("").Funcs(template.FuncMap{})
	t, err := vfstemplate.ParseGlob(assets, t, "/assets/*.tmpl")
	return t, err
}

func mainHandler(w http.ResponseWriter, req *http.Request) {
	t, err := loadTemplates()
	if err != nil {
		log.Println("loadTemplates:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = struct {
		Animals string
	}{
		Animals: "gophers",
	}

	err = t.ExecuteTemplate(w, "index.html.tmpl", data)
	if err != nil {
		log.Println("t.Execute:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	flag.Parse()

	http.HandleFunc("/", mainHandler)
	http.Handle("/assets/", httpgzip.FileServer(assets, httpgzip.FileServerOptions{ServeError: httpgzip.Detailed}))
	http.Handle("/favicon.ico", http.NotFoundHandler())

	printServingAt(*httpFlag)
	err := http.ListenAndServe(*httpFlag, nil)
	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}

func printServingAt(addr string) {
	hostPort := addr
	if strings.HasPrefix(hostPort, ":") {
		hostPort = "localhost" + hostPort
	}
	fmt.Printf("serving at http://%s/\n", hostPort)
}
