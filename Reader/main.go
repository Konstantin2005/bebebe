package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
	"unicode"
)

type StatisticChar struct {
	ToLowerChar     map[rune]int
	ToUpperChar     map[rune]int
	BestToLowerChar rune
	BestToUpperChar rune
	TotalChar       rune
}
type StatisticString struct {
	Statistic StatisticChar
	TotalWord int
	TotalLine int
}

var wordRE = regexp.MustCompile(`[\p{L}\p{N}]+`)

func (S *StatisticString) MainFail(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v\n", err)
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Ошибка закрытия файла: %v\n", err)
		}
	}()

	S.Statistic.ToUpperChar = make(map[rune]int)
	S.Statistic.ToLowerChar = make(map[rune]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		S.TotalLine++
		S.mostFrequentLowerChar(line)
		S.wordCountStrict(line)
	}
	S.BestToUpperChar()
	S.BestToLowerChar()

	if S.Statistic.ToUpperChar[S.Statistic.BestToUpperChar] > S.Statistic.ToLowerChar[S.Statistic.BestToLowerChar] {
		S.Statistic.TotalChar = S.Statistic.BestToUpperChar
	} else {
		S.Statistic.TotalChar = S.Statistic.BestToLowerChar
	}
}

func (S *StatisticString) wordCountStrict(s string) {
	S.TotalWord += len(wordRE.FindAllString(s, -1))
}

func (S *StatisticString) mostFrequentLowerChar(s string) {
	if s == " " {
		return
	}
	for _, r := range s {
		if unicode.IsLetter(r) {
			if unicode.IsLower(r) {
				S.Statistic.ToLowerChar[r]++
			} else {
				S.Statistic.ToUpperChar[r]++
			}
		}
	}
}

func (S *StatisticString) BestToUpperChar() {
	var CountBestChar int
	for Char, value := range S.Statistic.ToUpperChar {
		if value > CountBestChar {
			S.Statistic.BestToUpperChar = Char
			CountBestChar = value
		}
	}
}

func (S *StatisticString) BestToLowerChar() {
	var CountBestChar int
	for Char, value := range S.Statistic.ToLowerChar {
		if value > CountBestChar {
			S.Statistic.BestToLowerChar = Char
			CountBestChar = value
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("не коректрый путь")
	}
	path := os.Args[1]
	S := StatisticString{}
	start := time.Now()
	S.MainFail(path)
	elapsed := time.Since(start)

	fmt.Printf("Общея статистика в файле %v\n", path)
	fmt.Printf("Общие количество строк %d \n", S.TotalLine)
	fmt.Printf("Общие количество Слов %d \n", S.TotalWord)
	fmt.Printf("Самый популярный символ %s в нижнем регистре\n", string(S.Statistic.BestToLowerChar))
	fmt.Printf("Самый популярный символ %s в Верхнем регистре\n", string(S.Statistic.BestToUpperChar))
	fmt.Printf("Самый популярный символ %s в общем\n\n", string(S.Statistic.TotalChar))

	fmt.Printf("Время выполнения: %s\n\n", elapsed)

}
