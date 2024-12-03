package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)

type templateData struct {
	Part          string
	PartUpperCase string
	PackageName   string
	Year          int
	Day           int
}

func main() {

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	//var resp map[string]interface{}
	day, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("Generating", year, day, time.Now().Day(), "Args", os.Args[1], os.Args[2])

	err = os.MkdirAll(fmt.Sprintf("cmd/year%vday%v", year, day), 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	//TODO: Validation if file does not exists
	generateSolutionFile("a", "A", year, day)
	generateSolutionFile("b", "B", year, day)
	generateImportFile(year, day)
	_, err = os.Create(fmt.Sprintf("cmd/year%vday%v/input.txt", year, day))
	if err != nil {
		return
	}

	dir, err := os.ReadDir("cmd")
	if err != nil {
		fmt.Println(err)
		return
	}

	var dirsTemplate []string

	for _, directory := range dir {
		if directory.Name() != "root.go" {
			dirsTemplate = append(dirsTemplate, directory.Name())
		}
	}

	generateRootFile(dirsTemplate)

}

func generateRootFile(dirs []string) {
	data := struct {
		Dirs []string
	}{
		Dirs: dirs,
	}

	tmpl, err := template.New("rootTemplate.tmpl").Funcs(template.FuncMap{}).ParseFiles("gen/rootTemplate.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := os.Create("cmd/root.go")
	if err != nil {
		fmt.Println(err)
		return
	}

	tmpl.Execute(out, data)
}

func generateSolutionFile(part, upperCasePart string, year int, day int) {
	data := templateData{
		Part:          part,
		PartUpperCase: upperCasePart,
		Year:          year,
		Day:           day,
	}

	tmpl, err := template.New("dayProblemTemplate.tmpl").Funcs(template.FuncMap{}).ParseFiles("gen/dayProblemTemplate.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := os.Create(fmt.Sprintf("cmd/year%vday%v/%v.go", year, day, part))
	if err != nil {
		fmt.Println(err)
		return
	}

	tmpl.Execute(out, data)
}

func generateImportFile(year int, day int) {
	data := templateData{
		Year: year,
		Day:  day,
	}

	tmpl, err := template.New("importTemplate.tmpl").Funcs(template.FuncMap{}).ParseFiles("gen/importTemplate.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := os.Create(fmt.Sprintf("cmd/year%vday%v/import.go", year, day))
	if err != nil {
		fmt.Println(err)
		return
	}

	tmpl.Execute(out, data)
}
