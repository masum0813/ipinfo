/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import "github.com/masum0813/ipinfo/cmd"

var (
	// VERSION is set during build
	VERSION = "0.0.8"
)

func main() {
	cmd.Execute(VERSION)
}
