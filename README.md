# wc2 (the modern word-counter)

[![Go Report Card](https://goreportcard.com/badge/github.com/Loliver1224/wc2)](https://goreportcard.com/report/github.com/Loliver1224/wc2)
<a href="LICENSE" alt="MIT License"><img src="http://img.shields.io/badge/license-MIT-blue.svg?style=flat" /></a>

`wc2` is a modern wc(Word Count) command.

This is a modern take on that editor written in the Go-language for portability,
a modern word-counter with friendly output message.

GoでのCLI開発の練習がてら作ったwcの上位互換。
Emojiにも対応してるよ🤗

## Quick Start
Install `wc2` with a valid Go environment using `go get`:
```shell
> go get -u github.com/Loliver1224/wc2
```

## Usage
Basically, the same as the original.

```shell
> wc2 -h
Usage of main:
  -L    print the maximum display width
  -c    print the byte counts
  -l    print the line counts
  -m    print the character counts
  -w    print the word counts
```

In the original, the output was only numerical values, and it was difficult to understand.
Therefore, I added explanatory text for each number.

```shell
> wc2 -c -m -l -L -w alice.txt
byte counts:  603
character counts:  583
line counts:  11
maximum display width:  72
word counts: 112
```
