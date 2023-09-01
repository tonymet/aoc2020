package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	TypeDir int = iota
	TypeFile
)

const (
	StateListDir int = iota
	StateDefault
)

type Entry struct {
	size      int
	name      string
	entryType int
	children  []*Entry
	parent    *Entry
}

func (e *Entry) cd(name string) (*Entry, error) {
	if name == ".." {
		if e.parent == nil {
			panic("no parent")
		}
		return e.parent, nil
	}
	for _, child := range e.children {
		if child.name == name {
			return child, nil
		}
	}
	return nil, errors.New("child Not found " + name)
}

func newEntry() (n Entry) {
	n.children = make([]*Entry, 0)
	return
}

var (
	root Entry
)

func init() {
	root = newEntry()
	root.name = "root"
	root.entryType = TypeDir
}

func main() {
	if stdin := os.Getenv("STDIN"); len(stdin) != 0 {
		stdinFile, err := os.Open(stdin)
		if err != nil {
			panic(err)
		}
		os.Stdin = stdinFile

	}

	scanner := bufio.NewScanner(os.Stdin)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var scanState int = StateDefault
	var cwd = &root
	for scanner.Scan() {
		entry := newEntry()
		line := scanner.Text()
		fmt.Println(line)
		switch {
		case strings.HasPrefix(line, "$ cd"):
			// handle change of director
			entry.name = strings.TrimPrefix(line, "$ cd ")
			if entry.name == "" {
				panic("error parsing name")
			}
			if entry.name == "/" {
				cwd.children = append(root.children, &entry)
				cwd = &entry
			} else {
				var err error
				if cwd, err = cwd.cd(entry.name); err != nil {
					panic(err)
				}
			}
			fmt.Printf("cd : %+v\n", line)
		case strings.HasPrefix(line, "$ ls"):
			// ls
			// state = READDIR
			scanState = StateListDir
			fmt.Printf("ls : %+v\n", line)
		case strings.HasPrefix(line, "dir"):
			if scanState != StateListDir {
				panic("parsing error -- scanstate")
			}
			entry.entryType = TypeDir
			entry.parent = cwd
			fmt.Sscanf(line, "dir %s", &entry.name)
			cwd.children = append(cwd.children, &entry)
			fmt.Printf("direntry: %+v\n", entry)
		default:
			if scanState != StateListDir {
				panic("parsing error -- scanstate")
			}
			fmt.Sscanf(line, "%d %s", &entry.size, &entry.name)
			entry.entryType = TypeFile
			entry.parent = cwd
			cwd.children = append(cwd.children, &entry)
			fmt.Printf("entry: %+v\n", entry)
		}
	}
	fmt.Printf("DirTree: %+v\n", root)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
