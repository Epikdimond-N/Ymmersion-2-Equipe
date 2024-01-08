package routeur

import (
	"fmt"
	"net/http"
	controller "onepiece/controller"
	"os"
)

func InitServe() {
	http.HandleFunc("/", controller.DisplayHome)                 // template index
	http.HandleFunc("/persos/article", controller.DisplayPersos) //template article
	http.HandleFunc("/arcs/article", controller.DisplayArcs)
	http.HandleFunc("/events/article", controller.DisplayEvents)
	http.HandleFunc("/categories", controller.DisplayCategories) //template categories
	http.HandleFunc("/search", controller.DisplaySearch)         //template search
	http.HandleFunc("/admin", controller.DisplayAdmin)           //template admin
	http.HandleFunc("/addarticle", controller.DisplayAddArticle) //template addarticle
	http.HandleFunc("/error404", controller.Display404)          //template error404
	http.HandleFunc("/register", controller.RegisterHandler)
	http.HandleFunc("/confirmRegister", controller.ConfirmRegisterHandler)
	http.HandleFunc("/login", controller.LoginHandler)
	http.HandleFunc("/successLogin", controller.SuccessLoginHandler)
	http.HandleFunc("/logout", controller.LogoutHandler)
	http.HandleFunc("/changeLogin", controller.ChangeLoginHandler)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	fmt.Println("\nLien vers le site : http://localhost:8080 (CTRL+CLICK)")
	http.ListenAndServe("localhost:8080", nil)
}
