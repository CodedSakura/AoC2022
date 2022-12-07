package main

import (
	"AoC2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	A()
	B()
}

type File struct {
	size int
	name string
}

type Directory struct {
	name      string
	files     []File
	dirs      []*Directory
	parent    *Directory
	totalSize int
}

func mkdir(name string, parent *Directory) *Directory {
	return &Directory{
		name:   name,
		files:  []File{},
		dirs:   []*Directory{},
		parent: parent,
	}
}

func du(dir *Directory) int {
	size := 0
	for _, file := range dir.files {
		size += file.size
	}
	for _, subDir := range dir.dirs {
		size += du(subDir)
	}
	dir.totalSize = size
	return size
}

func aRecurse(dir Directory) int {
	sum := 0
	for _, subDir := range dir.dirs {
		if subDir.totalSize <= 100_000 {
			sum += subDir.totalSize
		}
		sum += aRecurse(*subDir)
	}
	return sum
}

func A() {
	var root Directory
	var pwd *Directory
	initialised := false

	for line := range utils.ReadDayByLine(07) {
		if !initialised {
			cmd := strings.Split(strings.Trim(line, "$ "), " ")
			if cmd[0] != "cd" {
				panic(cmd[0])
			}

			root = *mkdir(cmd[1], nil)
			pwd = &root
			initialised = true
			continue
		}

		if line[0:2] != "$ " {
			file := strings.Split(line, " ")
			if file[0] == "dir" {
				pwd.dirs = append(pwd.dirs, mkdir(file[1], pwd))
			} else {
				size, _ := strconv.Atoi(file[0])
				pwd.files = append(pwd.files, File{
					name: file[1],
					size: size,
				})
			}
			continue
		}

		cmd := strings.Split(strings.Trim(line, "$ "), " ")
		if cmd[0] == "cd" {
			if cmd[1] == ".." {
				pwd = pwd.parent
			} else {
				for _, dir := range pwd.dirs {
					if dir.name == cmd[1] {
						pwd = dir
					}
				}
			}
			continue
		} else if cmd[0] == "ls" {
			if len(cmd) > 1 {
				panic(len(cmd))
			}
		}
	}

	du(&root)

	fmt.Println(aRecurse(root))
}

func bRecurse(dir Directory, min int, reqSize int) int {
	for _, subDir := range dir.dirs {
		if subDir.totalSize >= reqSize {
			if subDir.totalSize < min {
				min = subDir.totalSize
			}
		}
		res := bRecurse(*subDir, min, reqSize)
		if res < min {
			min = res
		}
	}
	return min
}

func B() {
	var root Directory
	var pwd *Directory
	initialised := false

	for line := range utils.ReadDayByLine(07) {
		if !initialised {
			cmd := strings.Split(strings.Trim(line, "$ "), " ")
			if cmd[0] != "cd" {
				panic(cmd[0])
			}

			root = *mkdir(cmd[1], nil)
			pwd = &root
			initialised = true
			continue
		}

		if line[0:2] != "$ " {
			file := strings.Split(line, " ")
			if file[0] == "dir" {
				pwd.dirs = append(pwd.dirs, mkdir(file[1], pwd))
			} else {
				size, _ := strconv.Atoi(file[0])
				pwd.files = append(pwd.files, File{
					name: file[1],
					size: size,
				})
			}
			continue
		}

		cmd := strings.Split(strings.Trim(line, "$ "), " ")
		if cmd[0] == "cd" {
			if cmd[1] == ".." {
				pwd = pwd.parent
			} else {
				for _, dir := range pwd.dirs {
					if dir.name == cmd[1] {
						pwd = dir
					}
				}
			}
			continue
		} else if cmd[0] == "ls" {
			if len(cmd) > 1 {
				panic(len(cmd))
			}
		}
	}

	du(&root)

	reqSize := 30_000_000 - (70_000_000 - root.totalSize)
	fmt.Println(bRecurse(root, root.totalSize, reqSize))
}
