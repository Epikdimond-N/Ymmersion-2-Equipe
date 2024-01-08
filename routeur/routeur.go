package routeur

import (
	"fmt"
	"net/http"
	"os"
	controller "onepiece/controller"
)

func InitServe() {
	http.HandleFunc("/", controller.DisplayHome)                 // template index
	http.HandleFunc("/persos/article", controller.DisplayPersos) //template article
	http.HandleFunc("/arcs/article", controller.DisplayArcs)
	http.HandleFunc("/events/article", controller.DisplayEvents)
	http.HandleFunc("/categories", controller.DisplayCategories)
	http.HandleFunc("/search", controller.DisplaySearch)         //template search
	http.HandleFunc("/admin", controller.DisplayAdmin)           //template admin
	http.HandleFunc("/addarticle", controller.DisplayAddArticle) //template addarticle
	http.HandleFunc("/error404", controller.Display404)        //template erreur 404

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/site-web/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	fmt.Println("\nLien vers le site : http://localhost:8080 (CTRL+CLICK)")
	http.ListenAndServe("localhost:8080", nil)
}
