package file

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Create(filename string) *os.File {
	file, err := os.Create(fmt.Sprintf("%s.txt", filename))
	if err != nil {
		fmt.Printf("Can't create file %s\n", filename)
	}
	return file
}

func Open(filename string) {
	file, err := os.Open(fmt.Sprintf("./data/%s.txt", filename))
	if err != nil {
		log.Fatalf("Can't open file %s\n", filename)
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Can't read file %s\n", filename)
	}

	fmt.Println(string(data))
}

func Delete(filename string) {

}
