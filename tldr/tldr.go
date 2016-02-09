package tldr

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func GenPath(tldrpath, cmd string) (path string) {
	path += tldrpath + "/" + cmd + ".md"
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
			output += backGreenPrint(text[2:]) + "\n"
		} else if strings.HasPrefix(text, ">") {
			output += " " + readPrint(text[2:]) + "\n"
		} else if strings.HasPrefix(text, "-") {
			output += " " + whitePrint(text[2:]) + "\n"
		} else if strings.HasPrefix(text, "`") {
			output += "  " + codeHightLight(text[1:len(text)-1]) + "\n"
			output += "\n"
		} else {
			output += whitePrint(text) + "\n"
		}

		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	return output
}

func codeHightLight(text string) (output string) {
	codes := strings.Split(text, "}}")
	var code string
	for _, code = range codes {
		for i := 0; i < len(code); i++ {
			if strings.HasPrefix(code[i:], "{{") {
				output += greenPrint(code[:i])
				output += whitePrint(code[i+2:])
			}
		}
	}
	output += greenPrint(code)
	if strings.Compare(output, "") == 0 {
		output = text
	}
	return output
}

func readPrint(text string) (output string) {
	output += fmt.Sprintf("\x1b[31m%s\x1b[0m", text)
	return output
}

func bluePrint(text string) (output string) {
	output += fmt.Sprintf("\x1b[34m%s\x1b[0m", text)
	return output
}

func greenPrint(text string) (output string) {
	output += fmt.Sprintf("\x1b[32m%s\x1b[0m", text)
	return output
}

func backGreenPrint(text string) (output string) {
	output += fmt.Sprintf("\x1b[44m%s\x1b[0m", text)
	return output
}

func whitePrint(text string) (output string) {
	output += fmt.Sprintf("\x1b[37m%s\x1b[0m", text)
	return output
}

func yellowPrint(text string) (output string) {
	output += fmt.Sprintf("\x1b[33m%s\x1b[0m", text)
	return output
}
