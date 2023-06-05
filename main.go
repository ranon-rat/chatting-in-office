package main

import (
	"fmt"

	"github.com/ranon-rat/chatting-in-office/router"
)

func main() {
	fmt.Println(`
                                 
	 _           _       _       _   
	| |_ ___ _ _| |_ ___| |_ ___| |_ 
	| . |  _| | |   |  _|   | .'|  _|
	|___|_| |___|_|_|___|_|_|__,|_|  
									 
	
	

 simple chatting tool for using it in the office
 

	`)

	router.SetupRoutes()
}
