package gof

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {
	file1 := &RegularFile{name: "File1"}
	file2 := &RegularFile{name: "File2"}
	file3 := &RegularFile{name: "File3"}

	folder1 := &Directory{
		name: "Directory1",
	}

	folder1.add(file1)

	folder2 := &Directory{
		name: "Directory2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
}

// Abstract component
type Component interface {
	search(string)
}

// Composite
type Directory struct {
	components []Component
	name       string
}

func (d *Directory) search(keyword string) {
	fmt.Printf("Searching recursively for keyword %s in directory %s\n", keyword, d.name)
	for _, composite := range d.components {
		composite.search(keyword)
	}
}

func (d *Directory) add(c Component) {
	d.components = append(d.components, c)
}

// Leaf
type RegularFile struct {
	name string
}

func (f *RegularFile) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *RegularFile) getName() string {
	return f.name
}
