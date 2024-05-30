- [Que es una interfaz?](#1-que-es-una-interfaz)
  - [Declaración de interfaz](#12-declaraci%C3%B3n-de-interfaz)
  - [Implementación de interfaz](#13-implementaci%C3%B3n-de-interfaz)
  - [Interfaz vacía](#14-interfaz-vac%C3%ADa)
  - [Interfaces multiples](#15-interfaces-multiples)
  - [Type assertion](#16-type-assertion)
  - [Type switch](#17-type-switch)
  - [Interfaces Embebidas](#18-interfaces-embebidas)
  - [Pointer vs Value receiver](#19-pointer-vs-value-receiver)
  - [Comparación de interfaces](#110-comparaci%C3%B3n-de-interfaces)
  - [Uso de interfaces](#111-uso-de-interfaces)
- [Referencias](#2-referencias)

# 1. Que es una interfaz?

Las interfaces en Go no imponen un tipo para implementar métodos, pero las interfaces son herramientas muy poderosas. Un tipo puede optar por implementar métodos de una interfaz. Usando interfaces, un valor se puede representar en múltiples tipos, también conocido como polimorfismo.

Hablamos mucho sobre el objeto y el comportamiento en las lecciones de [`structs`](../structs/structs.md) y [`methods`](../methods/methods.md). También vimos cómo una estructura (y otros tipos) pueden implementar métodos. Una interfaz es otra pieza de un rompecabezas que acerca a Go al paradigma de la programación orientada a objetos.

Una interfaz es una colección de **firmas de métodos** que un Tipo puede implementar (usando métodos). Por lo tanto, la interfaz define (*no declara*) el comportamiento del objeto (*del tipo `Type`*).

Por ejemplo, un `Dog` puede `walk` y `bark`. Si una interfaz define firmas de métodos para `walk` y `bark` mientras `Dog` implementa métodos para `walk` y `bark`, se dice que `Dog` **implementa esa interfaz**.

El trabajo principal de una interfaz es proporcionar solo firmas de métodos que consisten en el **nombre del método, los argumentos de entrada y los return types**. Depende de un tipo (por ejemplo, tipo de estructura) declarar métodos e implementarlos.

Si es un programador de programación orientada a objetos, es posible que haya utilizado mucho la palabra clave implement al implementar una interfaz. Pero **en Go, no se menciona explícitamente si un tipo implementa una interfaz**.

Si un tipo implementa un método con nombre y firma definidos en una interfaz, entonces ese tipo implementa esa interfaz. Es como decir **"si camina como un pato, nada como un pato y grazna como un pato, entonces es un pato"**.

## 1.2 Declaración de interfaz

Al igual que `struct`, necesitamos crear un tipo derivado para simplificar la declaración de la interfaz usando la palabra clave interface.

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

> La convención de nomenclatura de las interfaces en Go es un poco complicada, puede seguir [este artículo de Medium para obtener más información](https://medium.com/@dotronglong/interface-naming-convention-in-golang-f53d9f471593). Si desea usar el prefijo I o el sufijo de interfaz, también está bien. Solo ténganlo en cuenta.

En el ejemplo anterior, hemos definido la interfaz `Shape` que tiene dos métodos `Area` y `Perimeter` que no aceptan argumentos y devuelve el valor `float64`. Cualquier tipo que implemente estos métodos (con firmas de métodos exactas) también implementará la interfaz `Shape`.

Dado que la interfaz es un tipo como una estructura, podemos crear una variable de su tipo. En el caso anterior, podemos crear una variable `s` de tipo interfaz `Shape`.

> Al igual que en un tipo de estructura, no hay absolutamente ninguna necesidad de crear un tipo derivado para una interfaz.

**Code**

```go
package main

import "fmt"

type Shape interface {
 Area() float64
 Perimeter() float64
}

func main() {
 var s Shape
 fmt.Println("value of s is", s)
 fmt.Printf("type of s is %T\n", s)
}
```

**Output**

```
value of s is <nil>
type of s is <nil>
```

[Ejemplo en vivo](https://go.dev/play/p/oGRDKbrEJYb)

Están sucediendo muchas cosas en este ejemplo, así que permitirme explicar algunos conceptos sobre las interfaces. Una interfaz tiene dos tipos. El tipo estático de una interfaz es la propia interfaz, por ejemplo, `Shape` en el programa anterior. Una interfaz no tiene un **valor estático**, sino que apunta a un **valor dinámico**.

Una variable de un tipo de interfaz puede contener un valor de un tipo que implementa la interfaz. El valor de ese tipo se convierte en el valor dinámico de la interfaz y ese tipo se convierte en el tipo dinámico de la interfaz.

Del ejemplo anterior, podemos ver que el valor cero y el tipo de interfaz es nulo. Esto se debe a que, en este momento, hemos declarado la variable s de tipo Shape pero no le asignamos ningún valor.

Cuando usamos la función `Println` del paquete `fmt` con el argumento de interfaz, apunta al **valor dinámico** de la interfaz y la sintaxis `%T` en la función `Printf` se refiere al **tipo dinámico** de interfaz.

> Pero en realidad, el tipo de interfaz es `Shape`, que es su tipo estático.

## 1.3 Implementación de interfaz

Declaremos los métodos `Area` y `Perimeter` con firmas proporcionadas por la interfaz `Shape`. Además, creemos la estructura `Shape` y hagamos que implemente la interfaz `Shape` (*implementando estos métodos*).

**Code**

```go
package main

import "fmt"

type Shape interface {
 Area() float64
 Perimeter() float64
}

type Rect struct {
 width  float64
 height float64
}

func (r Rect) Area() float64 {
 return r.width * r.height
}

func (r Rect) Perimeter() float64 {
 return 2 * (r.width + r.height)
}

func main() {
 var s Shape
 s = Rect{5.0, 4.0}
 r := Rect{5.0, 4.0}
 fmt.Printf("type of s is %T\n", s)
 fmt.Printf("value of s is %v\n", s)
 fmt.Println("area of rectange s", s.Area())
 fmt.Println("s == r is", s == r)
}
```

**Output**

```test
type of s is main.Rect
value of s is {5 4}
area of rectange s 20
s == r is true
```

[Ejemplo en vivo](https://play.golang.org/p/Hb__pA7Xp5V)

En el programa anterior, hemos creado la interfaz `Shape` y el tipo de estructura `Rect`. Luego definimos métodos como `Area` y `Perimeter` que pertenecen al tipo `Rect`, por lo tanto, `Rect` implementó esos métodos.

Dado que estos métodos están definidos por la interfaz `Shape`, *el tipo de estructura `Rect` implementa la interfaz `Shape`*. Dado que no hemos obligado a `Rect` a implementar la interfaz `Shape`, todo sucede automáticamente. Por lo tanto, podemos decir que **las interfaces en Go están implícitamente implementadas**.

Cuando un tipo implementa una interfaz, una variable de ese tipo también se puede representar como el tipo de una interfaz. Podemos confirmarlo creando una interfaz `nil` `s` de tipo `Shape` y asignando una estructura de tipo `Rect`.

> Acabamos de lograr el polimorfismo.

Dado que `Rect` implementa la interfaz `Shape`, esto es perfectamente válido. Del resultado anterior, podemos ver que el tipo dinámico de `s` ahora es `Rect` y el valor dinámico de `s` es el valor de la estructura `Rect` que es `{5 4}`.

> Lo llamamos dinámico porque podemos asignar `s` con una nueva estructura de un tipo de estructura diferente que también implementa la interfaz `Shape`.

A veces, el tipo dinámico de interfaz también se denomina **concrete type** (*tipo concreto*) porque cuando accedemos al tipo de interfaz, devuelve el tipo de su valor dinámico subyacente y su tipo estático permanece oculto.

Podemos llamar al método `Area` en `s` ya que la interfaz `Shape` define el método Area y el tipo concreto de `s` es `Rect` que implementa el método `Area`. Este método se llamará behind the scenes de la interfaz de valor dinámico.

Además, podemos ver que podemos comparar `s` con `r` ya que ambas variables tienen el mismo tipo dinámico (*estructura de tipo `Rect`*) y valor dinámico `{5 4}`.

> Puede aprender sobre la comparación de estructuras en la lección de [structs](../structs/structs.md).

Cambiemos el tipo dinámico y el valor dinámico de `s`.

**Code**

```go
package main

import (
 "fmt"
 "math"
)

type Shape interface {
 Area() float64
 Perimeter() float64
}

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

func (r Rect) Perimeter() float64 {
 return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
 return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
 return 2 * math.Pi * c.radius
}

func main() {
 var s Shape = Rect{10, 3}

 fmt.Printf("type of s is %T\n", s)
 fmt.Printf("value of s is %v\n", s)
 fmt.Printf("value of s is %0.2f\n\n", s.Area())

 s = Circle{10}
 fmt.Printf("type of s is %T\n", s)
 fmt.Printf("value of s is %v\n", s)
 fmt.Printf("value of s is %0.2f\n", s.Area())
}
```

**Output**

```
type of s is main.Rect
value of s is {10 3}
value of s is 30.00

type of s is main.Circle
value of s is {10}
value of s is 314.16
```

Si lees lecciones de [structs](../structs/structs.md) y [methods](../methods/methods.md), entonces el programa anterior no debería sorprenderte. Como el nuevo tipo de estructura `Circle` también implementa la interfaz `Shape`, podemos asignarle un valor de tipo de estructura `Circle`.

Supongo que ahora puedes relacionar por qué el tipo y el valor de la interfaz son dinámicos. De la lección de [slices](https://medium.com/rungo/the-anatomy-of-slices-in-go-6450e3bb2b94), aprendimos que una slice contiene la referencia a una matriz. De manera similar, podemos decir que **una interfaz también funciona de manera similar al mantener dinámicamente una referencia al underlying type (*tipo subyacente*)**.

¿Puedes adivinar qué pasará con el siguiente programa?

**Code**

```go
package main

import "fmt"

type Shape interface {
 Area() float64
 Perimeter() float64
}

type Rect struct {
 width  float64
 height float64
}

func (r Rect) Area() float64 {
 return r.width * r.height
}

func main() {
 var s Shape = Rect{10, 3}
 fmt.Println(s)
}
```

**Output**

```
./prog.go:20:16: cannot use Rect{…} (value of type Rect) as type Shape in variable declaration:
 Rect does not implement Shape (missing Perimeter method)
```

[Ejemplo en vivo](https://go.dev/play/p/pwhIwfHFzF9)

En el programa anterior, eliminamos el método `Perimeter`. Este programa no compilará y el compilador arrojará un error.

Debería ser obvio a partir del error anterior que para implementar con éxito una interfaz, debe implementar todos los métodos declarados por la interfaz con firmas exactas.

## 1.4 Interfaz vacía

Cuando una interfaz tiene cero métodos, se denomina interfaz vacía. Esto está representado por la `interface{}`. Dado que la interfaz vacía no tiene métodos, **todos los tipos implementan esta interfaz implícitamente**.

¿Se ha preguntado cómo la función `Println` del paquete integrado `fmt` acepta los diferentes tipos de valor como argumentos? Esto es posible debido a una interfaz vacía. Veamos la firma de la función `Println`.

```go
func Println(a ...interface{}) (n int, err error)
```

Como puede ver, `Println` es una función variable que acepta argumentos de tipo `interfaz{}`. Entendamos esto con profundidad.

Vamos a crear una función `explain` que acepta un argumento de tipo **empty interface** y explica el **dynamic value & type** de la interfaz.

**Code**

```go
package main

import "fmt"

type MyString string

type Rect struct {
 width  float64
 height float64
}

func explain(i interface{}) {
 fmt.Printf("value given to explain function is of type '%T' with value %v\n", i, i)
}

func main() {
 ms := MyString("Hello World!")
 r := Rect{5.5, 4.5}
 explain(ms)
 explain(r)
}
```

**Output**

```
value given to explain function is of type 'main.MyString' with value Hello World!
value given to explain function is of type 'main.Rect' with value {5.5 4.5}
```

[Example](https://go.dev/play/p/NhvO6Qjw_zp)

En el programa anterior, hemos creado un tipo `string` personalizado `MyString` y un tipo de estructura `Rect`. Dado que la función `explain` acepta un argumento del tipo **empty interface**, podemos pasar una variable de tipo `MyString`, `Rect` u otros.

Dado que todos los tipos implementan una interfaz vacía `interface{}`, esto es perfectamente legal. De nuevo **polimorfismo** for the win!!. El parámetro  `i` de la función `explain` es un tipo de interfaz pero su valor dinámico apuntará a cualquier valor que le hayamos pasado a la función como argumento.

## 1.5 Interfaces multiples

Un tipo puede implementar múltiples interfaces. Vamos a ver un ejemplo.

**Code**

```go
package main

import "fmt"

type Shape interface {
 Area() float64
}

type Object interface {
 Volume() float64
}

type Cube struct {
 side float64
}

func (c Cube) Area() float64 {
 return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
 return c.side * c.side * c.side
}

func main() {
 c := Cube{3}
 var s Shape = c
 var o Object = c
 fmt.Println("volume of s of interface type Shape is", s.Area())
 fmt.Println("area of o of interface type Object is", o.Volume())
}
```

**Output**

```
volume of s of interface type Shape is 54
area of o of interface type Object is 27
```

[Ejemplo en vivo](https://go.dev/play/p/YgW3NBxp8Fh)

En el programa anterior, creamos la interfaz `Shape` con el método `Area` y la interfaz `Object` con el método `Volume`. Dado que el tipo de estructura `Cube` implementa ambos métodos, implementa ambas interfaces. Por lo tanto, podemos asignar un valor de tipo de estructura `Cube` a la variable de tipo `Shape` u `Object`.

Esperamos que `s` tenga un valor dinámico de `c` y o que también tenga un valor dinámico de `c`. Usamos el método de `Area` en `s` de tipo de interfaz `Shape` porque define el método de `Area` y el método de volumen en `o` de tipo interfaz `Object` porque define el método de `Volume`. Pero, ¿qué sucederá si usamos el método de `Volume` en `s` y el método de `Area` en `o`?

Hagamos estos cambios en el programa anterior para ver qué sucede.

```go
fmt.Println("area of s of interface type Shape is", s.Volume())
fmt.Println("volume of o of interface type Object is", o.Area())
```

Los cambios anteriores producen el siguiente resultado.

```
program.go:31: s.Volume undefined (type Shape has no field or method Volume)
program.go:32: o.Area undefined (type Object has no field or method Area)
```

Este programa no se compilará debido a que el tipo estático de `s` es `Shape` y el tipo estático de `o` es `Object`. Dado que `Shape` no define el método `Volume` y `Object` no define el método `Area`, obtenemos el error anterior.

Para que funcione, necesitamos extraer de alguna manera el valor dinámico de estas interfaces, que es una estructura de tipo `Cube` y `Cube` implementa estos métodos. Esto se puede hacer usando **aserción de tipo**.

## 1.6 Type assertion

Podemos averiguar el valor dinámico subyacente de una interfaz usando la sintaxis **i.(Type)** donde `i` es una variable de tipo interfaz y `Type` es un tipo que implementa la interfaz. Go verificará si el tipo dinámico de `i` es idéntico a `Type` y devolverá el valor dinámico si es posible.

Reescribamos el ejemplo anterior y extraigamos el valor dinámico de la interfaz.

**Code**

```go
package main

import "fmt"

type Shape interface {
 Area() float64
}

type Object interface {
 Volume() float64
}

type Cube struct {
 side float64
}

func (c Cube) Area() float64 {
 return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
 return c.side * c.side * c.side
}

func main() {
 var s Shape = Cube{3}
 c := s.(Cube)
 fmt.Println("area of c of type Cube is", c.Area())
 fmt.Println("volume of c of type Cube is", c.Volume())
}
```

**Output**

```text
area of c of type Cube is 54
volume of c of type Cube is 27
```

[Ejemplo en vivo](https://go.dev/play/p/0e1XTpjuXJ_e)

El programa anterior, la variable de interfaz `s` de tipo `Shape` tiene el valor dinámico de `struct type Cube`. Usando notación abreviada, hemos extraído ese valor con la sintaxis `s.(Cube)` en la variable `c`.

Ahora, podemos usar métodos de `Area` y `Volume` en `c` ya que `c` es una estructura de tipo `Cube` y `Cube` implementa estos métodos.

**¡Tener cuidado!** En la sintaxis de type assertion (*aserción de tipos*) `i.(Type)`, si no puedo obtener el valor dinámico de `Type` porque `Type` no implementa la interfaz, entonces el compilador de Go generará un error de compilación.

```text
impossible type assertion:
XYZ does not implement Shape (missing Area method)
```

Pero si `Type` implementa la interfaz pero no tiene un valor concreto de `Type` (*porque es nulo en este momento*), entonces Go lanzará un panic en el tiempo de ejecución.

```text
panic: interface conversion: main.Shape is nil, not main.Cube
```

Afortunadamente, para evitar el **panic en el tiempo de ejecución**, hay otra variante de sintaxis de aserción de tipo que fallará silenciosamente.

```text
value, ok := i.(Type)
```

En la sintaxis anterior, podemos verificar usando la variable `ok` si `i` tiene un tipo concreto `Type` o un valor dinámico de `Typo`. Si no es así, `ok` será `false` y el valor será el valor cero de `Type`.

¿Cómo sabríamos si el valor subyacente de una interfaz implementa otras interfaces? Esto también es posible usando type assertion (*aserción de tipo*). Si `Type` en la sintaxis de aserción de tipo es una interfaz, Go comprobará si el tipo dinámico de `i` implementa la interfaz `Type`.

**Code**

```go
package main

import "fmt"

type Shape interface {
 Area() float64
}

type Object interface {
 Volume() float64
}

type Skin interface {
 Color() float64
}

type Cube struct {
 side float64
}

func (c Cube) Area() float64 {
 return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
 return c.side * c.side * c.side
}

func main() {
 var s Shape = Cube{3}
 value1, ok1 := s.(Object)
 fmt.Printf("dynamic value of Shape 's' with value %v implements interface Object? %v\n", value1, ok1)
 value2, ok2 := s.(Skin)
 fmt.Printf("dynamic value of Shape 's' with value %v implements interface Skin? %v\n", value2, ok2)
}
```

**Output**

```text
dynamic value of Shape 's' with value {3} implements interface Object? true
dynamic value of Shape 's' with value <nil> implements interface Skin? false
```

[Ejemplo en vivo](https://go.dev/play/p/Iu84WAzDEwx)

Dado que el tipo dinámico de `s` es `Cube` y `Cube` implementa la interfaz `Object`, la primera aserción tiene éxito. El `value1` es una interfaz de tipo `Object` y también apunta al valor dinámico de `s` (impreso por la función `Printf`).

Pero dado que `Cube` struct no implementa la interfaz `Skin`, obtuvimos `ok2` como `false` y`value2` como `nil` (valor cero de la interfaz). Si hubiéramos usado la sintaxis más simple de la sintaxis `v := i.(type)`, entonces nuestro programa habría lanzado un panic error en tiempo de ejecución con el siguiente error.

```
panic: interface conversion: main.Cube is not main.Skin: missing method Color
```

> Tome nota, necesitamos usar la aserción de tipo para obtener el valor dinámico de una interfaz de modo que podamos acceder a las propiedades de ese valor dinámico. Como por ejemplo, no puede acceder a los campos de una estructura en el objeto de tipo interfaz, incluso si tiene un valor dinámico de estructura.
> En pocas palabras, acceder a cualquier cosa que no esté representada por el tipo de interfaz provocará un panic error en tiempo de ejecución. Así que asegúrese de usar la aserción de tipo cuando sea necesario.

La aserción de tipo no solo se usa para verificar si una interfaz tiene un valor concreto de algún tipo dado, sino también para **convertir una variable dada de un tipo de interfaz a un tipo de interfaz diferente** (consulte el ejemplo anterior o [este ejemplo](https://go.dev/play/p/GUpJKfGQC6D)).

**Code**

```go
package main

import "fmt"
import "reflect"

// person interface
type Person interface {
 getFullName() string
}

// salaried interface
type Salaried interface {
 getSalary() int
}

// Employee struct represents an employee in an organization
type Employee struct {
 firstName string
 lastName string
 salary int
}

// using this method, Employee implements Person interface
func (e Employee) getFullName() string {
 return e.firstName + " " + e.lastName
}

// using this method, Employee implements Salaried interface
func (e Employee) getSalary() int {
 return e.salary
}

func main() {
 var johnP Person = Employee{"John", "Adams", 2000}

 // show john's salary
 fmt.Printf("full name : %v \n", reflect.ValueOf(johnP).Interface())

 // convert john Person to Salaried type
 johnS := johnP.(Salaried)

 fmt.Printf("salary : %v \n", johnS.getSalary()
}
```

**Output**

```
full name : {John Adams 2000}
salary : 2000
```

## 1.7 Type switch

Hemos visto una **interfaz vacía** y su uso. Reconsideremos la función `explain` que vimos antes. Como el tipo de argumento de la función `explain` es una **interfaz vacía**, podemos pasarle cualquier argumento.

Pero si el argumento es un `string`, queremos que la función `explain` imprima el resultado en mayúsculas. ¿Cómo podemos hacer que eso suceda?

Podemos usar la función `ToUpper` del paquete `strings`, pero dado que solo acepta un argumento `string`, debemos asegurarnos desde dentro de la función `explain` de que el **tipo dinámico** de la interfaz vacía `i` es un `string` mientras lo hacemos.

Esto se puede hacer usando el **Type switch**. La sintaxis para el type switch es similar a type assertion y es `i.(type)` donde `i` es interfaz y tipo es una **fixed keyword**. Usando esto, podemos obtener el tipo dinámico de la interfaz en lugar del valor dinámico.

> Pero esta sintaxis solo funcionará en un `switch`.

**Code**

```go
package main

import (
 "fmt"
 "strings"
)

func explain(i interface{}) {
 switch i.(type) {
 case string:
  fmt.Println("i stored string ", strings.ToUpper(i.(string)))
 case int:
  fmt.Println("i stored int", i)
 default:
  fmt.Println("i stored something else", i)
 }
}

func main() {
 explain("Hello World")
 explain(52)
 explain(true)
}
```

**Output**

```text
i stored string  HELLO WORLD
i stored int 52
i stored something else true
```

[Ejemplo en vivo](https://go.dev/play/p/ItSSq3VDMbB)

En el programa anterior, modificamos la función `explain` para usar un **type switch**. Cuando se llama a una función de explicación con cualquier tipo, `i` recibe su **valor dinámico** y su **tipo dinámico**.

Usando la declaración `i.(type)` dentro del `switch`, estamos obteniendo acceso a ese **tipo dinámico**. Usando **casos** dentro del `switch`, podemos hacer operaciones condicionales basadas en el tipo dinámico de la interfaz `i`.

En el caso del `string`, usamos la función `strings.ToUpper` para convertir el `string` a mayúsculas. Pero dado que solo acepta el tipo de datos `string`, necesitábamos obtener el valor dinámico subyacente. Por lo tanto, usamos **type assertion**.

## 1.8 Interfaces Embebidas

En Go, una interfaz no puede implementar otras interfaces ni extenderlas, pero podemos crear una nueva interfaz fusionando dos o más interfaces. Reescribamos nuestro programa `Shape-Cube`.

**Code**

```go
package main

import "fmt"

type Shape interface {
 Area() float64
}

type Object interface {
 Volume() float64
}

type Material interface {
 Shape
 Object
}

type Cube struct {
 side float64
}

func (c Cube) Area() float64 {
 return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
 return c.side * c.side * c.side
}

func main() {
 c := Cube{3}
 var m Material = c
 var s Shape = c
 var o Object = c
 fmt.Printf("dynamic type and value of interface m of static type Material is'%T' and '%v'\n", m, m)
 fmt.Printf("dynamic type and value of interface s of static type Shape is'%T' and '%v'\n", s, s)
 fmt.Printf("dynamic type and value of interface o of static type Object is'%T' and '%v'\n", o, o)
}
```

**Output**

```text
dynamic type and value of interface m of static type Material is'main.Cube' and '{3}'
dynamic type and value of interface s of static type Shape is'main.Cube' and '{3}'
dynamic type and value of interface o of static type Object is'main.Cube' and '{3}'
```

[Example](https://go.dev/play/p/s2U79IDaKqF)

En el programa anterior, dado que `Cube` implementa el método `Area` y `Volume`, implementa las interfaces `Shape` y `Object`. Pero dado que la interfaz `Material` es una interfaz embebida de estas interfaces, `Cube` también debe implementarla.

Esto es posible porque, al igual que la **estructura anidada de forma anónima**, todos los métodos de las interfaces anidadas se promocionan a interfaces principales.

## 1.9 `Pointer` vs `Value` receiver

Hasta ahora, en este tutorial, hemos visto methods with value receivers (*métodos con receptores de valor*). ¿La interfaz estará bien con el método que acepta el receptor del puntero? Vamos a ver.

**Code**

```go
package main

import "fmt"

type Shape interface {
 Area() float64
 Perimeter() float64
}

type Rect struct {
 width  float64
 height float64
}

func (r *Rect) Area() float64 {
 return r.width * r.height
}

func (r Rect) Perimeter() float64 {
 return 2 * (r.width + r.height)
}

func main() {
 r := Rect{5.0, 4.0}
 var s Shape = r
 area := s.Area()
 fmt.Println("area of rectangle is", area)
}
```

En el programa anterior, el método `Area` pertenece al tipo `*Rect`, por lo que su receptor obtendrá el puntero de la variable de tipo `Area`. Sin embargo, el programa anterior no se compilará y Go arrojará un error de compilación.

**Output**

```text
./prog.go:25:16: cannot use r (variable of type Rect) as type Shape in variable declaration:
 Rect does not implement Shape (Area method has pointer receiver)

Go build failed.
```

[Example](https://go.dev/play/p/vEkRuYo1JKu)

¿Qué demonios? Podemos ver claramente que el tipo de estructura `Rect` está implementando todos los métodos establecidos por la interfaz `Shape`, entonces, ¿por qué obtenemos que `Rect` no implementa el error `Shape`?

Si lee el error detenidamente, dice que el `Area method has pointer receiver`. Entonces, ¿qué pasa si el método `Area` tiene un receptor de puntero?

Bueno, hemos visto la lección de [`structs`](../structs/structs.md) que un método con un receptor de puntero funcionará tanto con el puntero como con el valor y si hubiéramos usado `r.Area()` en el programa anterior, se habría compilado bien.

Sin embargo, con las interfaces, esto es un poco diferente. El tipo dinámico de interfaz `s` es `Rect` y podemos ver claramente que `Rect` no implementa el método `Area` pero `*Rect` sí.

> Podrías preguntarte, ¿por qué la asignación `r := Rect{}` no falló porque claramente `Rect` no implementa la interfaz `Shape`?

> Esto se debe a que se puede llamar a un método con puntero o receptor de valor tanto en el valor como en el puntero, y la conversión de **valor a un puntero** o **puntero a un valor** adecuado para pasar como receptor para la llamada al método se realiza mediante Go under the hood ( como se ve en la lección de [métodos](../methods/methods.md) ).

> Por lo tanto, en el momento de la compilación, tanto el puntero como el valor se pueden almacenar en una variable de tipo interfaz. Sin embargo, al llamar a un método en la propia interfaz, el tipo dinámico se considera en tiempo de ejecución y el valor dinámico se pasa como receptor al método.

Para que este programa funcione, en lugar de asignar un valor del tipo `Rect` a la variable de interfaz `s`, necesitamos asignar un puntero (*de tipo `*Rect`*) para que el puntero se pase como receptor al método `Area`.

Reescribamos el programa anterior con este concepto.

**Code**

```go
package main

import "fmt"

type Shape interface {
 Area() float64
 Perimeter() float64
}

type Rect struct {
 width  float64
 height float64
}

func (r *Rect) Area() float64 {
 return r.width * r.height
}

func (r Rect) Perimeter() float64 {
 return 2 * (r.width + r.height)
}

func main() {
 r := Rect{5.0, 4.0}
 var s Shape = &r // assigned pointer
 area := s.Area()
 perimeter := s.Perimeter()
 fmt.Println("area of rectangle is", area)
 fmt.Println("perimeter of rectangle is", perimeter)
}
```

**Output**

```text
area of rectangle is 20
perimeter of rectangle is 18
```

[Enlace](https://go.dev/play/p/3OY4dBOSXdL)

El único cambio que hicimos está en la línea **no. 25** donde en lugar del valor de `r`, usamos el puntero a `r` para que el tipo dinámico de `s` se convierta en `*Rect` e implemente el método `Area`.

Sin embargo, la llamada a `s.Perimeter()` no falló a pesar de que Perimeter no está implementado por `*Area`.

Parece que Go está feliz de pasar una **copia del valor del puntero** como receptor al método `Perimeter` quizás porque no es una idea muy peligrosa, es solo una copia y nadie puede mutarla accidentalmente.

> Sin embargo, desearía que Go pudiera haber procesado la llamada al método en interfaces similares a la estructura para que no tengamos que preocuparnos por la conversión del puntero. Pero **podría deberse a la seguridad y la capacidad de la interfaz para almacenar datos**.

## 1.10 Comparación de interfaces

- Se pueden comparar dos interfaces con los operadores `==` y `!=`. Dos interfaces son siempre iguales si los valores dinámicos subyacentes son `nil`, lo que significa que dos `nil` interfaces son siempre iguales, por lo tanto, la operación `==` devuelve `true`.

```go
var a, b interface{}
fmt.Println( a == b ) // true
```

- Si estas interfaces no son `nil`, entonces sus tipos dinámicos (*el tipo de sus valores concretos*) deberían ser los mismos y los valores concretos deberían ser iguales.

- Si los tipos dinámicos de la interfaz **no son comparables**, como por ejemplo, `slice`, `map` and `function`, o el valor concreto de una interfaz es una estructura de datos compleja como `slice` or `array` que contiene estos valores incomparables, entonces `==` o `!=` operaciones dará como resultado un **runtime panic** (*panic en tiempo de ejecución*).

- Si una interfaz es nula, entonces la operación `==` siempre devolverá `false`.

## 1.11 Uso de interfaces

Hemos aprendido que son las interfaces y vimos que pueden tomar diferentes formas. Esa es la definición de polimorfismo.

Las interfaces son muy útiles en el caso de funciones y métodos en los que necesita argumentos de tipos dinámicos, como la función `Println`, que acepta cualquier tipo de valores.

Cuando varios tipos implementan la misma interfaz, resulta fácil trabajar con ellos. Por lo tanto, siempre que podamos usar interfaces, deberíamos hacerlo.

# 2. Referencias

[Documentación oficial de golang acerca de interfaces](https://go.dev/ref/spec#Interface_types)
[Interfaces en Go](https://medium.com/rungo/interfaces-in-go-ab1601159b3a)
