package main

func main {
	adder(add)
	adder1(add)

}

func adder(f func(int,int)int) {
	val := add(3,5)
	fmt.Println(val)
}

type myFunctionDataType func(int, int) int

func add( i int, j int) int {

	return i+j

}

func adder1(add myFunctionDataType) {
	val := add(3,5)
	fmt.Println(val)
}