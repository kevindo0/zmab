package main

import (
	"zmab/ab"
)

func main() {
	abs := ab.Load("./config.json")
	opt := abs.Init()
	ab.Start(opt)
}
