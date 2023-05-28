package main

import (
	"fmt"
	"github.com/lostinopensrc/boot.dev/Intro"
	"github.com/lostinopensrc/boot.dev/Variables"
	"github.com/lostinopensrc/boot.dev/Conditionals"
	"github.com/lostinopensrc/boot.dev/Functions"
)

func main() {
	Intro.Gophers()
	Variables.BasicVars()
	Variables.ShortHand()
	Variables.InferType()
	Conditionals.IsConditionals()
	concatOutput := Functions.DoConcat("My name " , "is lostinopensrc")
	fmt.Println(concatOutput)
	result,_ := Functions.Divide(4,2) // for correct division ignored the err returned from func Divide in functions package
	fmt.Println(result)
	_,err := Functions.Divide(4,0) // We cannot divide a divident by 0 , print the err 
	fmt.Println(err)
}