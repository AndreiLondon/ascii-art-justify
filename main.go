package main

import "fmt"

func main() {
	fileName, text, banner := getArguments()
	// If it's an empty string it will exit the program

	if text == "" {
		return
	}
	if text == "\\n" {
		fmt.Println()
		return
	}
	convertAndWrite(fileName, text, banner)

	// const s = "Hello world! 👋"

	// // print to the terminal
	// buf := alignCenter(s)
	// fmt.Println(buf)
	// fmt.Fprintln(os.Stdout, buf)
	// io.Copy(os.Stdout, buf)

}

func convertAndWrite(fileName string, text string, banner string) {
	path := banner + ".txt"
	output := readFile(path)
	myMap := parseTemplate(output)
	result := printMessage(text, myMap)
	writeToFile(fileName, []byte(result))
}
