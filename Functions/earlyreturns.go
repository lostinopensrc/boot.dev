package Functions

import "errors"

func Divide(divident , divisor int) (int , error){
	if divisor == 0 {
		return 0, errors.New("Cant divide by 0")
	}
	return divident/divisor,nil
}