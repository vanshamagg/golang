package streams

import (
	"io"
	"log"
	"os"
	"strings"
)

func FuncCopy() {
	readable := strings.NewReader("This is a really dope string")
	_, err := io.Copy(os.Stdout, readable)
	if err != nil {
		log.Fatal(err)
	}
}

func FuncCopyBuffer() {
	 r1 := strings.NewReader("This is one hell of a string")
	 r2 := strings.NewReader("This is yet another awesome string")
	 buf := make([]byte, 8)

	 if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
			log.Fatal(err)
	 }

	 
	//  Re-using the exiting buffer


	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
 }



}
