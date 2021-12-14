/*
 * Loan Calculator
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"log"
	"net/http"
//	"fmt"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	sw "LoanCalculator/go"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	//handler := http.HandlerFunc(sw.Index)
	//fmt.Println("Hello") //Test print statement to see if server works
 
	log.Fatal(http.ListenAndServe(":8080", router))
}
