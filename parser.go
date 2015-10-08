package acorn

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Parser tells how to parse files for code generation instructions
type Parser struct {
	Comment string
	Command string
}

// NewParser returns a new Acorn Parser
func NewParser(comment, command string) *Parser {
	return &Parser{Comment: comment, Command: command}
}

// Handler defines the callback interface
type Handler func([]string)

// Parse parses the reader, calling back with commands
func (p *Parser) Parse(reader io.Reader, callback Handler) error {
	var currentCommand []string

	endCommand := func() {
		if currentCommand != nil {
			callback(currentCommand)
		}
		currentCommand = nil
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, p.Comment) {
			endCommand()
			continue
		}

		line = strings.TrimSpace(strings.TrimPrefix(line, p.Comment))

		if currentCommand != nil {
			if line == "" {
				endCommand()
				continue
			}

			currentCommand = append(currentCommand, line)
			continue
		}

		if strings.HasPrefix(line, p.Command) {
			line := strings.TrimSpace(strings.TrimPrefix(line, p.Command))
			currentCommand = []string{line}
		}
	}

	endCommand()
	return scanner.Err()
}

// ParseFile opens a file for reading and parses it, calling back with commands
func (p *Parser) ParseFile(name string, callback Handler) error {
	file, err := os.Open(name)
	defer file.Close()

	if err != nil {
		return err
	}

	return p.Parse(file, callback)
}
