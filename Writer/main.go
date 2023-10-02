package main

import (
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) error {
	_, e := w.Write([]byte(p.first))
	return e
}
func main() {
	f, err :=os.Create("output.txt")
	if err != nil{
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	s:=[]byte("Hello ")

	_, err = f.Write(s)
	if err != nil{
		log.Fatalf("error %s", err)
	}
}
