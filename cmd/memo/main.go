package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/takymt/memo/internal/memo"
)

var revision = "dev"

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		return errors.New("usage: memo <description> | memo init | memo search <query> | memo open <query> | memo list --today|--week | memo capture | memo version")
	}

	configPath, err := memo.DefaultConfigPath()
	if err != nil {
		return err
	}

	switch args[0] {
	case "init":
		return runInit(configPath)
	case "search":
		if len(args) < 2 {
			return errors.New("usage: memo search <query>")
		}
		query := strings.Join(args[1:], " ")
		return runSearch(configPath, query)
	case "open":
		if len(args) < 2 {
			return errors.New("usage: memo open <query>")
		}
		query := strings.Join(args[1:], " ")
		return runOpen(configPath, query)
	case "version":
		fmt.Printf("revision: %s\n", revision)
		return nil
	default:
		description := strings.Join(args, " ")
		return runCreate(configPath, description)
	}
}

func runInit(configPath string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Printf("memo directory [%s]: ", cwd)
	var input string
	if _, err := fmt.Scanln(&input); err != nil {
		input = ""
	}
	if strings.TrimSpace(input) == "" {
		input = cwd
	}

	if err := os.MkdirAll(input, 0o755); err != nil {
		return err
	}

	cfg := memo.Config{MemoDir: input}
	if err := memo.SaveConfig(configPath, cfg); err != nil {
		return err
	}

	fmt.Printf("initialized: %s\n", configPath)
	return nil
}

func runCreate(configPath, description string) error {
	cfg, err := memo.LoadOrDefaultConfig(configPath)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(cfg.MemoDir, 0o755); err != nil {
		return err
	}

	fileName := memo.FileNameFromDescription(description)
	fullPath := filepath.Join(cfg.MemoDir, fileName)

	content := fmt.Sprintf("# %s\n\n", description)
	if err := os.WriteFile(fullPath, []byte(content), 0o644); err != nil {
		return err
	}

	fmt.Println(fullPath)
	return nil
}

func runSearch(configPath, query string) error {
	cfg, err := memo.LoadOrDefaultConfig(configPath)
	if err != nil {
		return err
	}

	matches, err := memo.SearchByFileName(cfg.MemoDir, query)
	if err != nil {
		return err
	}

	for _, match := range matches {
		fmt.Println(match)
	}
	return nil
}

func runOpen(configPath, query string) error {
	cfg, err := memo.LoadOrDefaultConfig(configPath)
	if err != nil {
		return err
	}

	matches, err := memo.SearchByFileName(cfg.MemoDir, query)
	if err != nil {
		return err
	}

	target := memo.BestMatch(matches)
	if target == "" {
		return errors.New("no memo matched")
	}

	editor := strings.TrimSpace(os.Getenv("EDITOR"))
	if editor == "" {
		editor = "vi"
	}

	command := exec.Command(editor, target)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}
