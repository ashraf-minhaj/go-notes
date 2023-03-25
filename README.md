# go-notes

<!-- <div align="center"> -->
![img](https://www.freecodecamp.org/news/content/images/2022/03/dancing-gopher.gif)

I'm learning golang from the book Head First Go. Let's see in how many days I get to finish this off.
 <!-- </div> -->

## Chapters [4/16]
 * Chapter 1. let’s get going: Syntax Basics [10001-10011]
 * Chapter 2. which code runs next?: Conditionals and Loops [20000-20008]
 * Chapter 3. call me: Functions, Pointers [30001-30005]
 * Chapter 4. bundles of code: Packages []
 * Chapter 5. on the list: Arrays [50001-50006]
 * Chapter 6. appending issue: Slices [60001-]


## To run (basic)
 * Init - `go mod init program_name`
 * Add module requirements and sums - `go mod tidy`
 * Format it - `go fmt`. Why? Because other devs want you to deliver standard code.
 * Run  - `go run .`

## Data types:
 * STRING - "yo bro"
 * RUNES - 'x' (int8, int32)
 * BOOL - true/false
 * NUMBER - integer (int), float (float32, float64)

## Data Structures
 * Array `var numbers [10]int`, `numbers := [3]float64{}`, `numbers := [3]float64{1, 2, 3}`
 * Slice `var mySlice []string`, `myslice := []string`
 * Map
 * Struct

## zero values
 declare vars without assigning anything in it.
 * int - 0
 * float - 0.0
 * bool - false
 * string - '' (empty string)

## Naming rules
 * a name must start with a letter
 * if it begins with a lower case letter, it's unexported or private element.
 * if it begins with a Capital letter, it's considered as exported and can be accessed from outside or public element.
   this is why `p` in `fmt.Println()` is capitalized.

conventions - 
 * names should be camel case i.e. `minhajIsGreat`
 * if it's obvious, short is preferred i.e. `max` instead of `maximum`.

## Conversions
 * math and comparison operations require same values to operate or compare with.

## Function vs Method
 * Function — a set of instructions that perform a task.
 * Method — a set of instructions that are associated with an object.
 * The dot indicates that the thing on its right belongs to the thing on its left.
  ```
  now := time.Now() // time package, Now() is function
  now.Year()        // now value/object, Year() is method
  ```

## Block vs Variable scope
 * blocks - segments of codes
 * scope - A variable’s scope consists of the block it’s declared in and any blocks nested within that block.
 * Just like with conditionals, the scope of any variables declared within a loop’s block is limited to that block (the init statement, condition expression, and post statement can be considered part of that scope too).

## Pointers
 * Values that represent the address of a variable are known as pointers, because **they point to the location where the variable can be found**.
 * type of a pointer is written with a * symbol, followed by the type of the variable the pointer points to.
 * The type of a pointer to an int variable.
 * *int (read as - pointer to int)
 * The * operator can also be used to update the value at a pointer.
 * In Go, it’s okay to return a pointer to a variable that’s local to a function. Even if that variable is no longer in scope, as long as you still have the pointer, Go will ensure you can still access the value.

## Arrays
 * Arrays hold collection of values that all share the same type.
 * Example - pill box.
 * **An array holds a specific number of elements, and it cannot grow or shrink.**
 * `var myArray [n]string`
 * If you know in advance what values an array should hold, you can initialize the array with those values using an **array literal** `[3]int{4, 9, 6}`.

## Slices
 * Slices are a collection type that can grow to hold additional items

## OS
 * os.Open function opens a file
 * os.File value to bufio.NewScanner returns a bufio.Scanner value whose Scan and Text methods can be used to read a line at a time from the file as strings.

## Build Go apps
 * compile the source code intor binary file - `go build file_name`
 * run the exe on windows, or just file name on mac/linux. `./file_name`
 * traverse -  for...range loop syntax.

## Notes:
 * go is a statically typed language. If you type wrong type of value in the wrong place, go will let you know.
 * In Go, := is for declaration + assignment.
 * blank identifier - single underscore if you are not using a variable `_`
 * declare two variables -> latest one will be assignment. 
 ```
 out1, err := something.SometingElse() // both vars declared
 out2, err := something.MoreSomething() // out2 declared, err assigned
 ```
 * Functions, conditionals, and loops all have blocks of code that appear within {} braces.
 * In addition to a name, a package may have an import path that is required when it is imported. i.e. `math/rand`
 * Man user input taking sucks!
 * Go is a “pass-by-value” language; function parameters receive a copy of the arguments from the function call.
 * error strings should not end with punctuation or newlines, or should not be capitalized.
 * & - ampersand
 * You can get the address of a variable using & (an ampersand), which is Go’s “address of” operator
 * **Go is a pass-by-value language**, meaning that function parameters receive a copy of any arguments from the caller
 * **PANIC** - an error that occurs while your program is running (not while compiling)



Feel free to learn and share.
Refeerence: Head First Go.
(c) Ashraf Minhaj