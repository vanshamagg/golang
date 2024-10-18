package ioop

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func Ioop() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize",
		"Cgo is not Go",
		"Errors are values",
		"Don't panic",
	}

	var writer bytes.Buffer
	for _, p := range proverbs {

		n, err := writer.Write([]byte(p))

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		if n != len(p) {
			log.Fatal("Failed to write data")
			os.Exit(1)
		}
	}

	fmt.Println(writer.String())

}
