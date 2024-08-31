package main

import (
	"fmt"
	"github.com/codepawfect/jwt-generator/pkg"
	"os"
	"strings"
)

func main() {
	// Parse command-line arguments into a map
	claims := parseArgs()

	// Create JWT with the parsed claims
	jwtString, err := pkg.CreateUnsignedJwt(claims)
	if err != nil {
		fmt.Println("Error while generating JWT:", err)
		os.Exit(1)
	}

	fmt.Println(jwtString)
}

// parseArgs parses command-line arguments and returns a map of claims.
func parseArgs() map[string]interface{} {
	argsMap := make(map[string]interface{})
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "--") {
			key := strings.TrimPrefix(arg, "--")

			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "--") {
				nextArg := args[i+1]

				if strings.HasPrefix(nextArg, "-list") {
					// Handle list values
					if i+2 < len(args) && !strings.HasPrefix(args[i+2], "--") {
						listValues := strings.Split(args[i+2], ",")
						argsMap[key] = listValues
						i += 2 // Skip the list argument
					}
				} else {
					// Handle single values
					argsMap[key] = nextArg
					i++ // Skip the next argument since it's a value for the current key
				}
			} else {
				// No value provided
				argsMap[key] = ""
			}
		}
	}
	return argsMap
}
