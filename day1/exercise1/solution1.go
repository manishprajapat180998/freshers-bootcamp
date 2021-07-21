
package main

import (
	"fmt"
)

type Matrix struct{
	rows,columns int
	twod [2][2]int
}

func (M1 Matrix) num_of_rows() int {
	return len(M1.twod)
}
func (M1 Matrix) num_of_columns() int {
	return len(M1.twod[0])
}
func (M1 Matrix) set_element(i int ,j int ,val int) {
	M1.twod[i][j] = val
}

func (M1 Matrix)add_Matrices(M2 Matrix)  Matrix{
	for i:=0;i<2;i++ {
		for j:=0;j<2;j++{
			M1.twod[i][j] = M1.twod[i][j]+M2.twod[i][j]
		}
	}
	return M1
}

func main() {
	var M2 = Matrix{2, 2, [2][2]int{{5, 6}, {7, 8}}}
	var M1 = Matrix{2, 2, [2][2]int{{1, 2}, {3, 4}}}
	M1 = M1.add_Matrices(M2)
	fmt.Println(M1.twod)
	fmt.Println(M1)
}