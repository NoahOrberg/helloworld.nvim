package main

import (
	"strings"

	"github.com/neovim/go-client/nvim"
	"github.com/neovim/go-client/nvim/plugin"
)

func hello(args []string) (string, error) {
	return "Hello " + strings.Join(args, " "), nil
}

func helloCmd(v *nvim.Nvim, args []string) error {
	return v.Command("echom 'hello " + args[0] + "'")
}

func main() {
	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleFunction(&plugin.FunctionOptions{Name: "Hello"}, hello)
		p.HandleCommand(&plugin.CommandOptions{Name: "HelloCmd", NArgs: "1"}, helloCmd)
		return nil
	})
}
