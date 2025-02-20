package main

import (
	"fmt"
	"net/http"
)

/*
	Go was built from the ground up
	to be aware of the web and to be able to
	handle with it in a meaningful fashion.
*/

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// From docs
		// Fprintf formats according to a format specifier and writes to 'w'
		// It returns the number of bytes written and any write error encountered
		n, err := fmt.Fprintf(w, "Hello World")
		if err != nil {
			fmt.Println(err)
		}

		// Sprintf takes multiple data types and returns them as a formatted string
		// There are many options for formatting integers. Use %d for standard, base-10 formatting.
		fmt.Println(fmt.Sprintf("bytes written: %d", n))

		/*
			'_' is a blank identifier
			Identifiers are the user-defined name of the program components used for the identification purpose.
			Golang has a special feature to define and use the unused variable using Blank Identifier.
			Unused variables are those variables that are defined by the user throughout the program but he/she never makes use of these variables.
			These variables make the program almost unreadable.

			As you know, Golang is a more concise and readable programming language so it doesn’t allow the programmer to define an unused variable if you do such,
			then the compiler will throw an error.
			The real use of Blank Identifier comes when a function returns multiple values,
			but we need only a few values and want to discard some values. Basically,
			it tells the compiler that this variable is not needed and ignored it without any error.
			It hides the variable’s values and makes the program readable. So whenever you will assign a value to Blank Identifier it becomes unusable.
		*/
		_ = http.ListenAndServe(":8080", nil)

	})
}
