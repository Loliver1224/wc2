package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"unicode"
)

type Counter struct {
	bytes         int
	chars         int
	lines         int
	maxLineLength int
	words         int
}

var (
	// 改行コード問題対策
	isSetPlatform bool
	isMacOrWin    bool
	isUnix        bool
)

func changeMax(a, b *int) {
	if *a < *b {
		*a = *b
	}
}

func isBreak(ch rune) bool {
	// 最初に出てくる改行文字で改行コードを判定
	if !isSetPlatform {
		if ch == '\n' {
			isUnix = true
			isSetPlatform = true
			fmt.Println("setLF")
			return true
		} else if ch == '\r' {
			isMacOrWin = true
			isSetPlatform = true
			fmt.Println("setCR")
			return true
		}
	} else {
		if isUnix && ch == '\n' {
			return true
		} else if isMacOrWin && ch == '\r' {
			return true
		}
	}
	return false
}

func count(file *os.File) (cnt *Counter) {
	scanner := bufio.NewReader(file)
	cnt = new(Counter)

	cnt.maxLineLength = -1
	lineWidth := 0
	isInWord := false
	prevCh := '^'

	for {
		ch, size, err := scanner.ReadRune()
		if err != nil {
			break
		}

		cnt.bytes += size
		if isBreak(ch) {
			cnt.lines++
			if !isBreak(prevCh) && isInWord {
				cnt.words++
			}
			isInWord = false
			changeMax(&cnt.maxLineLength, &lineWidth)
			lineWidth = 0
		} else if unicode.IsSpace(ch) && isInWord {
			cnt.words++
			isInWord = false
		} else if !unicode.IsSpace(ch) && unicode.IsGraphic(ch) {
			isInWord = true
		}
		if !isBreak(ch) && unicode.IsGraphic(ch) {
			cnt.chars++
			lineWidth++
		}
		prevCh = ch
	}
	// 入力が空でなければ，1行目分をカウント
	if cnt.bytes > 0 {
		cnt.lines++
	}

	// 最終行の処理
	// 最大長判定
	changeMax(&cnt.maxLineLength, &lineWidth)
	// 直前の文字が改行でなければwords++
	if !isBreak(prevCh) {
		cnt.words++
	}
	return
}

func isFileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func main() {
	var (
		// flag definition
		bytes         = flag.Bool("c", false, "print the byte counts")
		chars         = flag.Bool("m", false, "print the character counts")
		lines         = flag.Bool("l", false, "print the line counts")
		maxLineLength = flag.Bool("L", false, "print the maximum display width")
		words         = flag.Bool("w", false, "print the word counts")

		// for file open
		file *os.File
		err  error
	)

	flag.Parse()

	// set default flag
	if flag.NFlag() == 0 {
		*bytes = true
		*lines = true
		*words = true
	}

	// open file
	if arglen := len(flag.Args()); arglen == 1 {
		filename := flag.Arg(0)
		if isFileExists(filename) {
			if file, err = os.Open(filename); err != nil {
				fmt.Fprintf(os.Stderr, "fileCannotOpenError: file %s is cannot opened.\n", filename)
			}
			defer file.Close()
		} else {
			fmt.Fprintf(os.Stderr, "fileNotFoundError: file %s is not found.\n", filename)
			os.Exit(1)
		}
	} else if arglen == 0 {
		// set os.Stdin to file pointer
		file = os.Stdin
		defer file.Close()
	}

	result := count(file)

	// output
	if *bytes {
		fmt.Println("byte counts: ", result.bytes)
	}
	if *chars {
		fmt.Println("character counts: ", result.chars)
	}
	if *lines {
		fmt.Println("line counts: ", result.lines)
	}
	if *maxLineLength {
		fmt.Println("maximum display width: ", result.maxLineLength)
	}
	if *words {
		fmt.Println("word counts:", result.words)
	}
}
