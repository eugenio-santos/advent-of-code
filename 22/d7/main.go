package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type node struct {
	name     string
	size     int
	isDir    bool
	parent   *node
	children map[string]*node
}

var min int = math.MaxInt64
var maxStrg int = 70000000
var reqStrg int = 30000000
var target int = 1

func main() {
	dat, _ := os.ReadFile("in")

	datA := strings.Split(string(dat), "\n")

	root := &node{name: "/", size: 0, isDir: true, parent: nil, children: make(map[string]*node)}
	cursor := root

	for _, line := range datA[1:] {
		if strings.HasPrefix(line, "$") {
			if strings.HasPrefix(line, "$ cd") { // update cursor
				folder := strings.Split(line, " ")[2]
				if folder == ".." {
					cursor = cursor.parent
				} else {
					cursor = cursor.children[strings.Split(line, " ")[2]]
				}
			}
		} else { // reading contents of ls
			if strings.HasPrefix(line, "dir") { // update tree with new folder
				name := strings.Split(line, " ")[1]
				if _, ok := cursor.children[name]; !ok { // create new folder
					cursor.children[name] = &node{name: name, size: 0, isDir: true, parent: cursor, children: make(map[string]*node)}
				}
			} else { // update tree with file and buble up size
				size, _ := strconv.Atoi(strings.Split(line, " ")[0])
				file := &node{name: strings.Split(line, " ")[1], size: size, isDir: false, parent: cursor, children: nil}
				cursor.children[file.name] = file
				updateTreeSize(cursor, file.size)
			}
		}
	}

	target = reqStrg - (maxStrg - root.size)
	itAllChildren(root)
	fmt.Println("answ:", min)
}

func updateTreeSize(node *node, size int) {
	node.size += size
	if node.parent != nil {
		updateTreeSize(node.parent, size)
	}
}

func itAllChildren(node *node) {
	for _, child := range node.children {
		if child.isDir {
			itAllChildren(child)
			if child.size > 8748071 && child.size < min {
				min = child.size
			}
		}
	}
}
