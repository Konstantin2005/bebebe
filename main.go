package main

import (
	"bebebe/converter"
	"flag"
	"fmt"
	"os"
)

func main() {
	input := flag.String("input", "C:\\Users\\kisel\\GolandProjects\\bebebe\\resurse\\Input\\1Test.md", "путь к .md файлу (обязательно)")
	output := flag.String("output", "C:\\Users\\kisel\\GolandProjects\\bebebe\\resurse\\output\\1Test.html", "путь для сохранения .html (по умолчанию stdout)")
	help := flag.Bool("help", false, "показать справку")
	flag.Parse()

	if *help || *input == "" {
		flag.Usage()
		return
	}
	if *output == "" {
		flag.Usage()
		return
	}

	if err := converter.ConvertFile(*input, *output); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка:", err)
		os.Exit(1)
	}

}
