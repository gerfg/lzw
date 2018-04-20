package main

func main() {
	fileName := "teste"
	encode("instances/" + fileName)
	decodeFile("decoded/" + fileName + ".cpr")
}
