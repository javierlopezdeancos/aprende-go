# My Go Examples

<img src="./images/pet.jpeg" width="200">

Algunos ejemplos sobre tutoriales y ejercicios en Go.

## Summary

* [Condicionales e Iteradores](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md)
  * [Condicionales](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#1-condicionales)
    * [Condicional if](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#11-condicional-if)
    * [Condicional if-else](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#12-condicional-if-else)
    * [Condicional if else if](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#13-condicional-if-else-if)
      * [Estado inicial](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#131-estado-inicial)
      * Ternary condition
    * [Condicional switch](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#14-condicional-switch)
      * The syntax of the switch statement
      * [Default case](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#141-default-case)
      * [Múltiples valores en el case](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#142-multiples-valores-en-el-case)
      * [Inicial statement](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#143-inicial-statement)
      * [Expressionless switch statement](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#144-expressionless-switch-statement)
      * [Fallthrough statement](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#145-fallthrough-statement)
  * [Iteradores](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#2-Iteradores)
    * [Bucles for](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#21-bucles-for)
      * The for loop syntax
      * Variants of the for loop
        * [Opcional init statment](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#211-opcional-init-statment)
        * [Opcional post statment](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#212-opcional-post-statment)
        * [Opcional init y statment](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#213-opcional-init-y-post-statment)
        * [Sin ningún statment](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#214-sin-ningun-statment)
        * [El break statment](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#215-el-break-statement)
        * [El continue statment](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#216-el-continue-statement)
        * [El return statment](https://github.com/javierlopezdeancos/my-go-examples/blob/master/example-ifelse-iterations/ifelse-iterators.md#217-el-return-statement)
* [Structures in Go (structs)](https://medium.com/rungo/structures-in-go-76377cc106a2)
  * Declaring a struct type
  * Creating a struct
    * Getting and setting struct fields
  * Initializing a struct
  * Anonymous struct
  * Pointer to a struct
  * Anonymous fields
  * Nested struct
  * Promoted fields
  * Nested interface
  * Exported fields
  * Function fields
  * Struct comparison
  * Struct field meta-data
* [The anatomy of Functions in Go](https://medium.com/rungo/the-anatomy-of-functions-in-go-de56c050fe11)
  * What is a function
  * Function name convention
  * Function parameters
  * Return value
  * Multiple return values
* [Anatomy of methods in Go](https://medium.com/rungo/anatomy-of-methods-in-go-f552aaa8ac4a)
  * What is a method?
  * Methods with the same name
  * Pointer receivers
    * Calling methods with pointer receiver on values
  * Methods on nested struct
    * Methods on nested struct
    * Anonymously nested struct
    * Promoted methods
  * Methods can accept both pointer and value
  * Methods on non-struct type
* [Interfaces in Go](https://medium.com/rungo/interfaces-in-go-ab1601159b3a)
  * What is an interface?
  * Declaring interface
  * Implementing interface
  * Empty interface
  * Multiple interfaces
  * Type assertion
  * Type switch
  * Embedding interfaces
  * Pointer vs Value receiver
  * Interface comparison
  * Use of interfaces
* [Introduction to Streams and Buffers](https://medium.com/rungo/introduction-to-streams-and-buffers-d148c0cda0ad)
  * What are a stream and a buffer
  * Reading from a Data Source
    * io.Reader
    * strings.NewReader
    * ioutil.ReadAll
    * io.ReadFull
    * io.LimitReader
  * Writing to a Data Store
    * io.Writer
    * io.WriteString
    * Standard I/O Streams
  * Closing I/O Operations
  * Transferring Data between streams
    * io.Copy
    * io.Pipe
  * Buffered streams
