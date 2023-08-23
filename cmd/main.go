package main

import (
	"github.com/fatih/color"
)

var (
	progressMessage = color.GreenString("==>")
	redColor        = color.RedString("this is" + "red string")
	buldColor       = color.CyanString("ths is thire string")
)

func main() {
	c := color.New(color.FgCyan)
	c.Println("打印青色文字")

	c.DisableColor()
	c.Println("这会无颜色打印")

	c.EnableColor()
	c.Println("这再次打印青色...")
}
