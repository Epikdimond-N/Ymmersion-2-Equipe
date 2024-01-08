package main

import (
	// One "onepiece/go"
	routeur "onepiece/routeur"
	initTemplate "onepiece/temp"
)

func main() {
	// One.GetChar()
	initTemplate.InitTemplate()
	routeur.InitServe()
}
