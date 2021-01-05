package main

import (
	"github.com/rapidloop/skv"
	"log"
	"fmt"
)
type session struct { 
	info string
	result string
}
func main() {
	store,err := skv.Open("./store.db")
	if err != nil {
		log.Fatal(err)
	}

// store a complex object without making a fuss
var info session
store.Put("safi", "This is Safi")

// get it back later, identifying the object with a string key
store.Get("safi", &info.result)
fmt.Println(info.result)
// delete it when we no longer need it
store.Delete("sess-341356")

// bye
store.Close()

	


}
