package maxima

import (
	"bufio"
	"io"
	"log"
	"os/exec"
	"regexp"
)

var inprompt = regexp.MustCompile(`\(%i`)
var outprompt = regexp.MustCompile(`\(%o\d+\) `)
var newline = regexp.MustCompile(`\n`)

type Maxima struct {
	Stderr io.ReadCloser
	Stdin  io.WriteCloser
	Stdout io.ReadCloser
	Reader *bufio.Reader
}

// NewMaxima starts an instance of maxima, reads through the stdout buffer until ready for input and returns StderrPipe, StdinPipe, StdoutPipe and bufio.Reader of StdoutPipe
func NewMaxima() Maxima {
	maxima := Maxima{}
	var err error

	cmd := exec.Command("maxima")

	maxima.Stderr, err = cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	// maxima.Stderr = stderr
	maxima.Stdin, err = cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	maxima.Stdout, err = cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	maxima.Reader = bufio.NewReader(maxima.Stdout)

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	for {
		pk, _ := maxima.Reader.Peek(3)
		if inprompt.Match(pk) {
			maxima.Reader.ReadString(' ')
			break
		}
		maxima.Reader.ReadString('\n')
	}

	return maxima
}

func (m Maxima) Declare(vari, val string) (out string) {
	cmd := "declare(" + vari + ", " + val + ");\n"
	io.WriteString(m.Stdin, cmd)

	for {
		pk, _ := m.Reader.Peek(3)
		if inprompt.Match(pk) {
			m.Reader.ReadString(' ')
			break
		}
		ln, _ := m.Reader.ReadString('\n')
		out += ln
	}

	out = outprompt.ReplaceAllString(out, "")
	out = newline.ReplaceAllString(out, "")

	return out
}

func (m Maxima) Eval(val string) (out string) {
	cmd := val + ";\n"
	io.WriteString(m.Stdin, cmd)

	for {
		pk, _ := m.Reader.Peek(3)
		if inprompt.Match(pk) {
			m.Reader.ReadString(' ')
			break
		}
		ln, _ := m.Reader.ReadString('\n')
		out += ln
	}

	out = outprompt.ReplaceAllString(out, "")
	out = newline.ReplaceAllString(out, "")

	return out
}

func (m Maxima) Func(name, vari, val string) (out string) {
	cmd := name + "(" + vari + ") := " + val + ";\n"
	io.WriteString(m.Stdin, cmd)

	for {
		pk, _ := m.Reader.Peek(3)
		if inprompt.Match(pk) {
			m.Reader.ReadString(' ')
			break
		}
		ln, _ := m.Reader.ReadString('\n')
		out += ln
	}

	out = outprompt.ReplaceAllString(out, "")
	out = newline.ReplaceAllString(out, "")

	return out
}

func (m Maxima) Set(vari, val string) (out string) {
	cmd := vari + ": " + val + ";\n"
	io.WriteString(m.Stdin, cmd)

	for {
		pk, _ := m.Reader.Peek(3)
		if inprompt.Match(pk) {
			m.Reader.ReadString(' ')
			break
		}
		ln, _ := m.Reader.ReadString('\n')
		out += ln
	}

	out = outprompt.ReplaceAllString(out, "")
	out = newline.ReplaceAllString(out, "")

	return out
}

func (m Maxima) Close() {
	io.WriteString(m.Stdin, "quit();\n")
}

// BuildMatrix builds the string for a matrix based on a row major set of strings for use in Maxima methods
func BuildMatrix(rows, cols int, vals ...string) (out string) {
	if len(vals) < rows*cols {
		log.Fatal("Not enough values to generate matrix!")
	}

	out = "matrix("
	for r := 0; r < rows; r++ {
		out += "["
		for c := 0; c < cols; c++ {
			out += vals[c+r*cols]
			if c < cols-1 {
				out += ","
			}
		}
		out += "]"
		if r < rows-1 {
			out += ","
		}
	}
	out += ")"

	return
}
