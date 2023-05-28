package Variables

import "fmt"

// short hand var declaration only works inside a function (including main function) not outside/globally
func ShortHand(){

	congrats := "happy Birthday Gophers"
	fmt.Println(congrats)

}