package main

import (
	"fmt"
	"regexp"
)

func checkArgs(args []string, needArgs int) string {
	if len(args) != needArgs {
		return fmt.Sprintf("expect %v many arguments, got %v", needArgs, len(args))
	}
	for _, s := range args {
		match, _ := regexp.MatchString("[a-zA-Z0-9_']", s)
		if !match {
			return fmt.Sprintf("invalid argument %v", s)
		}
	}
	return ""
}
