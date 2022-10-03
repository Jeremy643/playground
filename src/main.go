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

	fmt.Println("1 commit")
	fmt.Println("2 commit")
	fmt.Println("3 commit")

	println("hllo")
	println("wrld")
}
