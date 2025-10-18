package main

import (
	"bebebe/converter"
	"flag"
	"fmt"
	"os"
)

func main() {
	input := flag.String("input", "C:\\Users\\kisel\\GolandProjects\\bebebe\\resurse\\Input\\1Test.md", "путь к .md файлу (обязательно)")
	output := flag.String("output", "", "путь для сохранения .html(по умолчанию stdout)")
	help := flag.Bool("help", false, "показать справку")
	Paterns := "C:\\Users\\kisel\\GolandProjects\\bebebe\\resurse\\pattern\\Ru_langveg.html"
	flag.Parse()

	if *help || *input == "" {
		flag.Usage()
		return
	}

	if err := converter.ConvertFile(*input, *output, Paterns); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка:", err)
		os.Exit(1)
	}

}
