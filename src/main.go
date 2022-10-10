package src

import (
	"fmt"

	"github.com/Jeremy643/playground/src/utils"
)

func main() {
	fmt.Println("Hello World!")

	i := 1
	fmt.Println(utils.FromIntPtr(&i))

	human := utils.CreateHuman("John", 55)
	fmt.Printf("%s is %d!\n", human.GetName(), human.GetAge())

	messagePrinter("Hello World.")
}

func shouldStay(stay *bool) string {
	if stay == nil {
		return "You should not stay."
	}
	if *stay {
		return "You should stay."
	}
	return "You should not stay."
}

func messagePrinter(s string) {
	fmt.Println(s)
}
