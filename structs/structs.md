- [¿Qué es una estructura?](#1-%C2%BFqu%C3%A9-es-una-estructura)
  - [Declarar un tipo de estructura](#11-declarar-un-tipo-de-estructura)
  - [Creando una estructura](#12-creando-una-estructura)
  - [Obtener y rellenar campos de una estructura](#13-obtener-y-rellenar-campos-de-una-estructura)
  - [Inicializando una estructura](#14-inicializando-una-estructura)
  - [Estructura anónima](#15-estructura-an%C3%B3nima)
  - [Puntero a una estructura](#16-puntero-a-una-estructura)
  - [Campos anónimos](#17-campos-an%C3%B3nimos)
  - [Estructura anidada](#18-estructura-anidada)
  - [Campos promocionados](#19-campos-promocionados)
  - [Campos de función](#110-campos-de-funci%C3%B3n)
  - [Comparación de estructuras](#111-comparaci%C3%B3n-de-estructuras)
  - [Metadatos de campo de estructura](#112-metadatos-de-campo-de-estructura)
- [Referencias](#2-referencias)

# 1. Estructuras en Go

A diferencia de la Programación Orientada a Objetos tradicional, Go no tiene una arquitectura de clase-objeto. Más bien, tenemos estructuras que contienen estructuras de datos complejas.

Una `struct` puede compararse con `class` en el paradigma de la **Programación Orientada a Objetos**. Si no sabe qué es Programación Orientada a Objetos, imagine que `struct` es **una receta que declara los ingredientes y el tipo de cada ingrediente**.

Una estructura tiene diferentes campos del mismo o diferente tipo de datos. Si compara la estructura con una receta, los nombres de campo de la estructura se convierten en los ingredientes (**como la sal**) y los tipos de campo se convierten en el tipo de estos ingredientes (**como la sal de mesa**).

Una estructura se usa principalmente cuando necesita definir un `schema` (*esquema*) hecho de diferentes campos individuales (*propiedades*). Como una clase, podemos crear un objeto a partir de este esquema (*la `clase` es análoga al `esquema`*).

Dado que podemos instanciar una estructura, debe haber alguna distinción de nomenclatura entre la estructura y la instancia. Por lo tanto, el nombre tipo de estructura se usa para representar el esquema de estructura y `struct` o estructura se usa para representar la instancia.

Podemos decir que `ross` es un tipo de `Employee` (*tipo de estructura*) que tiene propiedades `firstName`, `LastName`, `salary` and `fullTime` (campos de estructura).

## 1.1 Declarar un tipo de estructura

Un `struct type` no es más que un esquema que contiene el plano de los datos que contendrá una estructura. Para simplificar las cosas, necesitamos crear un nuevo tipo derivado para que podamos referirnos fácilmente al tipo de estructura. Usamos la palabra clave `struct` para crear un nuevo tipo de estructura como se muestra en el siguiente ejemplo.

```go
type StructName struct {
  field1 fieldType1
  field2 fieldType2
}
```

En la sintaxis anterior, `StructName` es un tipo de estructura mientras que `field1` y `field2` son campos de tipo de datos `fieldType1` y `fieldType2` respectivamente.

Vamos a crear un tipo de estructura Empleado como discutimos pero con algunos campos reales.

```go
type Employee struct {
  firstName string
  lastName string
  salary int
  fullTime bool
}
```

También puede definir diferentes campos del mismo tipo de datos en la misma línea como hemos visto en la lección de [variables](../example-variables/variables.md).

```go
type Employee struct {
 firstName, lastName string
 salary int
 fullTime bool
}
```

## 1.2 Creando una estructura

Ahora que tenemos una `struct type` `Employee`, creemos una estructura `ross` a partir de ella. Dado que `Employee` es un tipo (*tipo de datos personalizado*), declarar una variable de tipo `Employee` seria lo mismo que de costumbre.

**Code**

```go
package main

import "fmt"

type Employee struct {
 firstName, lastName string
 salary              int
 fullTime            bool
}

func main() {
 var ross Employee
 fmt.Println(ross)
}
```

**Output**

```
{  0 false}
```

[Ejemplo](https://go.dev/play/p/c_Gf7YCXBJW)

El resultado del programa anterior puede parecerte raro, pero está dando el **zero value** de la estructura. Esto sucede porque hemos definido la variable `ross` del tipo de datos `Employee` pero no la hemos inicializado.

El **zero value** de una estructura es una estructura con todos los campos establecidos en sus propios valores cero. Por lo tanto,

- `string` tendrá el valor cero de "" (no se puede imprimir).
- `int` tendrá el valor cero de 0.
- `bool` tendrá el valor cero de falso.

> Cuando decimos estructura, nos referimos a la variable que contiene el valor del tipo de datos `Employee`. Por lo tanto, `Employee` es el `struct type`, por el contrario `ross` es una estructura mientras que la palabra clave `struct` es un `built-in type`. Si esto fuera así en el paradigma OOP, llamaríamos a `Employee` una clase y a `ross` un objeto.

## 1.3 Obtener y rellenar campos de una estructura

Obtener y configurar un campo de estructura es muy simple. Cuando se crea una variable de estructura, podemos acceder a sus campos usando `.` (*punto*) operador.

En el programa anterior, hemos creado una estructura `ross` que tiene 4 campos. Para asignar un valor al campo `firstName`, debe usar la sintaxis `ross.firstName = "ross"`. Démosle a `ross` algo de identidad en este ejemplo:

**Code**

```go
package main

import "fmt"

type Employee struct {
 firstName, lastName string
 salary              int
 fullTime            bool
}

func main() {
 var ross Employee
 ross.firstName = "ross"
 ross.lastName = "Bing"
 ross.salary = 1200
 ross.fullTime = true

 fmt.Println("ross.firstName =", ross.firstName)
 fmt.Println("ross.lastName =", ross.lastName)
 fmt.Println("ross.salary =", ross.salary)
 fmt.Println("ross.fullTime =", ross.fullTime)
}
```

**Output**

```
ross.firstName = ross
ross.lastName = Bing
ross.salary = 1200
ross.fullTime = true
```

[Ejemplo](https://go.dev/play/p/Fw_Z6uBCinZ)

## 1.4 Inicializando una estructura

En lugar de crear una estructura vacía (*simplemente declarando una variable con su `zero value`*) y luego asignando valores a sus campos individualmente, podemos crear una estructura con valores de campo inicializados en la misma sintaxis, como una variable.

**Code**

```go
package main

import "fmt"

type Employee struct {
 firstName, lastName string
 salary              int
 fullTime            bool
}

func main() {
 ross := Employee{
  firstName: "ross",
  lastName:  "Bing",
  fullTime:  true,
  salary:    1200,
 }

 fmt.Println(ross)
}
```

**Output**

```
{ross Bing 1200 true}
```

[Ejemplo](https://play.golang.org/p/FGJ0Ja4WM-F)

Hemos usado la notación abreviada (*usando la sintaxis `:=`*) para crear la variable `ross` para que Go pueda inferir el tipo `Employee` automáticamente. El orden de aparición de los campos de struct no importa, como puedes ver, hemos inicializado el campo `fullTime` antes que el campo de `salary`.

> La **coma** (*,*) es absolutamente necesaria después de la asignación de valor del último campo al crear una estructura usando la sintaxis anterior. De esta manera, Go no agregará un **punto y coma** justo después del último campo mientras compila el código.

También puede inicializar solo algunos campos de una estructura y dejar otros en sus `zero values`. En el siguiente ejemplo, el valor de la estructura `ross` será `{ross Bing 0 true}` ya que `salary`, de tipo `int`, tiene un `zero value` de 0.

```go
ross := Employee {
  firstName: "ross",
  lastName:  "Bing",
  fullTime:  true,
}
```

Hay otra forma de inicializar una estructura que no incluye **field name declarations** como la siguiente.

```go
ross := Employee{"Ross", "Geller", 1200, true}
```

La sintaxis anterior es perfectamente válida. Pero al crear una estructura sin declarar **field name declarations**, debe proporcionar todos los **field name declarations** en el orden en que aparecen en el tipo de estructura.

## 1.5 Estructura anónima

Una estructura anónima es una estructura sin un tipo de estructura derivada definida explícitamente. Hasta ahora, hemos creado el tipo de estructura `Employee` que infiere `ross`. Pero en el caso de una estructura anónima, no definimos ningún tipo de estructura derivada y creamos una estructura definiendo el tipo de estructura en línea y los valores iniciales de los campos de estructura en la misma sintaxis.

**Code**

```go
package main

import "fmt"

func main() {
 monica := struct {
  firstName, lastName string
  salary              int
  fullTime            bool
 }{
  firstName: "Monica",
  lastName:  "Geller",
  salary:    1200,
 }

 fmt.Println(monica)
}
```

**Output**

```
{Monica Geller 1200 false}
```

[Ejemplo](https://go.dev/play/p/Np7Y8LuOdql)

En el programa anterior, estamos creando una estructura `monica` sin definir un `type struct` derivado. Esto es útil cuando no desea reutilizar un tipo de estructura.

Entonces, estarías adivinando si `ross` es del tipo de `Employee`, ¿cuál es el tipo de `monica` aquí? Al usar la función `fmt.Printf` y la sintaxis de formato `%T`, obtenemos el siguiente resultado.

**Code**

```go
fmt.Printf("%T", monica)
```

**Output**

```
struct {firstName string; lastName string; salary int; fullTime bool}
```

¿Se ve raro? Pero no del todo. Porque así es como se vería `Employee` si no lo hubiéramos creado. La creación de un tipo derivado del tipo de estructura integrado nos brinda la flexibilidad de reutilizarlo sin tener que escribir una sintaxis compleja una y otra vez.

## 1.6 Puntero a una estructura

En lugar de crear una estructura, podemos crear un puntero que apunte al valor de una estructura, en una sola declaración. Esto ahorra un paso más para crear una estructura (*variable*) y luego crear un puntero a esa variable (*valor al puntero*).

La sintaxis para crear un puntero a una estructura es la siguiente.

```go
s := &StructType{...}
```

Vamos a crear un puntero `ross` que apunte a un valor de estructura.

**Code**

```go
package main

import "fmt"

type Employee struct {
 firstName, lastName string
 salary              int
 fullTime            bool
}

func main() {
 ross := &Employee{
  firstName: "ross",
  lastName:  "Bing",
  salary:    1200,
  fullTime:  true,
 }

 fmt.Println("firstName", (*ross).firstName)
}
```

**Output**

```
firstName ross
```

[Ejemplo](https://go.dev/play/p/fph03X-T-bu)

En el programa anterior, dado que `ross` es un puntero, necesitamos usar la `dereferencing syntax` de `*ross` para obtener el valor real de la estructura a la que apunta y usar `(*ross).firstName` para acceder al `firstName` de ese valor de estructura.

Usamos paréntesis alrededor de la `dereferencing syntax` del puntero en el programa anterior para que el compilador no se confunda entre `(*ross).firstName` y `*(ross.firstName)`.

Pero Go proporciona una sintaxis alternativa fácil para acceder a los campos. **Podemos acceder a los campos de un puntero de estructura sin `dereferencing` primero**. Go se encargará de desreferenciar un puntero under the hood.

```go
ross := &Employee {
  firstName: "ross",
  lastName:  "Bing",
  salary:    1200,
  fullTime:  true,
}

fmt.Println("firstName", ross.firstName) // ross is a pointer
```

## 1.7 Campos anónimos

Puede definir un tipo de estructura sin declarar ningún `field name`. Solo tiene que definir los tipos de datos de campo y Go utilizará las declaraciones de tipos de datos (*keywords*) como nombres de campo.

**Code**

```go
package main

import "fmt"

type Data struct {
 string
 int
 bool
}

func main() {
 sample1 := Data{"Monday", 1200, true}
 sample1.bool = false

 fmt.Println(sample1.string, sample1.int, sample1.bool)
}
```

**Output**

```
Monday 1200 false
```

[Example](https://go.dev/play/p/oIIIvnrTWdO)

En el programa anterior, hemos definido solo los `data types` en el tipo de estructura de datos. Go under the hood utilizará estos tipos de campo como el nombre de los campos. No hay diferencia alguna entre el tipo de estructura definido de esta manera y los tipos de estructura que hemos definido anteriormente.

Simplemente, en este caso, Go nos ayudó a crear nombres de campo automáticamente. Puede mezclar algunos campos anónimos con campos con nombre como se muestra a continuación.

```go
type Employee struct {
 firstName, lastName string
 salary              int
 bool                // anonymous field
}
```

## 1.8 Estructura anidada

Un campo de estructura puede ser de cualquier tipo de datos. Por lo tanto, es perfectamente legal tener un campo de estructura que contenga otra estructura. Por lo tanto, un campo de estructura puede tener un tipo de datos que sea un tipo de estructura. Cuando un campo de estructura tiene un valor de estructura, ese valor de estructura se denomina estructura anidada, ya que está anidado dentro de una estructura principal.

**Code**

```go
package main

import "fmt"

type Salary struct {
 basic     int
 insurance int
 allowance int
}

type Employee struct {
 firstName, lastName string
 salary              Salary
 bool
}

func main() {
 ross := Employee{
  firstName: "Ross",
  lastName:  "Geller",
  bool:      true,
  salary:    Salary{1100, 50, 50},
 }
 fmt.Println(ross)
}
```

**Output**

```
{Ross Geller {1100 50 50} true}
```

[Example](https://go.dev/play/p/uGsHK2ztQ5o)

Como puede ver en el ejemplo anterior, hemos creado un nuevo tipo de estructura `Salary` que define el `salary` de un empleado. Luego, modificamos el campo de `salary` del tipo de estructura `Employee` que ahora tiene un valor de tipo `Salary`.

Al crear una estructura de `ross` del tipo `Employee`, inicializamos todos los campos, incluso el campo de `salary`. Dado que el campo `salary` contiene la estructura de tipo `Salary`, podemos asignarle un valor de estructura. Hemos utilizado el método abreviado de excluir los nombres de los campos al inicializar la estructura salarial.

Normalmente, accedería a un campo de una estructura usando la sintaxis `struct.field`, como hemos visto antes. Puede acceder al campo de `salary` de la misma manera que `ross.salary` que devuelve una estructura. Luego **puede acceder (o actualizar) los campos de esta estructura anidada usando el mismo enfoque**, como por ejemplo, `ross.salary.basic.` Veamos esto en acción.

**Code**

```go
package main

import "fmt"

type Salary struct {
 basic     int
 insurance int
 allowance int
}

type Employee struct {
 firstName, lastName string
 salary              Salary
 bool
}

func main() {
 ross := Employee{
  firstName: "Ross",
  lastName:  "Geller",
  bool:      true,
  salary:    Salary{1100, 50, 50},
 }
 fmt.Println("Ross's basic salary", ross.salary.basic)
}
```

**Output**

```
Ross's basic salary 1100
```

[Ejemplo](https://go.dev/play/p/jVs385qTwZo)

## 1.9 Campos promocionados

Como hemos visto en la lección de [paquetes](../packages/packages.md), cualquier variable o tipo que comience con una letra mayúscula se exporta desde ese paquete. En el caso de las estructuras, nos aseguramos de que todas nuestras estructuras utilizadas en esta lección se exporten, por lo tanto, comienzan con una letra mayúscula, a saber. `Employee`, `Salary`, `Data`, etc.

Pero lo realmente genial de `struct` es que **también podemos controlar qué campos de una estructura exportada son visibles fuera del paquete (*o exportados*)**. Para exportar los nombres de campo de una estructura, debemos seguir el mismo enfoque de letras mayúsculas.

```go
type Employee struct {
 FirstName, LastName string
 salary int
 fullTime bool
}
```

En el tipo de estructura anterior `Employee`, `FirstName` y `LastName` son los dos únicos campos que se exportan o se ven fuera del paquete.

Vamos a crear un simple paquete `organization` con el nombre de paquete `org`. Podemos crear un archivo `WORKSPACE/src/org/employee.go` y colocar el siguiente código dentro de él. Este archivo exporta el tipo de estructura `Employee`.

```go
// employee.go
package org

type Employee struct {
 FirstName, LastName string
 salary              int
 fullTime            bool
}
```

En el paquete principal, podemos importar el tipo de estructura `Employee` como se muestra a continuación.

```go
// main.go
package main

import (
 "fmt"
 "org"
)

func main() {
  ross := org.Employee{
  FirstName: "Ross",
  LastName:  "Geller",
  salary:    1200
 }

 fmt.Println(ross)
}
```

El programa anterior no se compilará y el compilador arrojará el siguiente error.

```
unknown field 'salary' in struct literal of type org.Employee
```

Esto sucede porque el campo de `salary` no se exporta del tipo de estructura `Employee`. También tuvimos que usar `org.Employee` como tipo de estructura porque el tipo de empleado proviene del paquete org. Pero podemos crear un `derived type` (**tipo derivado**) en el paquete principal para simplificar las cosas.

```go
// main.go
package main

import (
 "fmt"
 "org"
)

type Employee org.Employee

func main() {
 ross := Employee{
  FirstName: "Ross",
  LastName:  "Geller",
 }

 fmt.Println(ross)
}
```

Por encima de los rendimientos del programa por debajo del resultado.

```
{Ross Geller 0 false}
```

¿Se ve raro? Quizás, porque no esperábamos el valor de los campos `salary` y `fullTime`. Cuando importamos cualquier estructura de otro paquete, obtenemos el tipo de estructura tal como es, solo que no tenemos ningún control sobre los campos no exportados. Esto **es útil cuando desea proteger algunos campos pero aún hacerlos útiles como valores predeterminados o constantes o quizás algo complejo**.

¿Qué sucederá en el caso de una estructura anidada?

- Una estructura anidada también debe declararse con una letra mayúscula para que otros paquetes puedan importarla.
- Los campos de estructura anidados que comienzan con una letra mayúscula se exportan.
- Si una **estructura anidada es anónima**, sus campos que comienzan con una letra mayúscula estarán disponibles como campos promocionados.

## 1.10 Campos de función

Si recuerda lo que hablamos en **la función como un tipo** y **la función como un valor** de la lección de [funciones](../functions/functions.md), puede adivinar que los campos de estructura también pueden ser funciones.

Así que vamos a crear un campo de función de estructura simple que devuelva el nombre completo de un empleado.

**Code**

```go
package main

import "fmt"

type FullNameType func(string, string) string

type Employee struct {
 FirstName, LastName string
 FullName            FullNameType
}

func main() {
 rossGeller := Employee{
  FirstName: "Ross",
  LastName:  "Geller",
  FullName: func(firstName string, lastName string) string {
   return firstName + " " + lastName
  },
 }

 fmt.Println(rossGeller.FullName(rossGeller.FirstName, rossGeller.LastName))
}
```

**Output**

```
Ross Geller
```

[Ejemplo](https://go.dev/play/p/U73YGQcWkwa)

En el programa anterior, hemos definido el `struct type` `Employee` que tiene dos campos `string` y un campo `function`. Solo por simplicidad, hemos creado un **derived function type** `FullNameType`.

Al crear la estructura `rossGeller`, debemos asegurarnos de que el campo `FullName` siga la sintaxis del tipo de función. En el caso anterior, le asignamos una función anónima. Dado que la sintaxis de esta función anónima y la declaración `FullNameType` coinciden, esto es perfectamente legal.

Luego simplemente ejecutamos la función `rossGeller.FullName` con dos argumentos de cadena `rossGeller.FirstName` y `rossGeller.LastName`.

> If you are wondering, why we need to pass properties of `rossGeller` (viz. `rossGeller.FirstName` and `rossGeller.LastName`) to `rossGeller.FullName` function because `FullName` field belongs to the same struct `rossGeller`, then you need to see the [methods](../methods/methods.md) lesson.

## 1.11 Comparación de estructuras

Dos estructuras son comparables si pertenecen al mismo tipo y tienen los mismos valores de campo.

**Code**

```go
package main

import "fmt"

type Employee struct {
 firstName, lastName string
 salary              int
}

func main() {
 ross := Employee{
  firstName: "Ross",
  lastName:  "Geller",
  salary:    1200,
 }

 rossCopy := Employee{
  firstName: "Ross",
  lastName:  "Geller",
  salary:    1200,
 }

 fmt.Println(ross == rossCopy)
}
```

**Output**

```
true
```

[Ejemplo](https://go.dev/play/p/AFkN-AxDSk5)

El programa anterior imprime `true` porque tanto `ross` como `rossCopy` pertenecen al mismo tipo de estructura `Employee` y tienen el mismo conjunto de valores de campo.

Sin embargo, **si una estructura tiene un tipo de campo que no se puede comparar**, por ejemplo, el `map` que no es comparable, entonces la estructura no será comparable.

Por ejemplo, si el tipo de estructura `Employee` tiene `leaves` como `map` de tipo de datos, no podríamos hacer la comparación anterior.

```go
type Employee struct {
  firstName, lastName string
  salary              int
  leaves              map[string]int
}
```

## 1.12 Metadatos de campo de estructura

`Struct` brinda una capacidad más para agregar metadatos a sus campos. Por lo general, se usa para proporcionar información de transformación sobre cómo se codifica o decodifica un campo de estructura de otro formato (o se almacena/recupera de una base de datos), pero puede usarlo para almacenar cualquier meta información que desee, ya sea para otro paquete o para su propio uso.

Esta meta información está definida por el literal de cadena (**lección de lectura de [strings](https://medium.com/rungo/string-data-type-in-go-8af2b639478)**) como se muestra a continuación.

```go
type Employee struct {
 firstName string `json:"firstName"`
 lastName  string `json:"lastName"`
 salary    int    `json: "salary"`
 fullTime  int    `json: "fullTime"`
}
```

En el ejemplo anterior, estamos utilizando el tipo de estructura `Employee` para fines de **codificación/descodificación JSON**. Lea más sobre la codificación/descodificación de JSON en el tutorial ["Trabajar con JSON"](../work-with-json/work-with-json.md)

# 2. Referencias

[Documentación oficial de golang acerca de structs](https://go.dev/ref/spec#Struct_types)
[Estructuras en Go (structs)](https://medium.com/rungo/structures-in-go-76377cc106a2)
