// Let's see how protocol buffer works, practicing with a simple example!
// example 1 : store the personal data of name and age
package main

import "fmt"
import "github.com/golang/protobuf/proto"
import "log"

func main() {
	fmt.Println("Hello World!")

	// store structed data
	elliot := &Person{
		Name: "Elliot",
		Age: 24,
	}

	// marshall the data into proto data format
	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("Marshalling error:", err)
	}
	fmt.Println(data)

	// unmarshall the data into structured data format
	newElliot := &Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal(" Unmarshalling error: ", err)
	}
	fmt.Println(newElliot.GetAge()) 
	fmt.Println(newElliot.GetName())
		// GetAge(), GetName() is the function in person.pb.go
}