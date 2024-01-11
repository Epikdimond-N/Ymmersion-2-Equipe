package routeur

import (
	"fmt"
	"net/http"
	controller "onepiece/controller"
)

func InitServe() {
	http.HandleFunc("/", controller.NotFoundHandler)                              // catch-all route for any unspecified paths and display page 404 << working
	http.HandleFunc("/Home", controller.DisplayHome)                              // display root / page d'acceuil << working, 10 rondom article
	http.HandleFunc("/Persos", controller.DisplayPerso)                           // display given char ID << working
	http.HandleFunc("/Arcs", controller.DisplayArc)                               // display given arc ID << working
	http.HandleFunc("/EventsOnePiece", controller.DisplayEvent)                   // display given event ID << working
	http.HandleFunc("/search", controller.HandleSearch)                           // host la barre de recherche et la page d'article trouvé << working
	http.HandleFunc("/persos/article", controller.DisplayPersos)                  // display all chars articles << need update
	http.HandleFunc("/arcs/article", controller.DisplayArcs)                      // display all arcs articles << need update
	http.HandleFunc("/events/article", controller.DisplayEvents)                  // display all events article << need update
	http.HandleFunc("/categories", controller.DisplayCategories)                  // display categorie choice page << working
	http.HandleFunc("/admin", controller.DisplayAdmin)                            // display admin page << working
	http.HandleFunc("/admin/delete", controller.DisplayAdminDelete)               // display the form for delete << working
	http.HandleFunc("/admin/delete/gestion", controller.DeleteHandler)            // delete the selecte post
	http.HandleFunc("/admin/deleteThis", controller.DisplayAdminDeleteConf)       // display the article for confirmation for deleting << need update
	http.HandleFunc("/register", controller.RegisterHandler)                      // creation de compte << working
	http.HandleFunc("/confirmRegister", controller.ConfirmRegisterHandler)        // ecrit dans le json user.json  // gestion de creation de compte << working
	http.HandleFunc("/login", controller.LoginHandler)                            // possibilité de se log, page de redirection si logged = false << working
	http.HandleFunc("/successLogin", controller.SuccessLoginHandler)              // verifé les donnés entrée et les comparer au data du json user.json << working
	http.HandleFunc("/logout", controller.LogoutHandler)                          // reset les variable login / username / password / admin << need update
	http.HandleFunc("/changeLogin", controller.ChangeLoginHandler)                // ouvre la possibilité de changer de username << need update
	http.HandleFunc("/admin/newChar", controller.NewCharHandler)                  // page admin d'ajout de char << img of 10mb is ok ? << need update
	http.HandleFunc("/admin/newArc", controller.NewArcHandler)                    // page admin d'ajout d'arc << need update
	http.HandleFunc("/admin/newEvent", controller.NewEventHandler)                // page admin d'ajout d'event << working
	http.HandleFunc("/admin/newChar/gestion", controller.GestionNewPersosHandler) //system de gestion d'ajout char << need update
	http.HandleFunc("/admin/newArc/gestion", controller.GestionNewArcHandler)     //system de gestion d'ajout arc << need update
	http.HandleFunc("/admin/newEvent/gestion", controller.GestionNewEventHandler) //system de gestion d'ajout event << need update

	controller.ChargeImage()
	fmt.Println("\nLien vers le site : http://localhost:8080/Home (CTRL+CLICK)")
	http.ListenAndServe("localhost:8080", nil)
}
