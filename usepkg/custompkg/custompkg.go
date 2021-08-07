package custompkg

import (
	"fmt"
	"goproject/usepkg/exinit"
)

func PrintCustom() {
	fmt.Println("This is custom package!")
	exinit.PrintD()
}

func Printcustom2() {
	fmt.Println("This is custom2 package!")
}
