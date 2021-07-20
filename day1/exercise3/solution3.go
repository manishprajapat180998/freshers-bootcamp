package main

import "fmt"

type Salary interface{
	fulltime() int
	contractor() int
	freelancer() int
}

type employees struct{
	full int
	cont int
	free int
}
func(x employees) fulltime() int{
	return x.full*500
}
func(y employees) contractor() int{
	return y.cont*100
}
func(z employees) freelancer() int{
	return z.free*10
}
func main(){
	var employee Salary
	employee = employees{full: 28,cont: 28,free: 200}


	fmt.Println("fulltime employees total salary :",employee.fulltime())
	fmt.Println("contractors total salary: ",employee.contractor())
	fmt.Println("freelancers total salary: ",employee.freelancer())
}
