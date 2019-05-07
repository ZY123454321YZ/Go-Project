package main
import (
	"fmt"
	"geometry/rectangle"
)
func main() {
	rectLen, rectWidth := 6, 7
	fmt.Println("Geometrical shape properties")
	/*Area function of rectangle package used*/
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	/*Diagonal function of rectangle package used*/
	fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))
	fmt.Println("Geometrical shape properties")
}