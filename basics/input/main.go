package main

import (
	"fmt"
	"log"
	"strings"
)

/*
The fmt package provides functions for scanning strings, which is the process of parsing strings that contain values
separated by spaces.

Scan(...vals)
This function reads text from the standard in and stores the space-separated values into specified arguments.
Newlines are treated as spaces, and the function reads until it has received values for all of its arguments.
The result is the number of values that have been read and an error that describes any problems.

Scanln(...vals)
This function works in the same way as Scan but stops reading when it encounters a newline character.

Scanf(template, ...vals)
This function works in the same way as Scan but uses a template string to select the values from the input it receives.

Fscan(reader, ...vals)
This function reads space-separated values from the specified reader,newlines are treated as spaces, and
the function returns the number of values that have been read and an error that describes any problems.

Fscanln(reader, ...vals)
This function works in the same way as Fscan but stops reading when it encounters a newline character.

Fscanf(reader, template,...vals)
This function works in the same way as Fscan but uses a template to select the values from the input it receives.

Sscan(str, ...vals)
This function scans the specified string for space-separated values, which are assigned to the remaining arguments.
The result is the number of values scanned and an error that describes any problems.

Sscanln(str, template,...vals)
This function works in the same way as Sscanln but stops scanning the string as soon as a newline character is encountered.

Sscanf(str, template,...vals)
This function works in the same way as Sscan but uses a template to select values from the string.
*/
func main() {
	var input1, input2 string
	number, _ := fmt.Scan(&input1, &input2)
	fmt.Printf("指针接收了%d个键盘输入,分别是%v和%v\n", number, input1, input2)

	var (
		num int
		b   bool
		str string
	)
	r := strings.NewReader("10 false GFG")
	number2, _ := fmt.Fscan(r, &num, &b, &str)
	//n, err:= fmt.Fscanf(r, "%d %t %s", &num, &b, &str)
	fmt.Printf("指针从reader中接收了%d个键盘输入,分别是%v %v %v\n", number2, num, b, str)

	n, err := fmt.Sscan("GFG 3", &str, &num)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("指针从字符串中接收了%d个键盘输入,分别是%v %v \n", n, str, num)
}
