package main

import (
	"os"

	pinyin "github.com/hubenchang0515/go-pinyin"
)

func main() {
	var py = pinyin.NewPinyin()
	py.Parse(os.Args[1])
	println(py.String(" ", py.Tune))
}
