package lessons

import (
	"fmt"
	"log"
)

type Dog struct {
	Name  string
	Color string
}

func (d Dog) Run() {
	fmt.Println("dog is running")
}
func (d Dog) SayName() {
	log.Printf("My name is %v", d.Name)
}
