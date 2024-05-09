/*
Copyright © 2024 HUBERT SZARWINSKI <SZARHUBERT@GMAIL.COM>
*/
package main

import (
	"example.com/cli/cmd"
	"example.com/cli/db"
)

func main() {
	db.Run()
	cmd.Execute()
}
