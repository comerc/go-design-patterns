package main

import "fmt"

// Позволяет копировать объекты, не вдаваясь в подробности их реализации.

type Node interface {
	print(string)
	clone() Node
}

type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) clone() Node {
	return &File{name: f.name + "_clone"}
}

type Folder struct {
	children []Node
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *Folder) clone() Node {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []Node
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Node{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []Node{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone().(*Folder)
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}
