package arrays

import "fmt"

//Rows represent a row in 2d-nd array
type Rows []int
type TwoDArray []Rows

// type Shape struct {
// 	rows    int
// 	columns int
// }
type Arrays struct {
	arr  TwoDArray
	rows int
	cols int
}

func NewTwoDArray(arr TwoDArray) Arrays {
	columns := len(arr[0])
	rows := len(arr)
	return Arrays{arr: arr, rows: rows, cols: columns}
}

//PrintArr prints the array
func (ptr Arrays) PrintArr() {
	for _, row := range ptr.arr {
		fmt.Println("Row", row)
	}
}

//Only applicable to Square matrices
//RotateByNinety rotates a 2D array by 90 degress clockwise
func (ptr Arrays) RotateByNinety() {
	for x := 0; x < ptr.cols/2; x++ {
		for y := x; y < ptr.cols-x-1; y++ {
			temp := ptr.arr[x][y]
			//First row elements
			ptr.arr[x][y] = ptr.arr[ptr.cols-y-1][x]
			//First column elements
			ptr.arr[ptr.cols-y-1][x] = ptr.arr[ptr.cols-x-1][ptr.cols-y-1]
			//Last row elements
			ptr.arr[ptr.cols-x-1][ptr.cols-y-1] = ptr.arr[y][ptr.cols-x-1]
			//Last column elements
			ptr.arr[y][ptr.cols-x-1] = temp
		}
	}
}
