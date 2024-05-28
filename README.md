# Aprende Go en castellano 🇪🇸

<img src="./images/pet.jpeg" width="200">

Aprende Go en castellano con ejemplos tutoriales y ejercicios.

## Summary

* [Condicionales](./if-else-iterators/if-else-iterators.md#1-condicionales)
  * [Condicional if](./if-else-iterators/if-else-iterators.md#11-condicional-if)
  * [Condicional if-else](./if-else-iterators/if-else-iterators.md#12-condicional-if-else)
  * [Condicional if else if](./if-else-iterators/if-else-iterators.md#13-condicional-if-else-if)
    * [Estado inicial](./if-else-iterators/if-else-iterators.md#131-estado-inicial)
    * Ternary condition
  * [Condicional switch](./if-else-iterators/if-else-iterators.md#14-condicional-switch)
    * The syntax of the switch statement
    * [Default case](./if-else-iterators/if-else-iterators.md#141-default-case)
    * [Múltiples valores en el case](./if-else-iterators/if-else-iterators.md#142-multiples-valores-en-el-case)
    * [Inicial statement](./if-else-iterators/if-else-iterators.md#143-inicial-statement)
    * [Expressionless switch statement](./if-else-iterators/if-else-iterators.md#144-expressionless-switch-statement)
    * [Fallthrough statement](./if-else-iterators/if-else-iterators.md#145-fallthrough-statement)
* [Iteradores](./if-else-iterators/if-else-iterators.md#2-Iteradores)
  * [Bucles for](./if-else-iterators/if-else-iterators.md#21-bucles-for)
    * [Sintaxis del bucle for](./if-else-iterators/if-else-iterators.md#211-sintaxis-del-bucle-for)
    * [Variantes del bucle for](./if-else-iterators/if-else-iterators.md#212-variantes-del-bucle-for)
      * [Opcional init statment](./if-else-iterators/if-else-iterators.md#2121-opcional-init-statment)
      * [Opcional post statment](./if-else-iterators/if-else-iterators.md#2122-opcional-post-statment)
      * [Opcional init y statment](./if-else-iterators/if-else-iterators.md#2123-opcional-init-y-post-statment)
      * [Sin ningún statment](./if-else-iterators/if-else-iterators.md#2124-sin-ningun-statment)
      * [El break statment](./if-else-iterators/if-else-iterators.md#2125-el-break-statement)
      * [El continue statment](./if-else-iterators/if-else-iterators.md#2126-el-continue-statement)
      * [El return statment](./if-else-iterators/if-else-iterators.md#2127-el-return-statement)
      * [Range](./if-else-iterators/if-else-iterators.md#2128-range)
        * [Range sobre un array](./if-else-iterators/if-else-iterators.md#21281-range-sobre-un-array)
        * [Range sobre un map](./if-else-iterators/if-else-iterators.md#21282-range-sobre-un-map)
          * [Range sobre un map usando keys](./if-else-iterators/if-else-iterators.md#212821-range-sobre-un-map-usando-keys)
          * [Range sobre un map usando key value](./if-else-iterators/if-else-iterators.md#212822-range-sobre-un-map-usando-key-value)
* [Punteros en Go](./pointers/pointers.md#1-punteros-en-go)
  * [Cómo acceder a la dirección de memoria de una variable?](./pointers/pointers.md#12-cómo-acceder-a-la-dirección-de-memoria-de-una-variable)
  * [Que es un puntero?](./pointers/pointers.md#13-que-es-un-puntero)
  * [Desreferenciar un puntero](./pointers/pointers.md#14-desreferenciar-un-puntero)
  * [Cambiar el valor de la variable usando un puntero](./pointers/pointers.md#15-cambiar-el-valor-de-la-variable-usando-un-puntero)
  * [La función new](./pointers/pointers.md#16-la-función-new)
  * [Pasar un puntero a una función](./pointers/pointers.md#17-pasar-un-puntero-a-una-función)
  * [Aritmética de punteros](./pointers/pointers.md#18-aritmética-de-punteros)
* [Estructuras en go (structs)](./structs/structs.md)
  * [¿Qué es una estructura?](./structs/structs.md#1-qu%C3%A9-es-una-estructura)
  * [Declarar un tipo de estructura](./structs/structs.md#11-declarar-un-tipo-de-estructura)
  * [Creando una estructura](./structs/structs.md#12-creando-una-estructura)
  * [Obtener y establecer campos de estructura](./structs/structs.md#13-obtener-y-establecer-campos-de-estructura)
  * [Inicializando una estructura](./structs/structs.md#14-inicializando-una-estructura)
  * [Estructura anónima](./structs/structs.md#15-estructura-an%C3%B3nima)
  * [Puntero a una estructura](./structs/structs.md#16-puntero-a-una-estructura)
  * [Campos anónimos](./structs/structs.md#17-campos-an%C3%B3nimos)
  * [Estructura anidada](./structs/structs.md#18-estructura-anidada)
  * [Campos promocionados](./structs/structs.md#19-campos-promocionados)
  * [Campos de función](./structs/structs.md#110-campos-de-funci%C3%B3n)
  * [Comparación de estructuras](./structs/structs.md#111-comparaci%C3%B3n-de-estructuras)
  * [Metadatos de campo de estructura](./structs/structs.md#112-metadatos-de-campo-de-estructura)
* [Funciones en Go](./functions/functions.md)
  * [Qué es una función?](./functions/functions.md#11-qu%C3%A9-es-una-funci%C3%B3n)
  * [Convención de nombres para funciones](./functions/functions.md#12-convenci%C3%B3n-de-nombres-para-funciones)
  * [Parámetros en funciones](./functions/functions.md#13-par%C3%A1metros-en-funciones)
  * [Valor de retorno](./functions/functions.md#14-valor-de-retorno)
  * [Multiples valores de retorno](./functions/functions.md#15-multiples-valores-de-retorno)
  * [Valores de retorno nombrados](./functions/functions.md#16-valores-de-retorno-nombrados)
  * [Función recursiva](./functions/functions.md#17-funci%C3%B3n-recursiva)
  * [`defer` keyword](./functions/functions.md#18-defer-keyword)
  * [Función como tipo](./functions/functions.md#19-funci%C3%B3n-como-tipo)
  * [Función como valor (función anónima)](./functions/functions.md#110-funci%C3%B3n-como-valor-funci%C3%B3n-an%C3%B3nima)
  * [Función como valor (función anónima)](./functions/functions.md#111-expresi%C3%B3n-de-funci%C3%B3n-invocada-inmediatamente-iife)
* [Métodos en Go](./methods/methods.md#1-métodos-en-go)
  * [Que es un método?](./methods/methods.md#11-que-es-un-método)
  * [Métodos con el mismo nombre](./methods/methods.md#12-métodos-con-el-mismo-nombre)
  * [Pointer receivers](./methods/methods.md#13-pointer-receivers)
    * [Métodos de llamada con receiver de puntero en valores](./methods/methods.md#131-métodos-de-llamada-con-receiver-de-puntero-en-valores)
  * [Métodos en estructuras anidadas](./methods/methods.md#14-métodos-en-estructuras-anidadas)
    * [Métodos en estructuras anidadas](example-methods/methods.md#141-métodos-en-estructuras-anidadas)
    * [Estructuras anidadas anónimamente](./methods/methods.md#142-estructuras-anidadas-anónimamente)
    * [Métodos promocionados](./methods/methods.md#143-métodos-promocionados)
  * [Métodos pueden aceptar ambos, punteros y valores](./methods/methods.md#15-métodos-pueden-aceptar-ambos-punteros-y-valores)
  * [Métodos en no struct type](./methods/methods.md#16-métodos-en-no-struct-type)
* [Interfaces en Go](./interfaces/interfaces.md)
  * [Que es una interfaz?](./interfaces/interfaces.md#1-que-es-una-interfaz)
  * [Declaración de interfaz](./interfaces/interfaces.md#12-declaración-de-interfaz)
  * [Implementación de interfaz](./interfaces/interfaces.md#13-implementación-de-interfaz)
  * [Interfaz vacía](./interfaces/interfaces.md#14-interfaz-vacía)
  * [Interfaces multiples](./interfaces/interfaces.md#15-interfaces-multiples)
  * [Type assertion](./interfaces/interfaces.md#16-type-assertion)
  * [Type switch](./interfaces/interfaces.md#17-type-switch)
  * [Interfaces embebidas](./interfaces/interfaces.md#18-interfaces-embebidas)
  * [Pointer vs Value receiver](./interfaces/interfaces.md#19-pointer-vs-value-receiver)
  * [Comparación de interfaces](./interfaces/interfaces.md#110-comparación-de-interfaces)
  * [Uso de interfaces](./interfaces/interfaces.md#111-uso-de-interfaces)
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
  * [Leap Year](./katas/leapyear/leapyear.md)
  * [String Calculator](./katas/stringcalculator/stringcalculator.md)
