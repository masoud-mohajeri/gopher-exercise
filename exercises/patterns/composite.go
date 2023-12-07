package patterns

import "fmt"

type Component interface {
	search(string)
}

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}

type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)

}
func (f *File) getName() string {
	return f.name
}

func CompositePattern() {
	file1 := &File{name: "File One"}
	file2 := &File{name: "File Two"}
	file3 := &File{name: "File Three"}

	folder1 := &Folder{name: "Folder One"}
	folder1.add(file1)
	folder1.add(file2)
	folder1.add(file3)

	folder2 := &Folder{name: "Folder Two"}
	folder2.add(folder1)

	folder2.search("Keyword")

}
