package main

import (
	"fmt"
	"github.com/codepawfect/jwt-generator/pkg"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 3 {
		fmt.Println("Error: Exactly three arguments are required")
		fmt.Println("Usage: jwt-generator companyTag purchaseOrganisation role")
		os.Exit(1)
	}

	companyTag, purchaseOrganisation, role := argsWithoutProg[0], argsWithoutProg[1], argsWithoutProg[2]
	jwt, err := pkg.CreateUnsignedJwt(companyTag, purchaseOrganisation, role)

	if err != nil {
		fmt.Println("error while generating jwt")
	}

	fmt.Println(jwt)
}
