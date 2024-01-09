package main

import (
	routeur "onepiece/routeur"
	initTemplate "onepiece/temp"
	// One "onepiece/controller"
)

func main() {
	initTemplate.InitTemplate()
	routeur.InitServe()
	//One.GetArcs()
	//One.GetChar()
	//One.GetEvents()
	// One.UpdateChar()
}
