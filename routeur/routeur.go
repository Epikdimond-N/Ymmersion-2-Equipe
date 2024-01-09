package routeur

import (
	"fmt"
	"net/http"
	controller "onepiece/controller"
	"os"
)

func InitServe() {
	http.HandleFunc("/", controller.DisplayHome) // template index
	http.HandleFunc("/char", controller.DisplayChar)
	http.HandleFunc("/search", controller.HandleSearch)          //template search
	http.HandleFunc("/persos/article", controller.DisplayPersos) //template article
	http.HandleFunc("/arcs/article", controller.DisplayArcs)
	http.HandleFunc("/events/article", controller.DisplayEvents)
	http.HandleFunc("/categories", controller.DisplayCategories)           //template categories
	http.HandleFunc("/admin", controller.DisplayAdmin)                     //template admin
	http.HandleFunc("/addarticle", controller.DisplayAddArticle)           //template addarticle
	http.HandleFunc("/error404", controller.Display404)                    //template error404
	http.HandleFunc("/register", controller.RegisterHandler)               //template enregister le compte
	http.HandleFunc("/confirmRegister", controller.ConfirmRegisterHandler) // ecrit dans le json user.json si les donné son correcte
	http.HandleFunc("/login", controller.LoginHandler)                     // possibilité de se log, page de redirection si logged = false
	http.HandleFunc("/successLogin", controller.SuccessLoginHandler)       // verifé les donnés entrée et les data du json user.json
	http.HandleFunc("/logout", controller.LogoutHandler)                   // reset les variable login / username / password / admin
	http.HandleFunc("/changeLogin", controller.ChangeLoginHandler)         // ouvre la possibilité de changer de username

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	fmt.Println("\nLien vers le site : http://localhost:8080 (CTRL+CLICK)")
	http.ListenAndServe("localhost:8080", nil)
}
