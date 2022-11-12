//go:build !test
// +build !test

package main

import (
	"flag"
	"log"

	builder "godxvincent.com/go-patterns/patterns/creational/builder/execution"
	singleton "godxvincent.com/go-patterns/patterns/creational/singleton/execution"
)

type CallingPatterns map[string]func()

var patterns CallingPatterns

func main() {

	// variables declaration
	var patternSelected string

	// initLoadFunctions()
	// flags declaration using flag package
	flag.StringVar(&patternSelected, "pattern", "singleton", "Please set up the patter required. Default is singleton")

	// after declaring flags we need to call it
	flag.Parse()

	// check if cli params match
	if patternSelected != "" {
		log.Println("*************************************Starting Demo Pattern*************************************")
		log.Println("This is the pattern selected by the user =>", patternSelected)
		functionSelected := patterns[patternSelected]
		functionSelected()
		log.Println("*************************************Ending Demo Pattern*************************************")
	}

}

func init() {
	patterns = make(map[string]func())
	patterns["singleton"] = singleton.Execute
	patterns["builder"] = builder.Execute
}
