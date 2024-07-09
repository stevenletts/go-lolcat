package main

import (
	"bufio"
	"fmt"
	"math"
	"os/exec"
	"strings"
	"unicode/utf8"
)

func rgb(i int) (int, int, int) {
	f := 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func printColoredChar(output *strings.Builder, char string, j int) {
	r, g, b := rgb(j)
	fmt.Fprintf(output, "\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, char)
}

func main() {
	var output strings.Builder

	_, fortuneErr := exec.LookPath("fortune")
	_, cowsayErr := exec.LookPath("cowsay")

	if (fortuneErr != nil) && (cowsayErr != nil) {
		fmt.Println("to run this use brew install fortune and brew install cowsay")
		return
	}

	cmdStruct := exec.Command("fortune")

	out, err := cmdStruct.Output()
	if err != nil {
		fmt.Println(err)
	}

	cowStruct := exec.Command("cowsay", string(out))

	cowOut, err := cowStruct.Output()
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(cowOut)))

	j := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, w := 0, 0; i < len(line); i += w {
			runeValue, width := utf8.DecodeRuneInString(line[i:])
			w = width
			printColoredChar(&output, string(runeValue), j)
			j++
		}
		output.WriteString("\n")
	}
	fmt.Print(output.String())
}
