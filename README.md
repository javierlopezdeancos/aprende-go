# My Go Examples

<img src="./images/pet.jpeg" width="200">

Algunos ejemplos sobre tutoriales y ejercicios en Go.

## Summary

* [Condicionales e Iteradores](./example-ifelse-iterations/ifelse-iterators.md)
  * [Condicionales](./example-ifelse-iterations/ifelse-iterators.md#1-condicionales)
    * [Condicional if](./example-ifelse-iterations/ifelse-iterators.md#11-condicional-if)
    * [Condicional if-else](./example-ifelse-iterations/ifelse-iterators.md#12-condicional-if-else)
    * [Condicional if else if](./example-ifelse-iterations/ifelse-iterators.md#13-condicional-if-else-if)
      * [Estado inicial](./example-ifelse-iterations/ifelse-iterators.md#131-estado-inicial)
      * Ternary condition
    * [Condicional switch](./example-ifelse-iterations/ifelse-iterators.md#14-condicional-switch)
      * The syntax of the switch statement
      * [Default case](./example-ifelse-iterations/ifelse-iterators.md#141-default-case)
      * [Múltiples valores en el case](./example-ifelse-iterations/ifelse-iterators.md#142-multiples-valores-en-el-case)
      * [Inicial statement](./example-ifelse-iterations/ifelse-iterators.md#143-inicial-statement)
      * [Expressionless switch statement](./example-ifelse-iterations/ifelse-iterators.md#144-expressionless-switch-statement)
      * [Fallthrough statement](./example-ifelse-iterations/ifelse-iterators.md#145-fallthrough-statement)
  * [Iteradores](./example-ifelse-iterations/ifelse-iterators.md#2-Iteradores)
    * [Bucles for](./example-ifelse-iterations/ifelse-iterators.md#21-bucles-for)
      * [Sintaxis del bucle for](./example-ifelse-iterations/ifelse-iterators.md#211-sintaxis-del-bucle-for)
      * [Variantes del bucle for](./example-ifelse-iterations/ifelse-iterators.md#212-variantes-del-bucle-for)
        * [Opcional init statment](./example-ifelse-iterations/ifelse-iterators.md#2121-opcional-init-statment)
        * [Opcional post statment](./example-ifelse-iterations/ifelse-iterators.md#2122-opcional-post-statment)
        * [Opcional init y statment](./example-ifelse-iterations/ifelse-iterators.md#2123-opcional-init-y-post-statment)
        * [Sin ningún statment](./example-ifelse-iterations/ifelse-iterators.md#2124-sin-ningun-statment)
        * [El break statment](./example-ifelse-iterations/ifelse-iterators.md#2125-el-break-statement)
        * [El continue statment](./example-ifelse-iterations/ifelse-iterators.md#2126-el-continue-statement)
        * [El return statment](./example-ifelse-iterations/ifelse-iterators.md#2127-el-return-statement)
        * [Range](./example-ifelse-iterations/ifelse-iterators.md#2128-range)
          * [Range sobre un array](./example-ifelse-iterations/ifelse-iterators.md#21281-range-sobre-un-array)
          * [Range sobre un map](./example-ifelse-iterations/ifelse-iterators.md#21282-range-sobre-un-map)
            * [Range sobre un map usando keys](./example-ifelse-iterations/ifelse-iterators.md#212821-range-sobre-un-map-usando-keys)
            * [Range sobre un map usando key value](./example-ifelse-iterations/ifelse-iterators.md#212822-range-sobre-un-map-usando-key-value)
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
* [Anatomía de funciones en Go](./example-functions/functions.md)
  * [Qué es una función?](./example-functions/functions.md#11-qu%C3%A9-es-una-funci%C3%B3n)
  * [Convención de nombres para funciones](./example-functions/functions.md#12-convenci%C3%B3n-de-nombres-para-funciones)
  * [Parámetros en funciones](./example-functions/functions.md#13-par%C3%A1metros-en-funciones)
  * [Valor de retorno](./example-functions/functions.md#14-valor-de-retorno)
  * [Multiples valores de retorno](./example-functions/functions.md#15-multiples-valores-de-retorno)
  * [Valores de retorno nombrados](./example-functions/functions.md#16-valores-de-retorno-nombrados)
  * [Función recursiva](./example-functions/functions.md#17-funci%C3%B3n-recursiva)
  * [`defer` keyword](./example-functions/functions.md#18-defer-keyword)
  * [Función como tipo](./example-functions/functions.md#19-funci%C3%B3n-como-tipo)
  * [Función como valor (función anónima)](./example-functions/functions.md#110-funci%C3%B3n-como-valor-funci%C3%B3n-an%C3%B3nima)
  * [Función como valor (función anónima)](./example-functions/functions.md#111-expresi%C3%B3n-de-funci%C3%B3n-invocada-inmediatamente-iife)
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
* [Trabajando con JSON](./example-work-with-json/work-with-json.md#trabajando-con-json)
  * [Referencias](./example-work-with-json/work-with-json.md#referencias)
  * [Introducción](./example-work-with-json/work-with-json.md#introduccion)
  * [Codificando JSON](./example-work-with-json/work-with-json.md#codificando-json)
    * [Manejo de tipos de datos](./example-work-with-json/work-with-json.md#manejo-de-tipos-de-datos)
    * [Tipos de datos abstractos](./example-work-with-json/work-with-json.md#tipos-de-datos-abstractos)
    * [Conversión de tipos de datos](./example-work-with-json/work-with-json.md#conversion-de-tipos-de-datos)
    * [Codificando usando structure tags](./example-work-with-json/work-with-json.md#codificando-usando-structure-tags)
    * [Codificar trabajando con maps](./example-work-with-json/work-with-json.md#codificar-trabajando-con-maps)
  * [Decodificando JSON](./example-work-with-json/work-with-json.md#decodificando-json)
    * [Manejando estructuras de datos complejas](./example-work-with-json/work-with-json.md#manejando-estructuras-de-datos-complejas)
    * [Campos promocionados](./example-work-with-json/work-with-json.md#campos-promocionados)
    * [Decodificando usando structure tags](./example-work-with-json/work-with-json.md#decodificando-usando-structure-tags)
    * [Decodificar trabajando con maps](./example-work-with-json/work-with-json.md#decodificar-trabajando-con-maps)
    * [Usando Unmarshaler and TextUnmarshaler](./example-work-with-json/work-with-json.md#usando-unmarshaler-and-textunmarshaler)
  * [Codificador y Decodificador](./example-work-with-json/work-with-json.md#codificador-y-decodificador)
    * [Codificador](./example-work-with-json/work-with-json.md#codificador)
    * [Decodificador](./example-work-with-json/work-with-json.md#codificador)
* Katas
  [Leap Year](./katas/leapyear/leapyear.md)
  [String Calculator](./katas/stringcalculator/stringcalculator.md)
