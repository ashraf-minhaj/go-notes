# go-notes
 
## To run (basic)
 * Init - `go mod init program_name`
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

## Build Go apps
 * compile the source code intor binary file - `go build file_name`
 * run the exe on windows, or just file name on mac/linux. `./file_name`

## Notes:
 * go is a statically typed language. If you type wrong type of value in the wrong place, go will let you know.



Feel free to learn and share.
Refeerence: head first go.
(c) Ashraf Minhaj