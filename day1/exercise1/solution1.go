
package main

import (
	"fmt"
)

type Matrix struct{
	rows,columns int
	twod [2][2]int
}

func (M1 Matrix) num_of_rows() int {
	return M1.rows
}
func (M1 Matrix) num_of_columns() int {
	return M1.columns
}
func (M1 *Matrix) set_element(i int ,j int ,val int) {
	M1.twod[i][j] = val
}

func add_Matrices(M1, M2 Matrix)  Matrix{
	var M_temp = Matrix{}
	for i:=0;i<2;i++ {
		for j:=0;j<2;j++{
			M_temp.twod[i][j] = M1.twod[i][j]+M2.twod[i][j]
		}
	}
	return M_temp
}

func main(){
	var M1,M2,M3 Matrix = Matrix{},Matrix{},Matrix{}
	M1.rows = 2
	M1.columns = 2
	M1.set_element(0,0,1)
	M1.set_element(0,1,2)
	M1.set_element(1,1,3)
	M1.set_element(1,0,4)
	M2.rows = 2
	M2.columns = 2
	M2.set_element(0,0,5)
	M2.set_element(0,1,6)
	M2.set_element(1,1,7)
	M2.set_element(1,0,8)
	M3 = add_Matrices(M1,M2)
	M3.rows = 2
	M3.columns = 2
	fmt.Println(M3.twod)
	fmt.Println(M3)
}