package main

import (
	_ "embed"

	"github.com/ether-edge/ether-edge/command/root"
	"github.com/ether-edge/ether-edge/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
