package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

func main() {
	reader := strings.NewReader("Heyy what it is?")
	fmt.Println(reflect.TypeOf(reader))
	p := make([]byte, 6)

	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println(string(p[:n]))
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(p[:n]))
		fmt.Println(reflect.TypeOf(reader))
		fmt.Println("Hey there")
	}
}
