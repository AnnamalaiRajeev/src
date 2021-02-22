package main

import (
	"fmt"
	"net/http"
)

// Dog is dog
type Dog struct {
	Name  string
	Legs  int64
	Voice string
}

func (d *Dog) name() (l string) {
	d.Name = "jimmy"
	l = d.Name
	fmt.Printf("%T\n", d.Name)
	return
}

func (d *Dog) speak() (voice string) {
	d.Voice = "bark"
	voice = d.Voice
	return
}

func speak(d *Dog) (voice string) {
	d.Voice = "Bark"
	voice = d.Voice
	return
}

type server struct {
}

//say hello
func Handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world go to /get/...to get dog name")
}

//Return dog name
func Handler2(w http.ResponseWriter, r *http.Request) {
	d1 := Dog{}
	x := d1.name()
	fmt.Println(x)
	fmt.Fprint(w, x)

}

func main() {
	var x string
	fmt.Println("test")

	server := func() {
		http.HandleFunc("/", Handler1)
		http.HandleFunc("/get/", Handler2)
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(err)
		}
	}

	go server()

	n, err := fmt.Scanf("%s", &x)
	if err != nil {
		panic("error enter proper text")
	} else {
		fmt.Printf("%v\n", x)
		fmt.Printf("%d\n", n)
	}

}
