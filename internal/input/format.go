package input

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	GreenColor      = "\033[32m"
	BlueColor       = "\033[34m"
	ResetColor      = "\033[0m"
	separatorSymbol = "-"
)

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func eraseLinesCount(count int) {
	for i := 0; i < count; i++ {
		fmt.Print("\033[A\033[2K")
	}
}

func eraseTextLines(text string) {
	width := terminalWidth()
	usedLines := (len(text) / width) + 1
	eraseLinesCount(usedLines)
}

func prefixMessage(prefix string, message string, color string) string {
	currentTime := time.Now().Format("15:04:05")
	return fmt.Sprintf(color+"%s  %s  ~ "+ResetColor+" %s", prefix, currentTime, message)
}

func formatUserMessage(prefix string, message string) {
	input := prefixMessage(prefix, message, BlueColor)
	eraseTextLines(input)
	fmt.Println(input)
}

func lineSeparator() string {
	width := terminalWidth()
	line := ""
	for i := 0; i < width; i++ {
		line += separatorSymbol
	}
	return "\n" + line + "\n"
}

func centerMessage(message string) string {
	width := terminalWidth()
	line := ""
	paddingLength := float32(width-len(message)) / 2
	for i := 0; i < int(paddingLength); i++ {
		line += separatorSymbol
	}
	line += message
	for i := 0; i < int(paddingLength); i++ {
		line += separatorSymbol
	}
	if (width-len(message))%2 != 0 {
		line += separatorSymbol
	}
	return line
}

func addSpaces(str1 string, str2 string) (string, string) {
	diff := len(str1) - len(str2)
	if diff > 0 {
		return str1, str2 + strings.Repeat(" ", diff)
	}
	if diff < 0 {
		return str1 + strings.Repeat(" ", -diff), str2
	}

	return str1, str2
}

func terminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	splitted := strings.Split(string(out), " ")
	width, err := strconv.Atoi(strings.ReplaceAll(splitted[1], "\n", ""))
	if err != nil {
		log.Fatal(err)
	}
	return width
}
