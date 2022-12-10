package utilities

import (
	"bufio"
	"os"
	"path"
	"strconv"
	"strings"
)

type Directory struct {
	Name                   string
	AbsolutePath           string
	DirectChildrenFileSize uint64
	// VV: This is actually a set :)
	Directories map[string]int8

	calculatedSize *uint64
}

type Filesystem map[string]*Directory

func (d *Directory) CountSize(root *Filesystem) uint64 {
	if d.calculatedSize != nil {
		return *d.calculatedSize
	}
	sum := d.DirectChildrenFileSize

	for child, _ := range d.Directories {
		path := path.Join(d.AbsolutePath, child)
		cd := (*root)[path]
		sum += cd.CountSize(root)
	}

	d.calculatedSize = new(uint64)
	*d.calculatedSize = sum
	return *d.calculatedSize
}

func NewDirectory(name, absolutePath string, DirectChildrenFileSize uint64, directories map[string]int8) *Directory {
	new_dir := Directory{
		Name:                   name,
		AbsolutePath:           absolutePath,
		DirectChildrenFileSize: DirectChildrenFileSize,
		Directories:            directories,
	}

	return &new_dir
}

func ReadScanner(scanner *bufio.Scanner) (Filesystem, error) {
	root := make(map[string]*Directory)

	curr_path := "/"
	root[curr_path] = NewDirectory("", curr_path, 0, make(map[string]int8))

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		tokens := strings.Split(line, " ")

		if tokens[0] == "$" {
			cmd := tokens[1]

			if cmd == "cd" {
				dir := tokens[2]

				if dir == "/" {
					curr_path = "/"
				} else if dir == ".." {
					parent, _ := path.Split(curr_path)

					if len(parent) > 1 {
						curr_path = parent[:len(parent)-1]
					} else {
						curr_path = "/"
					}
				} else {
					// VV: this is a set
					root[curr_path].Directories[dir] = 0

					curr_path = path.Join(curr_path, dir)
					root[curr_path] = NewDirectory(dir, curr_path, 0, make(map[string]int8))
				}
			}
		} else if tokens[0] == "dir" {
			root[curr_path].Directories[tokens[1]] = 0
		} else {
			filesize, err := strconv.ParseUint(tokens[0], 10, 64)
			if err != nil {
				return root, err
			}

			root[curr_path].DirectChildrenFileSize += filesize
		}
	}

	return root, scanner.Err()
}

func ReadString(text string) (Filesystem, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) (Filesystem, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
