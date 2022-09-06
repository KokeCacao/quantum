package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type quantum struct {
	variables map[string]bool
}

func main() {
	quiet := false
	if len(os.Args) > 0 {
		switch os.Args[1] {
		case "--quiet":
			quiet = true
		case "-q":
			quiet = true
		}
	}
	reader := bufio.NewReader(os.Stdin)
	q := quantum{
		make(map[string]bool),
	}
	rand.Seed(time.Now().UnixNano())

	for {
		if !quiet {
			fmt.Print("$ ")
		}
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println()
			return
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		str = strings.TrimSuffix(str, "\n")
		if len(str) != 0 {
			ret := q.exec(str)
			if !quiet {
				fmt.Println(ret)
			} else {
				fmt.Print(ret)
			}
		}
	}
}

func (q *quantum) createVariables(vars []string) {
	for _, s := range vars {
		if _, ok := q.variables[s]; !ok {
			q.variables[s] = false
		}
	}
}

func (q *quantum) exec(cmd string) string {
	slice := strings.Fields(cmd)
	if len(slice) == 0 {
		return ""
	}

	switch slice[0] {
	case "Flip":
		if s := checkArgs(slice, 2); s != "" {
			return s
		}
		q.createVariables(slice[1:])
		q.variables[slice[1]] = rand.Intn(2) == 0
	case "Noise":
		if s := checkArgs(slice, 2); s != "" {
			return s
		}
		q.createVariables(slice[1:])
		if rand.Intn(3) == 0 {
			q.variables[slice[1]] = !q.variables[slice[1]]
		}
	case "Add_1_to_Coin":
		if s := checkArgs(slice, 2); s != "" {
			return s
		}
		q.createVariables(slice[1:])
		q.variables[slice[1]] = !q.variables[slice[1]]
	case "Add_Coin":
		if s := checkArgs(slice, 3); s != "" {
			return s
		}
		q.createVariables(slice[1:])
		q.variables[slice[2]] = q.variables[slice[1]] != q.variables[slice[2]]
	case "Add_NOT_Coin":
		if s := checkArgs(slice, 3); s != "" {
			return s
		}
		q.createVariables(slice[1:])
		q.variables[slice[2]] = q.variables[slice[1]] == q.variables[slice[2]]
	case "Add_AND_Coin":
		if s := checkArgs(slice, 4); s != "" {
			return s
		}
		q.createVariables(slice[1:])
		q.variables[slice[3]] = (q.variables[slice[1]] && q.variables[slice[2]]) == q.variables[slice[3]]
	case "Print":
		if s := checkArgs(slice, 2); s != "" {
			return s
		}
		q.createVariables(slice[1:])
		return strconv.FormatBool(q.variables[slice[1]])
	case "Print_All":
		if s := checkArgs(slice, 1); s != "" {
			return s
		}
		q.createVariables(slice[1:])
		return fmt.Sprintf("%v", q.variables)
	default:
		return "invalid command"
	}
	return ""
}
