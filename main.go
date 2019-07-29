// +build go1.12

//go:generate sqlboiler -c ./build/sqlboiler.toml --no-context --add-global-variants mysql
//go:generate packr2
package main

import (
	"github.com/forsam-education/cerberus/cmd"
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql/driver"
	_ "github.com/volatiletech/sqlboiler/queries/qm"
)

func main() {
	cmd.Execute()
}
