package repl

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type REPL struct{}

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorBold   = "\033[1m"
)

func NewREPL() *REPL {
	return &REPL{}
}

func (r *REPL) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	printBanner()

	for {
		fmt.Print(ColorCyan + "❯ " + ColorReset)
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		if strings.HasPrefix(input, ".") {
			handleCommand(input)
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, ColorRed+"Error reading input:"+ColorReset, err)
	}
}

func printBanner() {
	fmt.Println(ColorPurple + ColorBold)
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║                                        ║")
	fmt.Println("║       Welcome to the AdamScript        ║")
	fmt.Println("║                 REPL                   ║")
	fmt.Println("║                                        ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println(ColorReset)
	fmt.Println(ColorYellow + "Type " + ColorCyan + ".help" + ColorYellow + " for available commands" + ColorReset)
	fmt.Println()
}

func handleCommand(cmd string) {
	switch cmd {
	case ".help":
		printHelp()
	case ".exit":
		os.Exit(0)
	case ".clear":
		clearScreen()
	default:
		fmt.Printf(ColorRed+"Unknown command: "+ColorWhite+"%s\n"+ColorReset, cmd)
		fmt.Println(ColorYellow + "Type " + ColorCyan + ".help" + ColorYellow + " for available commands" + ColorReset)
	}
}

func printHelp() {
	fmt.Println(ColorBlue + ColorBold + "\nAvailable Commands:" + ColorReset)
	fmt.Println(ColorCyan + "  .help  " + ColorWhite + "- Show this help message" + ColorReset)
	fmt.Println(ColorCyan + "  .exit  " + ColorWhite + "- Exit the REPL" + ColorReset)
	fmt.Println(ColorCyan + "  .clear " + ColorWhite + "- Clear the screen" + ColorReset)
	fmt.Println()
}

func clearScreen() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()

	printBanner()
}
