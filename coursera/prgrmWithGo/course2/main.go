package main

func main() {

	x := 2
	ex(&x)
	print(x)
}

func ex(y *int) {
	*y = *y + 1
}
