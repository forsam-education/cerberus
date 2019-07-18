//go:generate sqlboiler --no-context --add-global-variants mysql
package main

import (
	"github.com/forsam-education/kerberos/cmd"
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql/driver"
	_ "github.com/volatiletech/sqlboiler/queries/qm"
)

func main() {
	cmd.Execute()
}
