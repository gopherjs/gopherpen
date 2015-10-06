package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/shurcooL/go/gzip_file_server"
	"github.com/shurcooL/httpfs/html/vfstemplate"
)

var httpFlag = flag.String("http", ":8080", "Listen for HTTP connections on this address.")

var t *template.Template

func loadTemplates() error {
	var err error
	t = template.New("").Funcs(template.FuncMap{})
	t, err = vfstemplate.ParseGlob(assets, t, "/assets/*.tmpl")
	return err
}

func mainHandler(w http.ResponseWriter, req *http.Request) {
	if !production {
		if err := loadTemplates(); err != nil {
			log.Println("loadTemplates:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	var data = struct {
		Animals string
	}{
		Animals: "gophers",
	}

	err := t.ExecuteTemplate(w, "index.html.tmpl", data)
	if err != nil {
		log.Println("t.Execute:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	flag.Parse()

	if production {
		err := loadTemplates()
		if err != nil {
			log.Fatalln("loadTemplates:", err)
		}
	}

	http.HandleFunc("/", mainHandler)
	http.Handle("/assets/", gzip_file_server.New(assets))
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
