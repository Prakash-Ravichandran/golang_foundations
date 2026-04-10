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