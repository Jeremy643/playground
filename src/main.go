package src

import (
	"fmt"

	"github.com/Jeremy643/playground/src/utils"
)

func main() {
	fmt.Println("Hello World!")

	i := 1
	fmt.Println(utils.FromIntPtr(&i))
}
