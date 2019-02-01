package main

import "reflect"

func main() {
	
}

func walk (x interface{}, fn func(input string)){
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}