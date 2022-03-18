// Let's see how protocol buffer works, practicing with a simple example!
// example 1
package main


import "fmt"
import "github.com/golang/protobuf/proto"
import "log"

func main() {
	fmt.Println("Hello World!")

	elliot := &Person{
		Name: "Elliot",
		Age: 24,
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("Marshalling error", err)
	}

	fmt.Println(data)

	newElliot := &Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("unmarshalling error: ", err)
	}
	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetName())
}