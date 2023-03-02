# go-notes

<!-- <div align="center"> -->
![img](https://www.freecodecamp.org/news/content/images/2022/03/dancing-gopher.gif)

I'm learning golang from the book Head First Go. Let's see with how many days I get to finish this off.
 <!-- </div> -->

## Chapters [2/16]
 * Chapter 1. let’s get going: Syntax Basics [10001-10011]
 * Chapter 2. which code runs next?: Conditionals and Loops [20000-20008]
 * Chapter 3. call me: Functions [30001-cont.]


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

# Block vs Variable scope
 * blocks - segments of codes
 * scope - A variable’s scope consists of the block it’s declared in and any blocks nested within that block.
 * Just like with conditionals, the scope of any variables declared within a loop’s block is limited to that block (the init statement, condition expression, and post statement can be considered part of that scope too).

## Build Go apps
 * compile the source code intor binary file - `go build file_name`
 * run the exe on windows, or just file name on mac/linux. `./file_name`

## Notes:
 * go is a statically typed language. If you type wrong type of value in the wrong place, go will let you know.
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



Feel free to learn and share.
Refeerence: Head First Go.
(c) Ashraf Minhaj