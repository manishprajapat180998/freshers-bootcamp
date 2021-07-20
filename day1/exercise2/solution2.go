package main


type expression struct{
	var x string
	left *expression
	right *expression
}

func check(c string) bool{
	if c == "+" || c == "-" || c == "*" || c == "/" {
		return true
	}
	return false
}


func tree(str string) expression{
	var str2 string


}
func main(){
	var str string = "a+b-c"
	var str1 expression = tree(str)

}