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

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func eraseConsoleLine() {
	fmt.Print("\033[A\033[2K")
}

func prefixMessage(prefix string, message string) string {
	currentTime := time.Now().Format("15:04:05")
	return fmt.Sprintf("%s  %s  ~  %s", prefix, currentTime, message)
}

func lineSeparator() string {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	splitted := strings.Split(string(out), " ")
	width, err := strconv.Atoi(strings.ReplaceAll(splitted[1], "\n", ""))
	if err != nil {
		return err.Error()
	}

	separator := ""
	for i := 0; i < width; i++ {
		separator = separator + "━"
	}

	separator = "\n" + separator + "\n"

	return separator
}
