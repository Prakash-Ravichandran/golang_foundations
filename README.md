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

## Read From file

[readFromFile](https://github.com/Prakash-Ravichandran/golang_foundations/commit/020b97fdcc7cd59de14fc52335892464b292fc61)

## Generate New Seed for random number generation

[Gen new seed](https://github.com/Prakash-Ravichandran/golang_foundations/commit/154bead6b77eb22a7b8f5651553e9ca134d412e8)

# Structs

[commit](https://github.com/Prakash-Ravichandran/golang_foundations/commit/9fb88fe6435e70a84e1ef9eee62d17eff1fd7349)

- A collection of related properties.
- Similar to objects in Javascript.
- A short version of stucture from Data Structure.

- Three of way of using structs in go

```go

type person struct {
	firstName string
	lastName  string
}

// first way
func main() {
	alex := person{"alex", "anderson"}
}

// second way
func main() {
	alex := person{firstName: "alex", lastName: "anderson"}
}

// third way
func main() {
   var alex person
   alex.firstName = "Alex"
   alex.lastName = "Anderson"
}

```

### Embedding Structs

- [commit](https://github.com/Prakash-Ravichandran/golang_foundations/commit/2dd4db7895e3395c5a3dd4ed3da23e86c5a3e253)

### Pointer

[commit](https://github.com/Prakash-Ravichandran/golang_foundations/commit/3dc136beef3ac044971ae155357e4e745f282a9d)

- Turn address into value using: \*addressTurn value into address using:  &values
- Go is a "pass by value" language, which _always_ copies arguments that are passed to a function.

```go

- here *person is a type description that receiver variable should be pointer to person
- *pointerToPerson - give the value that the pointer is pointing to. 

func (pointerToPerson *person) updateName() {
  (*pointerToPerson).firstName = "Alex"
}
```

### Pointer Shortcut in Go

[commit](https://github.com/Prakash-Ravichandran/golang_foundations/commit/5f635100a9e73db64cff6f1f97999c86be203bfe)

- GO can infer the pointer to variable from the receiver function

### Go: Value Types vs Reference Types

|                  | **Value Types**                            | **Reference Types**                              |
| ---------------- | ------------------------------------------ | ------------------------------------------------ |
| **Types**        | `int`, `float`, `string`, `bool`, `struct` | `slice`, `map`, `channel`, `pointer`, `function` |
| **In Functions** | Use pointers to modify these               | No pointers needed to modify these               |

### Maps

[different ways of defining maps](https://github.com/Prakash-Ravichandran/golang_foundations/commit/b022ef42060a720ff4b828f611b70b6047519b34)

- Maps are similar to Objects in Javascript, Hash in Ruby, Dict in Python
- Structs in Go are also similar to Objects in Javascript, Hash in Ruby, Dict in Python.
- They differ mainly
-

#### Difference between structs and maps

##### Map

| Property        | Details                                             |
| --------------- | --------------------------------------------------- |
| 🔑 Key Types    | All keys must be the same type                      |
| 📦 Value Types  | All values must be the same type                    |
| 🔁 Iteration    | Keys are indexed — we can iterate over them         |
| 🏗️ Use Case     | Use to represent a collection of related properties |
| ⏱️ Compile Time | Don't need to know all the keys at compile time     |
| 🧠 Memory       | Reference Type!                                     |

---

##### Struct

| Property        | Details                                                       |
| --------------- | ------------------------------------------------------------- |
| 📦 Value Types  | Values can be of different type                               |
| 🔑 Key Indexing | Keys don't support indexing                                   |
| 🏗️ Use Case     | Use to represent a "thing" with a lot of different properties |
| ⏱️ Compile Time | You need to know all the different fields at compile time     |
| 🧠 Memory       | Value Type!                                                   |

---

##### Quick Comparison

| Feature                     | Map                               | Struct                                   |
| --------------------------- | --------------------------------- | ---------------------------------------- |
| Key types                   | Must be the same                  | N/A (named fields)                       |
| Value types                 | Must be the same                  | Can be different                         |
| Key indexing                | ✅ Supported                      | ❌ Not supported                         |
| Know fields at compile time | ❌ Not required                   | ✅ Required                              |
| Memory model                | Reference Type                    | Value Type                               |
| Best for                    | Collections of related properties | A "thing" with many different properties |

## 57. Purpose of Interface

- Every variable in Go should have a type.
- Every function's argument should have type.

SO Does that mean..

- Every function we ever has to be rewritten to accomodate different types even when the logic inside those functions are identical ?
- Go doesn't function overloading like other languages.

[Without Interface](https://github.com/Prakash-Ravichandran/golang_foundations/commit/c5c11066af6c08eff096f12d1d4535ec9a598983)

### A question in example of interface:

- how the printGreeting function accepts both english and spanishbot struct, printGreeting should accept of type bot right ?

- Ans: `englishBot` and `spanishBot` satisfied the requirement of a bot interface in their receiver functions

```go

func (englishBot) getGreeting() string { // getGreeting here defined with receiver as englishBot
}

func (spanishBot) getGreeting() string { // getGreeting here defined with receiver as spanishBot
}

// Interfaces are implicit in Go, we don't have to do like

// type englishBot implements Bot - go takes care of it implicity

```

### Concrete Type vs Interface Type

Concrete Type: Can create a value out the type. Ex: map, struct, int, float, englishBot
Interface Type: Cannot create a value out the type

```go

type englishBot struct {}

eb:= englishBot{}

type myString string

s: myString

```

### Custom Writer Implementaion

[custom writer](https://github.com/Prakash-Ravichandran/golang_foundations/commit/7449e7a87b8e1126cf32d6994c20d704597e236d)

Question: Why we need define a custom implemenation of write() in go needs a Struct.

In Go, methods must be associated with a **defined type**. You cannot simply create a standalone function named `Write` and expect it to satisfy an interface because the Go compiler needs to know _what_ that function belongs to.

### Why the `struct` is required

When you define an interface like `io.Writer`, it is defined as:

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

For a type to satisfy this interface, it must "possess" a `Write` method. In Go, the only way to attach a method to a type is through a **receiver**:

```go
func (cw customWriter) Write(p []byte) (n int, err error) { ... }
```

The `(cw customWriter)` part is the **receiver**. It tells the compiler: _"This specific function belongs to the `customWriter` type."_ Without the `struct` (or another named type), there is no "anchor" for that method to attach to.

---

### Does it _have_ to be a struct?

**No, it doesn't have to be a `struct`!** In Go, you can attach methods to **any user-defined type**. While `struct` is the most common way to do this because it holds data, you can use other underlying types as well.

#### Example: Using a custom string type

If the implementation doesn't need to hold complex data, could use a custom type based on `string` or even `int`:

```go
type myWriter string

func (mw myWriter) Write(p []byte) (n int, err error) {
    fmt.Println("I am a:", mw)
    return len(p), nil
}

func main() {
    var lw myWriter = "My custom string writer"
    // This works perfectly because myWriter has a Write method!
    io.Copy(lw, res.Body)
}
```

### The "Big Picture"

1.  **`interface`** = The set of rules (the behavior).
2.  **`type` (struct, string, etc.)** = The container that holds the logic.
3.  **`method`** = The actual logic that fulfills the rule.

The **`type`** to act as the identity for the **`method`**. When passed `lw` to `io.Copy`, Go looks at the type of `lw`, sees that it has a `Write` method, and confirms that it matches the `io.Writer` interface requirements.

### Assignment

[assignment](https://github.com/Prakash-Ravichandran/golang_foundations/commit/e5431b7ff2fc41ea517b18f37b6cd8325656fd7f)

## Channels and Go Routines

- Concurrency is not parallelism

Creating 5 routines and the main routine waiting for the only one channel response it a blocking. Called Blocking the channel [code](https://github.com/Prakash-Ravichandran/golang_foundations/commit/9ab9c2b905e89f66e90a382ed8430516a52004ba)

- Repeatedly call the links with new routine created as soon as we receive the link in channel.
