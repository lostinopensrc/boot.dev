package Variables

import "fmt"

func InferType(){
	pennypersms := 2
	fmt.Printf("The type of pennypersms is %T\n",pennypersms)

	pennypersms1 := 2.0
	fmt.Printf("The type of pennypersms1 is %T\n",pennypersms1)
}