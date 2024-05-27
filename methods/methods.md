- [Métodos en Go](#1-m%C3%A9todos-en-go)
  - [Que es un método?](#11-que-es-un-m%C3%A9todo)
  - [Métodos con el mismo nombre](#12-m%C3%A9todos-con-el-mismo-nombre)
  - [Pointer receivers](#13-pointer-receivers)
    - [Métodos de llamada con receiver de puntero en valores](#131-m%C3%A9todos-de-llamada-con-receiver-de-puntero-en-valores)
  - [Métodos en estructuras anidadas](#14-m%C3%A9todos-en-estructuras-anidadas)
    - [Métodos en estructuras anidadas](#141-m%C3%A9todos-en-estructuras-anidadas)
    - [Estructuras anidadas anónimamente](#142-estructuras-anidadas-an%C3%B3nimamente)
    - [Métodos promocionados](#143-m%C3%A9todos-promocionados)
  - [Métodos pueden aceptar ambos, punteros y valores](#15-m%C3%A9todos-pueden-aceptar-ambos-punteros-y-valores)
  - [Métodos en no struct type](#16-m%C3%A9todos-en-no-struct-type)
- [Referencias](#2-referencias)

# 1. Métodos en Go

Go no es compatible con el paradigma orientado a objetos, pero la `struct` se asemeja a la arquitectura de clases. Para agregar métodos a una estructura, necesitamos usar funciones con un `receiver` (*receptor*).

Go no proporciona clases, pero podemos usar estructuras para crear objetos como hemos aprendido en el capitulo de [structs](../structs/structs.md). Pero en *Object Oriented Programming*, las clases tienen `properties` (campos) así como `behaviors` (métodos) y hasta ahora solo hemos aprendido acerca de las propiedades de una estructura que son structure fields.

> Behavior es una acción que un objeto puede realizar. Por ejemplo, `Dog` es un tipo de `Animal` y `Dog` puede `bark`(*ladrar*). Por lo tanto, ladrar es un comportamiento de la clase `Dog`. Por lo tanto, cualquier objeto (instancia) de la clase `Dog` tendrá este comportamiento.

Hemos visto en la lección de [structs](../structs/structs.md#campos-de-funci%C3%B3n), especialmente en la sección de campo de función, que un campo de estructura también puede ser una función. Podemos agregar un campo `bark` de función de tipo que no toma argumentos y devuelve un string `woof woof!`. Esta podría ser una forma de agregar métodos a la estructura.

Pero esto no se adhiere al concepto Object Oriented Programing ya que los campos de estructura no tienen idea de la estructura a la que pertenecen. Por lo tanto, los métodos vienen al rescate.

# 1.1 Que es un método?

En la sección anterior de [structs](../structs/structs.md), jugamos con los campos de función de una estructura, por lo tanto, el concepto de método será muy fácil de entender.

Un método no es más que una función, pero pertenece a cierto tipo. Un método se define con una sintaxis ligeramente diferente a la de una función normal. Requería un parámetro adicional conocido como receiver (*receptor*), que es un tipo al que pertenece la función. De esta forma, un método (función) puede acceder a las propiedades del receptor al que pertenece (como campos de una estructura).

Escribamos un programa para obtener el `full name` de una estructura `Employee` usando una simple función.

**Code**

```go
package main

import "fmt"

type Employee struct {
 FirstName, LastName string
}

func fullName(firstName string, lastName string) (fullName string) {
 fullName = firstName + " " + lastName
 return
}

func main() {
 e := Employee{
  FirstName: "Ross",
  LastName:  "Geller",
 }

 fmt.Println(fullName(e.FirstName, e.LastName))
}
```

**Output**

```
Ross Geller
```

[Ejemplo](https://go.dev/play/p/ANE_IN9cmk4)

En el programa anterior, hemos creado una estructura simple de tipo `Employee` que tiene dos campos string, `FirstName` y `LastName`. Luego definimos la función `fullName` que toma dos strings de argumentos y devuelve un string. La función `fullName` devuelve el nombre completo de un empleado concatenando estos dos strings.

Luego creamos una estructura de tipo `Empoyee` al proporcionar los valores de los campos `FirstName` y `LastName`. Para obtener el nombre completo del empleado `e`, usamos la función `fullName` y proporcionamos los argumentos apropiados.

Esto funciona, pero lo malo es que cada vez que necesitamos obtener el nombre completo de un empleado (y podría haber miles), debemos pasar los valores de nombre y apellido a la función de nombre completo manualmente.

Un **método** puede resolver este problema fácilmente. Para convertir una función al método, solo necesitamos un parámetro de receptor adicional en la definición de la función. La sintaxis para definir un método es la siguiente.

```go
func (r Type) functionName(...Type) Type {
    ...
}
```

De la sintaxis anterior, podemos decir que el método y la función tienen la misma sintaxis excepto por una declaración de argumento receiver (`r Type`) justo antes del nombre de la función. `Type` es cualquier tipo legal en Go y los argumentos de función y los valores devueltos son opcionales.

Vamos a crear el método `fullName` usando la sintaxis anterior.

**Code**

```go
package main

import "fmt"

type Employee struct {
 FirstName, LastName string
}

func (e Employee) fullName() string {
 return e.FirstName + " " + e.LastName
}

func main() {
 e := Employee{
  FirstName: "Ross",
  LastName:  "Geller",
 }
 fmt.Println(e.fullName())
}
```

**Output**

```
Ross Geller
```

[Ejemplo](https://go.dev/play/p/68JWEHO9Yep)

En el programa anterior, hemos definido el método `fullName` que no toma ningún argumento pero devuelve una cadena. Como podemos ver en la declaración del receiver, este método pertenece al tipo `Employee`.

El método `fullName` pertenecerá a cualquier **objeto** del tipo `Employee`. Por lo tanto, ese **objeto** obtendrá automáticamente este método como una propiedad. Cuando se llama a este método en el objeto, recibirá el **objeto** como el receiver `e`.

Se puede acceder al **receiver** del método dentro del cuerpo del método. Por lo tanto, podemos acceder a `e` dentro del cuerpo del método de `fullName`. En el ejemplo anterior, dado que el receiver es una estructura de tipo `Employee`, podemos acceder a cualquier campo de la estructura. Como hicimos en el ejemplo anterior, estamos concatenando los campos `FirstName` y `LastName` y devolviendo el resultado.

Como un método pertenece a un tipo de receiver y está disponible en ese tipo como una propiedad, podemos llamar a ese método usando la sintaxis `Type.methodName(...)`. En el programa anterior, hemos usado `e.fullName()` para obtener el nombre completo de un empleado ya que el método `fullName` pertenece a `Employee`.

> Esto no es diferente de lo que vimos en el capitulo de [`structs`](../structs/structs.md) donde la función `fullName` era un campo de `struct`. Pero en el caso de los métodos, no tenemos que proporcionar propiedades de `struct` porque el método ya las conoce.

## 1.2 Métodos con el mismo nombre

Una diferencia importante entre funciones y métodos es que podemos tener varios métodos con el mismo nombre, mientras que no se pueden definir dos funciones con el mismo nombre en un mismo paquete.

Se nos permite crear métodos con el mismo nombre siempre que sus receptores sean diferentes. Vamos a crear dos tipos de estructura `Circle` y `Rectangle` y crear dos métodos del mismo nombre `Area` que calcula el área de su receiver.

**Code**

```go
package main

import (
 "fmt"
 "math"
)

type Rect struct {
 width  float64
 height float64
}

type Circle struct {
 radius float64
}

func (r Rect) Area() float64 {
 return r.width * r.height
}

func (c Circle) Area() float64 {
 return math.Pi * c.radius * c.radius
}

func main() {
 rect := Rect{5.0, 4.0}
 cir := Circle{5.0}
 fmt.Printf("Area of rectangle rect = %0.2f\n", rect.Area())
 fmt.Printf("Area of circle cir = %0.2f\n", cir.Area())
}
```

**Output**

```
Area of rectangle rect = 20.00
Area of circle cir = 78.54
```

[Ejemplo](https://go.dev/play/p/3f_MScaNXUw)

En el programa anterior, hemos creado los tipos de estructura `Rect` y `Circle` y hemos creado dos métodos del mismo nombre `Area` con el tipo de receptor `Rect` y `Circle`. Cuando llamamos al método `Area()` en `Rect` y `Circle`, se ejecutan sus respectivos métodos.

## 1.3 Pointer receivers

Hasta ahora, hemos visto que los métodos pertenecen a un tipo. Pero un método también puede pertenecer al puntero de un tipo.

Cuando un método pertenece a un tipo, su receptor recibe una copia del objeto sobre el que fue llamado. Para verificar eso, podemos crear un método que muta una estructura que recibe. Vamos a crear un método `changeName` que cambie el campo de nombre de una estructura de `Employee`.

**Code**

```go
package main

import "fmt"

type Employee struct {
 name   string
 salary int
}

func (e Employee) changeName(newName string) {
 e.name = newName
}

func main() {
 e := Employee{
  name:   "Ross Geller",
  salary: 1200,
 }

 // e before name change
 fmt.Println("e before name change =", e)

 // change name
 e.changeName("Monica Geller")

 // e after name change
 fmt.Println("e after name change =", e)
}
```

**Output**

```
e before name change = {Ross Geller 1200}
e after name change = {Ross Geller 1200}
```

[Ejemplo](https://go.dev/play/p/cV-W27SGcnl)

En el programa anterior, hemos llamado al método `changeName` en la estructura de tipo `Employee`. En el método, estamos mutando el valor del campo `name` de esa estructura.

A partir del resultado anterior, podemos verificar que aunque hayamos mutado el objeto receiver, no afectó al objeto original en el que se llamó al método.

Esto prueba que el método changeName recibió solo una copia de la estructura real (del método principal). Por lo tanto, cualquier cambio realizado en la copia dentro del método no afectó a la estructura original.

Pero un método también puede pertenecer al puntero de un tipo. La sintaxis para la definición del método que pertenece al puntero de un tipo es la siguiente.

```go
func (r *Type) functionName(...Type) Type {
    ...
}
```

Como puede ver en la definición anterior, la sintaxis para definir un método con un receiver de puntero es muy similar al método normal. En la siguiente definición, le indicamos a Go que este método pertenecerá al puntero `Type` en lugar del valor de `Type`.

Cuando un método pertenece al puntero de un tipo, su receiver recibirá el puntero al objeto en lugar de una copia del objeto. Reescribamos el ejemplo anterior con un método que recibe un receptor de puntero.

**Code**

```go
package main

import "fmt"

type Employee struct {
 name   string
 salary int
}

func (e *Employee) changeName(newName string) {
 (*e).name = newName
}

func main() {
 e := Employee{
  name:   "Ross Geller",
  salary: 1200,
 }

 // e before name change
 fmt.Println("e before name change =", e)
 // create pointer to `e`
 ep := &e
 // change name
 ep.changeName("Monica Geller")
 // e after name change
 fmt.Println("e after name change =", e)
}
```

**Output**

```
e before name change = {Ross Geller 1200}
e after name change = {Monica Geller 1200}
```

[Ejemplo](https://go.dev/play/p/PfSQ_GP-GPN)

Veamos qué cambios hicimos.

- Cambiamos la definición del método para recibir un receptor de puntero utilizando la sintaxis `*Employee`. De esta manera, el receptor `e` es el puntero al objeto de estructura al que se llamó este método.

- Dentro del cuerpo del método, estamos convirtiendo el puntero del receptor al valor del receptor usando la sintaxis de desreferenciación de puntero `(*p)`. Por lo tanto `(*e)` será el valor real de la estructura almacenada en la memoria.

- Entonces cambiamos el valor del campo `name` de la estructura `e`. Cualquier cambio realizado en `e` se reflejará en la estructura original.

- En el método principal, creamos un puntero `ep` que apunta a la estructura `e`.

- Dado que el método `changeName` pertenece al puntero de tipo `Employee` o tipo `*Empleado`, se puede invocar en el valor de tipo `*Empleado`.

- Dado que el tipo de `ep` es `*Employee`, podemos llamar al método `changeName` usando la sintaxis `ep.changeName()`. Esto pasará el puntero `ep` al método como receptor (en lugar del valor `e`).

> En el programa anterior, solo creamos el puntero `ep` de `e` solo para llamar al método `changeName` en él, pero también puede usar la sintaxis `(&e).changeName("Monica Geller")` en lugar de crear un nuevo puntero.

### 1.3.1 Métodos de llamada con receiver de puntero en valores

Quizá se preguntara, ¿siempre necesito crear un puntero para trabajar con métodos con receptor de puntero? Pero Go ya supuso que se haría esta pregunta.

Reescribamos el ejemplo anterior usando los atajos de Go.

**Code**

```go
package main

import "fmt"

type Employee struct {
 name   string
 salary int
}

func (e *Employee) changeName(newName string) {
 e.name = newName
}

func main() {
 e := Employee{
  name:   "Ross Geller",
  salary: 1200,
 }

 // e before name change
 fmt.Println("e before name change =", e)
 // change name
 e.changeName("Monica Geller")
 // e after name change
 fmt.Println("e after name change =", e)
}
```

**Output**

```
e before name change = {Ross Geller 1200}
e after name change = {Monica Geller 1200}
```

[Ejemplo](https://go.dev/play/p/D2zm7lpSme4)

El programa anterior funcionará bien como antes. Entonces, qué cambió.

- Si un método tiene un **receiver de puntero**, entonces no necesariamente necesita usar la sintaxis de desreferenciación de puntero `(*e)` para obtener el valor del receptor. Puede usar `e` simplemente, que será la dirección del valor al que apunta el puntero, pero Go entenderá que está tratando de realizar una operación en el valor mismo y, under the hood, convertirá `e` en `(*e)`.

- Además, no necesariamente necesita llamar a un método desde un puntero si el método tiene un receptor de puntero. En su lugar, se le permite llamar a este método en el valor y Go pasará el puntero del valor como receiver automáticamente.

> Puede decidir entre el método con receptor de puntero o receiver de valor según su caso de uso. Pero, en general, incluso si no desea mutar el receiver, se prefieren los métodos con receiver de puntero ya que no se crea nueva memoria para las operaciones (en el caso de métodos con receiver de valor).

## 1.4 Métodos en estructuras anidadas

Aprendimos mucho sobre la estructura anidada en el capitulo de [structs](../structs/structs.md#estructura-anidada). Como un campo de estructura también puede ser una estructura, podemos definir un método en la estructura principal y acceder a la estructura anidada para hacer lo que queramos.

**Code**

```go
package main

import "fmt"

type Contact struct {
 phone, address string
}

type Employee struct {
 name    string
 salary  int
 contact Contact
}

func (e *Employee) changePhone(newPhone string) {
 e.contact.phone = newPhone
}

func main() {
 e := Employee{
  name:    "Ross Geller",
  salary:  1200,
  contact: Contact{"011 8080 8080", "New Delhi, India"},
 }
 // e before phone change
 fmt.Println("e before phone change =", e)
 // change phone
 e.changePhone("011 1010 1222")
 // e after phone change
 fmt.Println("e after phone change =", e)
}
```

**Outline**

```
e before phone change = {Ross Geller 1200 {011 8080 8080 New Delhi, India}}
e after phone change = {Ross Geller 1200 {011 1010 1222 New Delhi, India}}
```

[Ejemplo](https://go.dev/play/p/8_QhLqqb9Ot)

En el ejemplo anterior, hemos definido el método `changePhone` en `*Employee` que recibe el puntero de `e`. Dentro de este método, podemos acceder a las propiedades de `e` que también contiene la estructura anidada de tipo `Contact`.

Dado que `e` es el puntero en el método, podemos mutar la estructura anidada. En el ejemplo anterior, hemos cambiado la estructura anidada `contact` mutando el valor del campo del `phone`.

### 1.4.1 Métodos en estructuras anidadas

Una estructura anidada también puede tener métodos. Si la estructura interna implementa un método, puede llamar a un método usando `.` (dot) accessor.

**Code**

```go
package main

import "fmt"

type Contact struct {
 phone, address string
}

type Employee struct {
 name    string
 salary  int
 contact Contact
}

func (c *Contact) changePhone(newPhone string) {
 c.phone = newPhone
}

func main() {
 e := Employee{
  name:   "Ross Geller",
  salary: 1200,
  contact: Contact{
   phone:   "011 8080 8080",
   address: "New Delhi, India",
  },
 }
 // e before phone change
 fmt.Println("e before phone change =", e)
 // change phone
 e.contact.changePhone("011 1010 1222")
 // e after phone change
 fmt.Println("e after phone change =", e)
}
```

**Output**

```
e before phone change = {Ross Geller 1200 {011 8080 8080 New Delhi, India}}
e after phone change = {Ross Geller 1200 {011 1010 1222 New Delhi, India}}
```

[Ejemplo](https://go.dev/play/p/DBikGzozbAy)

### 1.4.2 Estructuras anidadas anónimamente

En el capitulo sobre [structs](../structs/structs.md#campos-promocionados), también aprendimos sobre campos anónimos y promociones de campo. En pocas palabras, si un campo de una estructura es una estructura anónima, los campos de la estructura anidada se promoverán al padre.

Veamos cómo podemos usar los campos promocionados dentro de un método.

**Code**

```go
package main

import "fmt"

type Contact struct {
 phone, address string
}
type Employee struct {
 name   string
 salary int
 Contact
}

func (e *Employee) changePhone(newPhone string) {
 e.phone = newPhone
}

func main() {
 e := Employee{
  name:   "Ross Geller",
  salary: 1200,
  Contact: Contact{
   phone:   "011 8080 8080",
   address: "New Delhi, India",
  },
 }
 // e before phone change
 fmt.Println("e before phone change =", e)
 // change phone
 e.changePhone("011 1010 1222")
 // e after phone change
 fmt.Println("e after phone change =", e)
}
```

**Output**

```
e before phone change = {Ross Geller 1200 {011 8080 8080 New Delhi, India}}
e after phone change = {Ross Geller 1200 {011 1010 1222 New Delhi, India}}
```

[Ejemplo](https://go.dev/play/p/d7D8FYYmiof)

Como podemos ver en el ejemplo anterior, dado que la estructura `contact` está anidada de forma anónima dentro de la estructura `Employee`, sus campos se promoverán a `Employee` y podremos acceder a él en el objeto `e`.

Por lo tanto, cualquier método que acepte un receiver de estructura también tendrá acceso a los campos promocionados. Utilizando este principio, pudimos acceder a la propiedad `phone` del campo anidado `Contact` en el objeto `e` de tipo `Employee`.

### 1.4.3 Métodos promocionados

Al igual que los campos promocionados, los métodos implementados por la estructura anidada anónima también se promocionan a la estructura principal. Como vimos en el ejemplo anterior, el campo `Contact` está anidado de forma anónima. Por lo tanto, podríamos acceder al campo `phone` de la estructura interna en el padre.

En el mismo escenario, cualquier método implementado por la estructura `contact` estará disponible en la estructura `Employee`. Reescribamos el ejemplo anterior.

**Code**

```go
package main

import "fmt"

type Contact struct {
 phone, address string
}

type Employee struct {
 name   string
 salary int
 Contact
}

func (c *Contact) changePhone(newPhone string) {
 c.phone = newPhone
}

func main() {
 e := Employee{
  name:   "Ross Geller",
  salary: 1200,
  Contact: Contact{
   phone:   "011 8080 8080",
   address: "New Delhi, India",
  },
 }
 // e before phone change
 fmt.Println("e before phone change =", e)
 // change phone
 e.changePhone("011 1010 1222")
 // e after phone change
 fmt.Println("e after phone change =", e)
}
```

**Output**

```
e before phone change = {Ross Geller 1200 {011 8080 8080 New Delhi, India}}
e after phone change = {Ross Geller 1200 {011 1010 1222 New Delhi, India}}
```

[Ejemplo](https://go.dev/play/p/x2L-rR-TKNF)

Hicimos solo un cambio en el método `changePhone`. En lugar de recibir el tipo `*Empleado`, este método ahora espera un receptor del tipo `*Contact`. Dado que se promocionan los campos de la estructura anidada `Contact`, también se promocionará cualquier método implementado por ella. Por lo tanto, podríamos llamar a `e.changePhone()` como si el tipo `Employee` de struct e implementara el método`changePhone`.

> Sin embargo, una cosa para recordar aquí es que incluso estamos llamando al método `changePhone()` en `e`, el receptor enviado por Go será del tipo *Contact ya que este método le pertenece.

## 1.5 Métodos pueden aceptar ambos, punteros y valores

Cuando una función normal tiene una definición de parámetro, solo aceptará el argumento del tipo definido por el parámetro. Si pasó un puntero a la función que espera un valor, no funcionará. Esto también es cierto cuando la función acepta el puntero pero en su lugar está pasando un valor.

> Debe ver esto desde la perspectiva del tipo de datos. Una función que acepta un valor de tipo Type tiene una definición de parámetro `func (arg Type)`, mientras que una función que acepta un puntero tiene una definición de `func (arg *Type)`.

Pero cuando se trata de métodos, esa no es una regla estricta. Podemos definir un método con valor o receiver de puntero y llamarlo como puntero o valor. **Go hace el trabajo de conversión de tipo** under the hood, como hemos visto en los ejemplos anteriores.

**Code**

```go
package main

import "fmt"

type Employee struct {
 name   string
 salary int
}

func (e *Employee) changeName(newName string) {
 e.name = newName
}

func (e Employee) showSalary() {
 e.salary = 1500
 fmt.Println("Salary of e =", e.salary)
}

func main() {
 e := Employee{
  name:   "Ross Geller",
  salary: 1200,
 }
 // e before change
 fmt.Println("e before change =", e)
 // calling `changeName` pointer method on value
 e.changeName("Monica Geller")
 // calling `showSalary` value method on pointer
 (&e).showSalary()
 // e after change
 fmt.Println("e after change =", e)
}
```

**Output**

```
e before change = {Ross Geller 1200}
Salary of e = 1500
e after change = {Monica Geller 1200}
```

[Ejemplo](https://go.dev/play/p/tiQQbimhQ8_O)

En el programa anterior, definimos el método `changeName` que recibe un puntero pero llamamos al valor `e` que es legal porque Go under the hood le pasará un puntero de `e` (de tipo `*Employee`).

Además, definimos el método `showSalary` que recibe valor, pero lo llamamos en el puntero a `e`, lo cual es legal porque Go under the hood le pasará el valor del puntero (de tipo `*Employee`).

> Intentamos cambiar el salario de `e` dentro del método `showSalary` pero no funcionó como podemos ver en el resultado. Esto se debe a que incluso llamamos a este método en un puntero, Go enviará solo una copia del valor a ese método.

## 1.6 Métodos en no struct type

Hasta ahora hemos visto métodos pertenecientes al tipo struct pero por la definición de los métodos, es una función que puede pertenecer a cualquier tipo. Por lo tanto, **un método puede recibir cualquier tipo siempre que la definición de tipo y la definición de método estén en el mismo paquete**.

Hasta ahora, definimos la estructura y el método en el mismo paquete `main`, por lo que funcionó. Pero para verificar si podemos agregar métodos en tipos externos, intentaremos agregar un método a `toUpperCase` en el `string` de tipo incorporada.

**Code**

```go
package main

import (
 "fmt"
 "strings"
)

func (s string) toUpperCase() string {
 return strings.ToUpper(s)
}

func main() {
 str := "Hello World"
 fmt.Println(str.toUpperCase())
}
```

**Output**

```
program.go:8: cannot define new methods on non-local type string
program.go:14: str.toUpperCase undefined (type string has no field or method toUpperCase)
```

[Ejemplo](https://go.dev/play/p/FVvIh-YQBLH)

A partir del programa anterior, creamos el método `toUpperCase` que acepta strings como tipo receiver. Por lo tanto, esperamos que `string.toUpperCase()` funcione y devuelva la versión en mayúsculas del receiver `s`.

> Usamos el paquete integrado `strings` para convertir una cadena a mayúsculas.

Pero el programa anterior se ejecutará con un error de compilación.

Esto se debe a que el tipo `string` y el método `toUpperCase` no están definidos en el mismo paquete. Vamos a crear un nuevo tipo derivado `MyString` from `string`. De esta forma, tanto el método como el tipo `MyString` recién definido pertenecerán al mismo paquete y deberían funcionar.

**Code**

```go
package main

import (
 "fmt"
 "strings"
)

type MyString string

func (s MyString) toUpperCase() string {
 normalString := string(s)
 return strings.ToUpper(normalString)
}

func main() {
 str := MyString("Hello World")
 fmt.Println(str.toUpperCase())
}
```

**Output**

```
HELLO WORLD
```

[Ejemplo](https://go.dev/play/p/N-lLB3xPSDM)

A partir del programa anterior, creamos el método `toUpperCase` que ahora pertenece al tipo `MyString`. Necesitábamos modificar las partes internas de este método para pasar el tipo `string` a la función`strings.ToUpper`, pero lo conseguimos.

Ahora podemos llamar a `str.toUpperCase()` porque str es de tipo `MyString` ya que usamos conversión de tipo en la línea **no. 16** para convertir del tipo de cadena al tipo `MyString`.

# Referencias

[Métodos en Go](https://medium.com/rungo/anatomy-of-methods-in-go-f552aaa8ac4a)
