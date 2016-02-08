package tldr

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func GenPath(tldrpath, pf, cmd string) (path string) {
	path += tldrpath + "/tldr/pages/" + pf + "/" + cmd + ".md"
	return path
}

func PrintMd(path string) (output string) {
	var err error

	fp, err := os.Open(path)
	if err != nil {
		return "File could not be found\n"
	}
	defer fp.Close()
	reader := bufio.NewReaderSize(fp, 4096)
	for {
		line, _, err := reader.ReadLine()
		text := string(line)
		if len(text) == 0 {
		} else if strings.HasPrefix(text, "#") {
			output += backGreenPrint(text[2:])
		} else if strings.HasPrefix(text, ">") {
			output += " " + readPrint(text[2:])
		} else if strings.HasPrefix(text, "-") {
			output += " " + whitePrint(text[2:])
		} else if strings.HasPrefix(text, "`") {
			output += "  " + greenPrint(text[1:len(text)-1])
			output += "\n"
		} else {
			output += whitePrint(text)
		}

		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	return output
}

func readPrint(text string) (output string) {
	output += fmt.Sprintf("\x1b[31m%s\x1b[0m\n", text)
	return output
}

func bluePrint(text string) (output string) {
	output += fmt.Sprintf("\x1b[34m%s\x1b[0m\n", text)
	return output
}

func greenPrint(text string) (output string) {
	output += fmt.Sprintf("\x1b[32m%s\x1b[0m\n", text)
	return output
}

func backGreenPrint(text string) (output string) {
	output += fmt.Sprintf("\x1b[44m%s\x1b[0m\n", text)
	return output
}

func whitePrint(text string) (output string) {
	output += fmt.Sprintf("\x1b[37m%s\x1b[0m\n", text)
	return output
}

func yellowPrint(text string) (output string) {
	output += fmt.Sprintf("\x1b[33m%s\x1b[0m\n", text)
	return output
}
