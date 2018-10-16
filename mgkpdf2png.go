package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Songmu/prompter"
)

func main() {
	input := (&prompter.Prompter{
		Message: "Enter path to PDF",
	}).Prompt()
	fmt.Printf("Input PDF: %v\n", input)
	pdf, _ := filepath.Abs(input)
	if _, err := os.Stat(pdf); os.IsNotExist(err) {
		log.Fatal(err)
	}
	dir := filepath.Dir(pdf)
	filename := filepath.Base(pdf)
	ext := filepath.Ext(filename)
	if ext != ".pdf" {
		log.Fatalf("File is not PDF but %v", ext)
	}
	base := strings.TrimSuffix(filename, ext)
	png := filepath.Join(dir, strings.Join([]string{base, ".png"}, ""))
	fmt.Printf("Output PNG: %v\n", png)
	if prompter.YN("Start convert ?", true) {
		args := []string{"-density", "300", pdf, png}
		cmd := exec.Command("magick", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("Exit")
	}
}
