package main

import (
	"onepiece/routeur"
	initTemplate "onepiece/temp"
)

func main() {
	initTemplate.InitTemplate()
	routeur.InitServe()
	//One.GetArcs()
	//One.GetChar()
	//One.GetEvents()
	//Two.UpdateChar()
}
