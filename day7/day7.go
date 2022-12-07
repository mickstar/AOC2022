package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	Text string
}

type File struct {
	Filename string
	Size     int64
}

type Directory struct {
	Parent           *Directory
	Location         string
	Files            []File
	ChildDirectories []*Directory
	// really just for debug
	ChildDirectoryStrings []string

	DirectorySizeTotal   int64
	DirectorySizeShallow int64
}

func readFile(filename string) ([]Line, error) {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bytes), "\n")
	data := make([]Line, len(lines))

	for i, line := range lines {
		data[i] = Line{
			Text: line,
		}
	}
	return data, nil
}

/*
* every ls, we create a new directory and fill it with the content
* we're going to have to do this recusively
 */
func interpretFileStructure(lines []Line, currentIndex int, currentDirectory *Directory) (*Directory, int, error) {
	rootLine := lines[currentIndex]

	if !strings.HasPrefix(rootLine.Text, "$ cd ") {
		return nil, -1, fmt.Errorf("expected line to start with $ cd, got %s", rootLine.Text)
	}

	directoryName := strings.TrimPrefix(rootLine.Text, "$ cd ")

	if directoryName == ".." {
		// go up a directory
		return currentDirectory.Parent, currentIndex + 1, nil
	}

	files := make([]File, 0)
	directoryPaths := make([]string, 0)

	newIndex := len(lines)

	for i := currentIndex + 1; i < len(lines); i++ {
		line := lines[i]
		if strings.HasPrefix(line.Text, "$ cd ") {
			// we need to return at this point.
			newIndex = i
			break
		} else if strings.HasPrefix(line.Text, "$ ls") {
			// dont care
			continue
		} else if strings.HasPrefix(line.Text, "dir") {
			// handle directory in form
			// dir vtb
			directoryPaths = append(directoryPaths, strings.Split(line.Text, " ")[1])
			continue
		} else {
			// must be a file.
			// 299692 rbssdzm.ccn
			splitString := strings.Split(line.Text, " ")
			files = append(files, File{
				Filename: splitString[1],
				Size:     forceReadInt(splitString[0]),
			})
		}
	}

	var directory = Directory{
		Parent:                currentDirectory,
		Location:              directoryName,
		Files:                 files,
		ChildDirectories:      make([]*Directory, 0),
		ChildDirectoryStrings: directoryPaths,
	}
	if currentDirectory != nil {
		currentDirectory.ChildDirectories = append(currentDirectory.ChildDirectories, &directory)
	}
	return &directory, newIndex, nil
}

func calculateDirectorySizeForTree(directory *Directory) {
	if directory == nil {
		return
	}

	for _, file := range directory.Files {
		directory.DirectorySizeShallow += file.Size
		directory.DirectorySizeTotal += file.Size
	}

	for _, childDirectory := range directory.ChildDirectories {
		calculateDirectorySizeForTree(childDirectory)
		directory.DirectorySizeTotal += childDirectory.DirectorySizeTotal
	}
}

func forceReadInt(s string) int64 {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return int64(n)
}

func main() {
	data, err := readFile("day7/day7_input.txt")
	//data, err := readFile("bigboys/bigboy_day7.txt")
	if err != nil {
		panic(err)
	}

	i := 0
	var rootDirectory *Directory
	var currentDirectory *Directory
	for true {
		if i >= len(data) {
			fmt.Println("Finished reading")
			break
		}
		var directory *Directory
		directory, i, err = interpretFileStructure(data, i, currentDirectory)
		if err != nil {
			panic(err)
		}
		if rootDirectory == nil || directory.Location == "/" {
			rootDirectory = directory
		}
		currentDirectory = directory
	}

	calculateDirectorySizeForTree(rootDirectory)
	fmt.Println(rootDirectory)
	//printDirectoryTree(rootDirectory, 0)

	fmt.Println("Part A ", partA(rootDirectory, 0))

	partB(rootDirectory)

}

func printDirectoryTree(directory *Directory, indent int) {
	if directory == nil {
		return
	}

	if directory.Location == "a" {
		fmt.Println(directory)
	}

	for i := 0; i < indent; i++ {
		fmt.Print(" ")
	}
	fmt.Println(directory.Location, "total ", directory.DirectorySizeTotal, "shallow", directory.DirectorySizeShallow)
	for _, file := range directory.Files {
		for i := 0; i < indent; i++ {
			fmt.Print(" ")
		}
		fmt.Print("- ")
		fmt.Println(file.Filename, file.Size)
	}
	for _, childDirectory := range directory.ChildDirectories {
		printDirectoryTree(childDirectory, indent+1)
	}
}

func partA(root *Directory, currentTotal int64) int64 {
	var total int64 = 0
	if root == nil {
		return 0
	}
	if root.DirectorySizeTotal < 100000 {
		total += root.DirectorySizeTotal
	}
	for _, child := range root.ChildDirectories {
		total += partA(child, currentTotal)
	}

	return total
}

const TOTAL_DISK_SPACE = 70000000
const NEED_TO_FREE = 30000000

func partB(root *Directory) {
	var curFree int64 = TOTAL_DISK_SPACE - 49192532
	var toFree int64 = NEED_TO_FREE - curFree
	if root == nil {
		return
	}

	if root.DirectorySizeTotal >= toFree {
		fmt.Println("Can Delete ", root.Location, root)
	}
	for _, child := range root.ChildDirectories {
		partB(child)
	}
}
