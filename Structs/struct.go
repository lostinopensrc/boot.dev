package Structs

type Rect struct{
	Width int 
	Height int
}

func Area(r Rect) int {	
	return r.Width * r.Height
}