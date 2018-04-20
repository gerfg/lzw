package main

func main() {
	fileName := "generated.fib25"
	encode("instances/" + fileName)
	decodeFile("decoded/" + fileName + ".cpr")
}
