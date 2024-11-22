package file

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Create(name string) *os.File {
	file, err := os.Create(fmt.Sprintf("%s.txt", name))
	if err != nil {
		fmt.Printf("Can't create file %s\n", name)
	}
	return file
}

func Open(name string) *[]byte {
	file, err := os.Open(fmt.Sprintf("./data/%s.json", name))
	if err != nil {
		log.Fatalf("Can't open file %s\n", name)
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Can't read file %s\n", name)
	}

	return &data
}

func Write(name string, data *[]byte) error {
	if err := os.WriteFile(fmt.Sprintf("./data/%s.json", name), *data, 0644); err != nil {
		return err
	}
	return nil
}

func Delete(name string) {

}
