// Let's see how protocol buffer works, practicing with a simple example!
// example 2 : store personal data of name, age, social followers

package main

import "fmt"
import "log"
import "github.com/golang/protobuf/proto"

func main() {
	fmt.Println("Helloworld")

	// Store struct data
	elliot := &Person{
		Name: "Elliot",
		Age: 24,
		SocialFollowers: &SocialFollowers{
			Twitter: 1400,
			Youtube: 2500,
		},
	}

	// marshall
	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("Marshalling error:", err)
	}
	fmt.Println(data)

	// unmarshall
	newElliot := &Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("Unmarshalling error:", err)
	}
	fmt.Println(newElliot.GetName())
	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.SocialFollowers.GetTwitter())
	fmt.Println(newElliot.SocialFollowers.GetYoutube())
}