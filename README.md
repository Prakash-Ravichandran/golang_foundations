## What does a package main mean ?


1. Executable packages: Defines a package that can be compiled and executed. Must have a func called Main.
2. Reusable packages: Defines a package that can be used as a dependency (helper code).

## what does import "fmt" mean ?

- It means the main function has access to the fmt package. Simply linking the fmt package to the main package.

## what does func thing mean ?

- Its the definition that we are creating function with a keyword 'func'

## how does the file organization in go looks ?

1. Package declaration.
2. Import statements.
3. Function definitions.


# Section 3

## Variable Declaration

```c.go
package main

import "fmt"


func main(){
   // var card string = "Ace of Spades"
   card := "Ace of Spades"
   card = "Five of Diamonds"
}
```

## Function Declaration and return type


```main.go
/** 
Function declaration and return type.
**/

package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
```

## Slices and For loops

[commit](https://github.com/Prakash-Ravichandran/golang_foundations/commit/d2266aca122608cdb46cb1217050bfd440e6e5c8) 

```c.go
func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades") // add new card to the slice

	for i, card := range cards { 
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
```

## OO approach vs Go Approach


OO Approach:

- Class deck has a string with member functions like print, share, saveDeck.

Go Approach:

- Function with 'deck' as a 'receiver' (A function with a receiver is like a 'method' - a function that belongs to an instance)

Similar to Javascript, python we can use this/self for the receiver variable.

receiver variable d = cards or this or self

```c.go
func (d deck) print() {    
	for i, card := range this/self {        
		fmt.Println(i, card)
	}
}
```

## Custom Type Definitions and Receiver Functions - [commit](https://github.com/Prakash-Ravichandran/golang_foundations/commit/ce3ec700d9a991a622f85cffb47cf620c6a87a78#diff-91d7af91df1707302d8dfe290cf342b8cbcba0b563c429d1b05ca8e1843cd051)

## Multiple Return Values

- [Multiple Return Values](https://github.com/Prakash-Ravichandran/golang_foundations/commit/ad1405f2b6f38e2dd936c8bb1a7e3f491e791f1a)


## Byte Slices

 "Hi there!" is a human readable code, byte slice is machine readable code. conversion of string to decimal code - [asciitable.com](https://www.asciitable.com/)


```c.go


  "Hi there!"  string

  [72 105 32 116 104 101 114 101 33] byte slice

  ```

  ## Deck to String

```c.go

  deck -> []string -> string -> []byte

``` 

## Saving File to Hard Drive

[save file to hard drive](https://github.com/Prakash-Ravichandran/golang_foundations/commit/a4a94fc49588aaa341bd40e7ce1faadee5de6aa7)

