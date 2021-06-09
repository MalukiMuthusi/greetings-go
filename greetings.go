package greetings

import (
	"fmt"

	"github.com/malukimuthusi/hellogo"
)

func main() {
	message := hellogo.Hello("Maluki")
	fmt.Println(message)
}
