# Trabajando con JSON

## Referencias

[Working with JSON in Go](https://medium.com/rungo/working-with-json-in-go-7e3a37c5a07b)

## Introducción

En Javascript, podemos usar la función `JSON.parse(json_string)` para convertir un JSON string en un objeto Javascript.

Si queremos mandar un objeto javascript al servidor, nosotros podemos usar la función `JSON.stringify(js_object)` la cual nos devuelve un JSON string.

Como podemos ver, el formato JSON es similar a un Mapa de datos con claves y valores. Una clave JSON es estrictamente un string, mientras el valor puede ser cualquier tipo de dato que sea soportado por JSON.

JSON format soporta 6 tipos de datos, `string`, `number`, `boolean`, `null`, `array` y `object`.

En Go, lo más cerca que podemos llegar a la representación JSON a traves de los tipos de datos `map` o `struct`.

Ambos de estas estructuras de datos pueden almacenar datos complejos en pares de clave/valor.

Mientras **decodificamos** un `JSON string`, necesitamos un contenedor que pueda almacenar los datos JSON en un tipo de dato váĺido como un `map` o un `struct`. De igual forma, cuando **codificamos** un `JSON data`, necesitamos un objeto como un `mapa` o un `struct` que pueden ser convertidos a un formato de JSON válido.

Analicemos las API proporcionadas por el codificador y decodificador de JSON de Go. Go proporciona la mayoría de las API de codificación / decodificación en el paquete `encoding / json`.

## Codificando JSON

Para codificar un JSON a partir de una estructura de datos adecuada, utilizamos la función `json.Marshal` proporcionada por el paquete `json`. Esta función tiene la siguiente sintaxis.

```go
func Marshal(v interface{}) ([]byte, error)
```

Posemos usar un struct o un `map` como el argumento `v` para la función `Marshal` (para codificar JSON data). Esta función devuelve un slice de bytes el cual no es mas que los datos JSON codificacos en UTF-8 y un error si el objeto `v` no se puede codificar como una cadena JSON.

Let’s create a simple struct and encode JSON from it.

```go
package main

import(
 "fmt"
 "encoding/json"
)

// Student declares `Student` structure
type Student struct {
  FirstName, lastName string
  Email string
  Age int
  HeightInMeters float64
  IsMale bool
}

func main() {
  // define `john` struct
  john := Student{
    FirstName: "John",
    lastName: "Doe",
    Age: 21,
    HeightInMeters: 1.75,
    IsMale: true,
  }

  // encode `john` as JSON
  johnJSON, _ := json.Marshal( john )

  // print JSON string
  fmt.Println( string(johnJSON) )
}
```

```shell
{"FirstName":"John","Email":"","Age":21,"HeightInMeters":1.75,"IsMale":true}

Program exited.
```

> Si deseas formatear el JSON con nuevas líneas y sangría, puede usar la función `json.MarshalIndent` que llama internamente a la función Marshal. Esta función tomará el prefijo y la sangría como argumento.

Si echas un vistazo al resultado anterior, es posible que encuentres algo extraño. En primer lugar, la función Marshal considera el nombre del campo del struct como la clave del elemento JSON, lo que puede ser útil, pero la mayoría de las veces, necesitamos nombres de campo personalizados en los datos JSON. Los nombres de campo en mayúsculas parecen raros.

We can also encode map data type into JSON data. The necessary condition is that map keys should be either string or int. If a map key is an integer, it will be coerced to a string for encoding.

También podemos codificar map data type en JSON data. La condición necesaria es que las claves del mapa deben ser de cadena o int. Si una clave de mapa es un número entero, se convertirá en una cadena para la codificación.

```go
package main

import(
  "fmt"
  "encoding/json"
)

// Student declares `Student` map
type Student map[string]interface{}

func main() {

  // define `john` struct
  john := Student{
    "FirstName": "John",
    "lastName": "Doe",
    "Age": 21,
    "HeightInMeters": 1.75,
    "IsMale": true,
  }

  // encode `john` as JSON
  johnJSON, _ := json.Marshal( john )

  // print JSON string
  fmt.Println( string(johnJSON) )
}
```

```shell
{"Age":21,"FirstName":"John","HeightInMeters":1.75,"IsMale":true,"lastName":"Doe"}

Program exited.
```

Esta vez, no tuvimos que preocuparnos por eliminar ningún campo del JSON, ya que estamos usando un `map` para codificar los datos. Es posible que sientas que el mapa es el camino correcto a seguir para la codificación JSON, pero el struct aporta muchas características, simplemente no se pueden comparar.

### Manejo de tipos de datos

Como hemos aprendido, JSON admite principalmente 6 tipos de datos, `string`, `number`, `boolean`, `null`, `array` and `object`. Esta puede ser una gran noticia para un desarrollador de JavaScript porque todos los tipos de datos se suponen en JavaScript, pero en Go, debemos considerar varios tipos de datos durante la codificación.

1. **Number**: Un `int` o un `float` o un json. Un valor number es codificado como un  `JSON number value`.

2. **String**: Un valor `string` es saneado y codificado como `JSON string value`. El valor tipo `[]byte` es codificado como value un `Base64 string`.

3. **Boolean**: Un valor `bool` es codificado como un `JSON boolean value`.

4. **Null**: Un valor `nil` (como un de un puntero, interfaz u otro tipo de datos) es codificado como `JSON null value`.

5. **Object**: Un valor `map` o un `struct` es codificado como `JSON object value`.

6. **Array**: Un valor `array` o un `slice` es codificado como un `JSON array value` excepto por el slide de bytes (`[]byte`).

### Tipos de datos abstractos

En los ejemplos anteriores, hemos codificado valores de tipos de datos concretos como `int`, `string`, `bool`, etc. Agreguemos valores de datos más complejos como `struct`, `map` e `interface` a un objeto y veamos cómo se codifica en JSON.

```go
package main

import(
  "fmt"
  "encoding/json"
)

// Profile declares `Profile` structure
type Profile struct {
  Username string
  followers int
  Grades map[string]string
}

// Student declares `Student` structure
type Student struct {
  FirstName, lastName string
  Age int
  Profile Profile
  Languages []string
}

func main() {

  var john Student

  // define `john` struct
  john = Student{
    FirstName: "John",
    lastName: "Doe",
    Age: 21,
    Profile: Profile{
      Username: "johndoe91",
      followers: 1975,
      Grades: map[string]string{ "Math": "A", "Science": "A+" },
    },
    Languages: []string{ "English", "French" },
  }

  // encode `john` as JSON
  johnJSON, err := json.MarshalIndent( john, "", "  " )

  // print JSON string
  fmt.Println( string(johnJSON), err )
}
```

En el ejemplo anterior, hemos agregado el campo `Profile` y `Languages` a la estructura del `Student` struct, que tiene `Profile` y `map[string]` respectivamente. Este programa produce el siguiente resultado.

```shell
{
  "FirstName": "Mike",
  "Age": 21,
  "Profile": {
    "Username": "mikedoe91",
    "Grades": {
      "Math": "A",
      "Science": "A+"
    }
  },
  "Languages": [
    "English",
    "French"
  ]
} <nil>

Program exited.
```

Si observa este resultado con atención, no obtuvimos el campo de `followers` en el objeto `Profile` porque no se exporta desde el tipo `Profile`.

Como hemos aprendido de la lección "Structs en Go", una estructura puede tener una estructura anidada anónimamente. En tal caso, los campos de esa estructura (así como los métodos) se promueven a la estructura principal.

```go
package main

import(
  "fmt"
  "encoding/json"
)

// Profile declares `Profile` structure
type Profile struct {
  Username string
  followers int
  Grades map[string]string
}

// Student declares `Student` structure
type Student struct {
  FirstName, lastName string
  Age int
  Profile
  Languages []string
}

func main() {

  var john Student

  // define `john` struct
  john = Student{
    FirstName: "John",
    lastName: "Doe",
    Age: 21,
    Profile: Profile{
      Username: "johndoe91",
      followers: 1975,
  },
    Languages: []string{ "English", "French" },
  }

  // encode `john` as JSON
  johnJSON, _ := json.MarshalIndent( john, "", "  " )

  // print JSON string
  fmt.Println( string(johnJSON) )
}
```

En el ejemplo anterior, hemos hecho anónimo el campo `Profile` del tipo de estructura `Student`. Esto hará que todos los campos del `profile` sean promovidos al tipo de estructura de `Student` padre. Este cambio produce el siguiente resultado.

```shell
{
  "FirstName": "John",
  "Age": 21,
  "Username": "johndoe91",
  "Grades": null,
  "Languages": [
    "English",
    "French"
  ]
}

Program exited.
```

Observe que los campos `Username` y `Grades` ahora forman parte del objeto principal. Sin embargo, el campo `Grades` está codificado como nulo porque no se inicializó y dado que el valor cero de un mapa es nulo, se codificó como nulo según la regla.

 > Si un nombre de campo promocionado entra en conflicto con el nombre de campo de la estructura principal, se selecciona el campo menos anidado para la clasificación.

Si un valor es un puntero, entonces el valor del puntero se usa por el `marshaling`. Si un valor es una interfaz, entonces el valor concreto de la interfaz se utiliza por el `marshaling`.

```go
package main

import(
  "fmt"
  "encoding/json"
)

// ProfileI interface defines `Follow` method
type ProfileI interface {
  Follow()
}

// Profile declares `Profile` structure
type Profile struct {
  Username string
  Followers int
}

// Follow method implementation
func (p *Profile) Follow(){
  p.Followers++
}

// Student declares `Student` structure
type Student struct {
  FirstName, lastName string
  Age int
  Primary ProfileI
  Secondary ProfileI
}

func main() {

  // define `john` struct (pointer)
  john := &Student{
    FirstName: "John",
    lastName: "Doe",
    Age: 21,
    Primary: &Profile{
      Username: "johndoe91",
      Followers: 1975,
    },
  }

  // follow `john`
  john.Primary.Follow()

  // encode `john` as JSON
  johnJSON, _ := json.MarshalIndent( john, "", "  " )

  // print JSON string
  fmt.Println( string(johnJSON) )
}
```

En el ejemplo anterior, hemos creado el tipo de interfaz `ProfileI` que declara el método `Follow`. Hemos implementado este método en el tipo de estructura de `Profile` con el receptor `* Profile`, lo que significa que `*Profile` ahora implementa la interfaz `ProfileI`.

En la estructura `Student`, hemos agregado el campo `Primary` y `Secondary` del tipo de interfaz `ProfileI`. Eso significa que cualquier valor que implemente la interfaz ``ProfileI` puede asignarse a estos campos.

En la función principal, hemos declarado la estructura `john` de tipo `*Student` asignando un puntero del tipo `Student`. No hemos asignado el campo `Secondary` de `john`, pero el campo `Primary` contiene un puntero a una estructura de `Profile`. Esto es legal porque el campo `Primary` es del tipo de interfaz `ProfileI` y el tipo `*Profile` implementa la interfaz `ProfileI`.

Más tarde, llamamos al método `john.Primary.Follow()` para incrementar el recuento de seguidores. Si observa con atención, hemos pasado un puntero a la función `MarshalIndent`, lo cual está bien porque Marshal usa el valor del puntero *under the hood* para calcular las referencias. Este programa produce el siguiente resultado.

```shell
{
  "FirstName": "John",
  "Age": 21,
  "Primary": {
    "Username": "johndoe91",
    "Followers": 1976
  },
  "Secondary": null
}

Program exited.
```

El campo `Secondary` es nulo porque el valor del campo `Secondary` en la estructura `john` es nulo (el valor cero de una interfaz es nulo). El campo `Primary` es un objeto JSON porque el campo `Primary` es una interfaz y tiene un puntero a la estructura del perfil como valor concreto.

### Conversión de tipos de datos

A veces, no queremos codificar un valor de un campo tal como está, sino proporcionar un valor personalizado para el `marshaling`. Esto se puede lograr implementando la interfaz `json.Marshaler` o `encoding.TextMarshaler`.

```go
// from `encoding/json` package
type Marshaler interface {
    MarshalJSON() ([]byte, error)
}
// from `encoding` package
type TextMarshaler interface {
    MarshalText() (text []byte, err error)
}
```

Si el valor de un campo implementa una de las interfaces anteriores, entonces la función `Marshal` no considerará el valor del campo para el cálculo y, en su lugar, utilizará el valor devuelto por el método `MarshalJSON` o el método `MarshalText`.

```go
package main

import(
  "fmt"
  "encoding/json"
)

// Profile declares `Profile` structure
type Profile struct {
  Username string
  Followers int
}

// MarshalJSON - implement `Marshaler` interface
func (p Profile) MarshalJSON() ([]byte, error) {
  // return JSON value
  // TODO: handle error gracefully
  return []byte(fmt.Sprintf(`{"f_count": "%d"}`, p.Followers)), nil;
}

// Age declares `Age` type
type Age int

// MarshalText - implement `TextMarshaler` interface
func (a Age) MarshalText() ([]byte, error){
  // return string value
  // TODO: handle error gracefully
  return []byte(fmt.Sprintf(`{"age": %d}`, int(a))), nil;
}

// Student declares `Student` structure
type Student struct {
  FirstName, lastName string
  Age Age
  Profile Profile
}

func main() {

  // define `john` struct (pointer)
  john := &Student{
    FirstName: "John",
    lastName: "Doe",
    Age: 21,
    Profile: Profile{
      Username: "johndoe91",
      Followers: 1975,
    },
  }

  // encode `john` as JSON
  johnJSON, _ := json.MarshalIndent( john, "", "  " )

  // print JSON string
  fmt.Println( string(johnJSON) )
}
```

En el ejemplo anterior, el tipo de estructura `Profile` implementa el método `MarshalJSON`, por lo que implementa la interfaz `Marshaler`. Cuando Marshal encuentra un valor de tipo Marshaler, llamará a este método para obtener los datos JSON (porción de bytes) en lugar de calcular el valor.

También hemos modificado el campo `Age` del tipo de estructura `Student`. Ahora tiene el valor de tipo `Age`, que es un alias de tipo `int`. El tipo de estructura de `Age` implementa el método `MarshalText`, lo que significa que implementa la interfaz `TextMarshaler` definida en el paquete de codificación.

Cuando `Marshal` encuentra un valor de tipo `TextMarshaler`, llamará al método `MarshalText` y usará el valor devuelto por este método para codificar el valor de cadena JSON. Por lo tanto, este programa produce lo siguiente a continuación.

```shell
{
  "FirstName": "John",
  "Age": "{\"age\": 21}",
  "Profile": {
    "f_count": "1975"
  }
}

Program exited.
```

### Codificando usando Structure Tags

Un `struct` puede contener metadatos adicionales que otros programas pueden usar para procesar ese campo de manera diferente. Estos metadatos se asignan a un campo mediante un literal de `string` (`raw string`, cadena sin formato `[``]` o `interpreted string` cadena interpretada `[“”]`)

```go
type Data struct {
  FieldOne string `json:"fname" xml:"first-name" gorm:"size:255"`
  FieldTwo string `json:"lname" xml:"last-name" gorm:"size:255"`
}
```

Las `structure tags` pueden ser utilizadas por muchos programas como codificador/ decodificador para obtener información adicional sobre el campo o motores de validación como validador para obtener criterios de validación o motores ORM como GORM para datos de columna de una  tabla.

Las `structure tags` se ignoran en los casos de uso general, pero si desea leer la etiqueta de un campo de estructura, puede usar el paquete `reflect` incorporado.

In our case, Marshal function uses the tag of a struct field to obtain additional encoding/encoding information from the field. For JSON encoding, we need to use json:"options" tag value. Here, the options are comma-separated string values.

En nuestro caso, la función `Marshal` usa la etiqueta de un `structure tag` para obtener información adicional de codificación/codificación del campo. Para la codificación JSON, necesitamos usar `json:"opciones" tag value`. Aquí, las opciones son valores de cadena separados por comas.

El primer valor de la opción es el nombre del campo que debería aparecer en el JSON. Los otros valores de las opciones pueden ser `omitempty` para descartar un campo si su valor está vacío o es un `string` para convertir el valor del campo en un `string`.

If we want to ignore a field unconditionally, we can use - as the options value. However, - can also be a valid JSON element key, hence -, value will specify that we want the field name to be -.

Si queremos ignorar un campo incondicionalmente, podemos usarlo `-` como valor de opciones. Sin embargo, `-` también puede ser una clave de elemento JSON válida, por lo tanto, el valor especificará que queremos que el nombre del campo sea `-`.

```go
package main

import(
  "fmt"
  "encoding/json"
)

// Profile declares `Profile` structure
type Profile struct {
  Username string `json:"uname"`
  Followers int `json:"followers,omitempty,string"`
}

// Student declares `Student` structure
type Student struct {
  FirstName string `json:"fname"` // `fname` as field name
  LastName string `json:"lname,omitempty"` // discard if value is empty
  Email string `json:"-"` // always discard
  Age int `json:"-,"` // `-` as field name
  IsMale bool `json:",string"` // keep original field name, coerce to a string
  Profile Profile `json:""` // no effect
}

func main() {

  // define `john` struct (pointer)
  john := &Student{
    FirstName: "John",
    LastName: "", // empty
    Age: 21,
    Email: "john@doe.com",
    Profile: Profile{
      Username: "johndoe91",
      Followers: 1975,
    },
  }

  // encode `john` as JSON
  johnJSON, _ := json.MarshalIndent( john, "", "  " )

  // print JSON string
  fmt.Println( string(johnJSON) )
}
```

Lo único extraño del programa anterior es la etiqueta `json: ""` que no hace realmente nada, por lo que la función Marsal la ignora. Todo lo demás debería explicarse por sí mismo. Este programa anterior produce el siguiente resultado.

```shell
{
  "fname": "John",
  "-": 21,
  "IsMale": "false",
  "Profile": {
    "uname": "johndoe91",
    "followers": "1975"
  }
}

Program exited.
```

Las `structure tags` proporcionan una gran ayuda para tratar los campos promocionados. Hasta ahora, hemos aprendido que si los campos promocionados entran en conflicto con los campos de la estructura principal, entonces se seleccionan los campos menos anidados para el cálculo de referencias.

Sin embargo, las etiquetas nos dan más control sobre las promociones de campos usando esta simple regla. Los campos en conflicto (etiquetados o no etiquetados) se agrupan por los nombres de campo JSON y los campos menos anidados se seleccionan para la clasificación.

### Codificar trabajando con  maps

En el ejemplo de mapa anterior, vimos que un `map` con un `string` o `int` se puede codificar como JSON y todas las keys `int` se convierten a JSON. Sin embargo, el mapa en Go puede ser más complejo y sus keys pueden ser de un tipo de datos complejo.

En tales casos, si las claves de un mapa implementan la interfaz `encoding.TextMarshaler, Marshal` intentará obtener la key JSON de la función `MarshalText()` en su lugar, mientras que los valores pueden ser cualquier cosa (como valores de estructura).

## Decodificando JSON

Decodificar JSON es un poco complicado porque necesitamos traducir algunos datos basados en texto en una estructura de datos compleja. Para decodificar JSON en una estructura de datos válida como `map` o `struct`, primero debemos asegurarnos antes de si el JSON es un JSON válido.

```go
func Valid(data []byte) bool
```

Podemos usar la función `json.Valid` para verificar si JSON es válido. Esta función devuelve verdadero si los datos JSON son válidos o falso en caso contrario.

```go
package main

import(
  "fmt"
  "encoding/json"
)

func main() {

  // some JSON data
  data := []byte(`
  {
    "FirstName": "John",
    "Age": 21,
    "Username": "johndoe91",
    "Grades": null,
    "Languages": [
      "English",
      "French"
    ]
  }`)

  // check if `data` is valid JSON
  isValid := json.Valid(data)
  fmt.Println( isValid )
}
```

En el programa anterior, tenemos algunos datos JSON almacenados dentro de la variable `data` y simplemente estamos usando la función `json.Valid(data)` para verificar si `data` contienen un JSON válido (para que se puedan decodificar sin un error). Dado que el JSON representado anteriormente es válido, devuelve verdadero.

```shell
true

Program exited.
```

Como usamos la función `json.Marshal` para codificar datos JSON de una estructura de datos, tenemos que usar la función `json.Unmarshal` **para decodificar datos JSON en una estructura de datos** como `map` o `struct`.

```go
func Unmarshal(data []byte, v interface{}) error
```

La función `Unmarshal` toma los datos JSON como primer argumento y el contenedor `v`, que contendrá los datos como segundo argumento. El argumento `v` es un puntero a una estructura de datos válida o a una interfaz.

Si `v` es nil (excepto por una interfaz nil) o no es un puntero, `Unmarshal` devuelve el error `json.InvalidUnmarshalError`. También devuelve un error si el JSON no se puede decodificar en el valor almacenado en `v`.

```go
package main

import(
  "fmt"
  "encoding/json"
)

// Student declares `Student` structure
type Student struct {
  FirstName, lastName string
  Email string
  Age int
  HeightInMeters float64
}

func main() {

  // some JSON data
  data := []byte(`
  {
    "FirstName": "John",
    "lastName": "Doe",
    "Age": 21,
    "HeightInMeters": 175,
    "Username": "johndoe91"
  }`)

  // create a data container
  var john Student

  // unmarshal `data`
  fmt.Printf( "Error: %v\n", json.Unmarshal( data, &john ) );

  // print `john` struct
  fmt.Printf( "%#v\n", john )
}
```

En el ejemplo anterior, hemos declarado un tipo de estructura `Student` simple con campos básicos. En la función principal, el objeto de datos contiene algunos datos JSON válidos. Algunos de los nombres de campo en este JSON coinciden con la estructura del estudiante.

Más tarde, declaramos una estructura `john` vacía de tipo `Student`. Una estructura vacía tendrá todos sus campos configurados en sus respectivos valores cero (como 0 para enteros tipo `int` y "" para cadenas tipo `string`). Luego pasamos el puntero a la estructura `john` dentro de la función `Unmarshal`. Este programa cuando se ejecuta produce el siguiente resultado.

```shell
Error: <nil>
main.Student{FirstName:"John", lastName:"", Email:"", Age:21, HeightInMeters:175}

Program exited.
```

Como podemos ver en el resultado, el campo `lastName` permanece vacío porque no se exportó a pesar de que existía el mismo campo en los datos JSON. Cualquier campo JSON adicional, si no se declara en la estructura, no será decodificado del JSON a la estructura.

También puede observar que el campo `HeightInMeters` del JSON es un entero tipo `int`, pero se interpretó con éxito en `float64`. Como sabemos, JSON solo tiene un tipo de datos numérico para representar números, por lo tanto, `Unmarshal` coaccionará el valor numérico a un float64.

Unfortunately, the reverse is not true. Unmarshal will return an error if JSON field contains a floating-point number and the respective field is designated as an int. Hence, it’s better to designate a number field as float32/64 in scenarios where a field can either be an integer or a floating-point number.

Desafortunadamente, lo contrario no es cierto. `Unmarshal` devolverá un error si el campo JSON contiene un número de punto flotante y el campo respectivo está designado como `int`. Por lo tanto, es mejor designar un campo numérico como float32/float64 en escenarios donde un campo puede ser un número entero o un número de punto flotante.

```shell
Error: json: cannot unmarshal number 1.75 into Go struct field Student.HeightInMeters of type int
```

> Si un campo en JSON no contiene el valor del tipo de datos declarado en la estructura, Unmarshal no forzará ese valor a un tipo de datos apropiado del campo y, en su lugar, devolverá un error.

### Manejando estructuras de datos complejas

Si un JSON contiene datos complejos, como un objeto o una array, entonces una estructura debe declarar los campos de los tipos apropiados en orden para desarmar el JSON sin error.

```go
package main

import(
  "fmt"
  "encoding/json"
)

// Profile declares `Profile` structure
type Profile struct {
  Username string
  Followers int
}

// Student declares `Student` structure
type Student struct {
  FirstName, lastName string
  HeightInMeters float64
  IsMale bool
  Languages [2]string
  Subjects []string
  Grades map[string]string
  Profile Profile
}

func main() {

  // some JSON data
  data := []byte(`
  {
    "FirstName": "John",
    "HeightInMeters": 1.75,
    "IsMale": null,
    "Languages": [ "English", "Spanish", "German" ],
    "Subjects": [ "Math", "Science" ],
    "Grades": { "Math": "A" },
    "Profile": {
      "Username": "johndoe91",
      "Followers": 1975
    }
  }`)

  // create a data container
  var john Student = Student{
    IsMale: true,
    Subjects: []string{ "Art" },
    Grades: map[string]string{ "Science": "A+" },
  }

  // unmarshal `data`
  fmt.Printf( "Error: %v\n", json.Unmarshal( data, &john ) );

  // print `john` struct
  fmt.Printf( "%#v\n", john )
}
```

The one extra thing we did was to initialize the john struct with some field values. We have initialized the Subjects field of slice data type and Grades field of map data type. This program yields the following result.

En el programa anterior, hemos declarado el tipo de estructura `Profile` y la estructura de `Student` modificada para contener un campo de `Profile` de tipo `Profile`.

También hemos introducido algunos campos nuevos en el tipo de estructura `Student` de estructura de datos de `array`, `slice` y `map`. También hemos proporcionado los datos JSON para estos campos en el valor de los datos.

Lo único que hicimos fue inicializar la estructura `john` con algunos valores de campos. Hemos inicializado el campo `Subjects` del tipo de datos `slice` y el campo `Grades` del tipo de datos map.

Este programa produce el siguiente resultado.

```shell
Error: <nil>
main.Student{
    FirstName: "John",
    lastName: "",
    HeightInMeters: 1.75,
    IsMale: true,
    Languages: [2]string{
        "English",
        "Spanish"
    },
    Subjects: []string{
        "Math",
        "Science"
    },
    Grades: map[string]string{
        "Math": "A",
        "Science": "A+"
    },
    Profile: main.Profile{
        Username: "johndoe91",
        Followers: 1975
    }
}

Program exited.
```

Están sucediendo muchas cosas aquí, así que repasemos los resultados.

  1. Los campos que no se exportan en la estructura o que faltan en el JSON no se eliminan. Si un valor de campo en JSON es nulo y el valor cero de su tipo de campo correspondiente es nulo (como `interface`, `map`, `pointer` o `slice`), el valor se reemplaza por nulo; de lo contrario, ese campo se ignora para la desordenación y conserva su valor original .

  2. Si `Unmarshal` encuentra un tipo de `array` y los valores de `array` en el JSON son más de los que el `array` puede contener, los valores adicionales se descartan. Si los valores del `array` en el JSON son menores que la longitud del `array`, los elementos restantes del `array` se establecen en sus valores cero. El tipo `array` debe ser compatible con los valores del JSON.

  3. Si `Unmarshal` encuentra un tipo `slice`, el `slice` de la estructura se establece en 0 de longitud y los elementos del `array` JSON se agregan uno a la vez. Si el JSON contiene un `array` vacío, `Unmarshal` reemplaza el `slice` en la estructura con un `slice` vacío. El tipo de `slice` debe ser compatible con los valores del JSON.

  4. Si `Unmarshal` encuentra un tipo `map` y el valor del  `map` en la estructura es nulo, se crea un nuevo  `map`  y se añaden los valores del objeto en el JSON. Si el valor del mapa no es nulo, el valor original del `map` se reutiliza y se agregan nuevas entradas. El tipo de `map` debe ser compatible con los valores del JSON

Si `Unmarshal` encuentra un campo de puntero y el valor de ese campo en el JSON es nulo, ese campo se establece en un valor de puntero nulo. Si el campo en JSON no es nulo, se asigna nueva memoria para el puntero en caso de que el puntero sea nulo o se reutilice el valor anterior del puntero.

```go
package main

import(
  "fmt"
  "encoding/json"
)

// Profile declares `Profile` structure
type Profile struct {
  Username string
  Followers int
}

// Student declares `Student` structure
type Student struct {
  FirstName, lastName string
  HeightInMeters float64
  IsMale bool
  Languages [2]string
  Subjects []string
  Grades map[string]string
  Profile *Profile
}

func main() {

  // some JSON data
  data := []byte(`
  {
    "FirstName": "John",
    "HeightInMeters": 1.75,
    "IsMale": null,
    "Languages": [ "English" ],
    "Subjects": [ "Math", "Science" ],
    "Grades": null,
    "Profile": { "Followers": 1975 }
  }`)

  // create a data container
  var john Student = Student{
    IsMale: true,
    Languages: [2]string{ "Korean", "Chinese" },
    Subjects: nil,
    Grades: map[string]string{ "Math": "A" },
    Profile: &Profile{ Username: "johndoe91" },
  }

  // unmarshal `data`
  fmt.Printf( "Error: %v\n\n", json.Unmarshal( data, &john ) );

  // print `john` struct
  fmt.Printf( "%#v\n\n", john )
  fmt.Printf( "%#v\n", john.Profile )
}
```

En el programa anterior, el campo `Subject` se establece en nil explícitamente (será nil implícitamente ya que su valor cero es nil). El campo `Profile` de la estructura del estudiante tiene un puntero al tipo de perfil y se inicializa con el valor del campo `Username`.

En el JSON, hemos establecido el campo `Grades` en nulo. También hemos proporcionado el valor del campo `Profile.Followers`. Veamos el resultado de este programa.

```shell
Error: <nil>

main.Student{FirstName:"John", lastName:"", HeightInMeters:1.75, IsMale:true, Languages:[2]string{"English", ""}, Subjects:[]string{"Math", "Science"}, Grades:map[string]string(nil), Profile:(*main.Profile)(0xc00000c080)}
&main.Profile{Username:"johndoe91", Followers:1975}
```

Como podemos ver en el resultado, el valor del campo del `array` `Languages` fue anulado por los valores en el JSON. El campo `Grades` se estableció en nil porque es un campo tipo `map` (y su valor cero es nil) y su valor en elJSON es nulo.

Dado que el campo `Profile` es un puntero y su valor en el struct `john` no es nulo, `Unmarshal allocates` usa el valor del puntero existente y los valores de campo asignados del JSON. Si el valor del campo `Profile` en JSON fuera nulo, John habría establecido incondicionalmente el valor del campo `Profile` en cero.

### Campos promocionados

Si una `struct` contiene un campo de `struct` anidado de forma anónima, el campo de `struct` anidado se promoverá a la estructura principal. Por lo tanto, el JSON debe contener los valores de campo en el objeto principal

```go
package main

import(
  "fmt"
  "encoding/json"
)

// Profile declares `Profile` structure
type Profile struct {
  Username string
  Followers int
}

// Account declares `Account` structure
type Account struct {
  IsMale bool
  Email string
}

// Student declares `Student` structure
type Student struct {
  FirstName, lastName string
  HeightInMeters float64
  IsMale bool
  Profile
  Account
}

func main() {

  // some JSON data
  data := []byte(`
  {
    "FirstName": "John",
    "HeightInMeters": 1.75,
    "IsMale": true,
    "Username": "johndoe91",
    "Followers": 1975,
    "Account": { "IsMale": true, "Email": "john@doe.com" }
  }`)

  // create a data container
  var john Student

  // unmarshal `data`
  fmt.Printf( "Error: %v\n", json.Unmarshal( data, &john ) );

  // print `john` struct
  fmt.Printf( "%#v\n", john )
}
```

Como podemos ver en el ejemplo anterior, las estructuras de `Account` y `Profile` se anidan de forma anónima dentro de la estructura `Student`. En los datos JSON, hemos agregado los campos `Username` y `Followers` de la estructura del perfil en el objeto principal, pero los datos del campo `Account` se proporcionan a través del objeto `Account`.

```shell
Error: <nil>

main.Student{FirstName:"John", lastName:"", HeightInMeters:1.75, IsMale:true, Profile:main.Profile{Username:"johndo
e91", Followers:1975}, Account:main.Account{IsMale:false, Email:""}}
```

From the result of this program above, we can see that Account field was not unmarshalled because this field is anonymously nested struct and it expected the field values to be present on the parent object.

A partir del resultado del programa anterior, podemos ver que el campo `Account` no se eliminó porque este campo es una estructura anidada anónimamente y esperaba que los valores del campo estuvieran presentes en el objeto principal.

### Decodificando usando Structure Tags

En la lección de codificación de JSON, aprendimos que las structure tags pueden ser muy útiles para decidir los nombres de los campos y los criterios de omisión. También podemos usar las structure tags para interpolar nombres de campo JSON para  nombres de campo de struct.

```go
package main

import(
  "fmt"
  "encoding/json"
)

// Profile declares `Profile` structure
type Profile struct {
  Username string `json:"uname"`
  Followers int `json:"f_count"`
}

// Student declares `Student` structure
type Student struct {
  FirstName string `json:"fname"`
  LastName string `json:"-"` // discard
  HeightInMeters float64 `json:"height"`
  IsMale bool  `json:"male"`
  Languages []string `json:",omitempty"`
  Profile Profile `json:"profile"`
}

func main() {

  // some JSON data
  data := []byte(`
  {
    "fname": "John",
    "LastName": "Doe",
    "height": 1.75,
    "IsMale": true,
    "Languages": null,
    "profile": {
      "uname": "johndoe91",
      "Followers": 1975
    }
  }`)

  // create a data container
  var john Student = Student{
    Languages: []string{ "English", "French" },
  }

  // unmarshal `data`
  fmt.Printf( "Error: %v\n", json.Unmarshal( data, &john ) );

  // print `john` struct
  fmt.Printf( "%#v\n", john )
}
```

En el programa anterior, hemos etiquetado todos los campos de estructura con los nombres de campo JSON apropiados. Sin embargo, hemos etiquetado el campo `LastName` con `json: "-"`, lo que significa que este campo **no se considerará para la decodificación**.

Como puede ver en el resultado anterior, el campo `IsMale` en el JSON no se decodifico porque está etiquetado como nombre de campo `male` (lo mismo con el campo anidado `Followers`).

Unfortunately, omitempty option does not work. You might’ve expected that the Unmarshal function will ignore a field if its value is null in the JSON by looking at the omitempty option value but unfortunately that’s not the case. I hope, Go will consider adding this feature in the future.

Desafortunadamente, la opción `omitempty` no funciona. Es posible que esperara que la función `Unmarshal` ignorara un campo si su valor es nulo en el JSON al observar el valor de la opción `omitempty`, pero desafortunadamente ese no es el caso.

### Decodificar trabajando con maps

Since a JSON contains string keys and values of supported data types, a map of type map[string]interface{} is a suitable candidate for storing JSON data. We can pass a pointer to nil or non-nil pointer of the map to the Unmarshal function and all JSON field values will be populated inside the map.

Dado que un JSON contiene claves de tipo `string` y valores de tipos de datos admitidos, un `map` de tipo `map[string]interface{}` es un candidato adecuado para almacenar datos JSON. Podemos pasar un puntero a un puntero nulo o no nulo del mapa a la función `Unmarshal` y todos los valores de los campos del JSON se completarán dentro del `map`.

```go
package main

import(
  "fmt"
  "encoding/json"
)

// Student declares `Student` map
type Student map[string]interface{}

func main() {

  // some JSON data
  data := []byte(`
  {
    "id": 123,
    "fname": "John",
    "height": 1.75,
    "male": true,
    "languages": null,
    "subjects": [ "Math", "Science" ],
    "profile": {
      "uname": "johndoe91",
      "f_count": 1975
    }
  }`)

  // create a data container
  var john Student

  // unmarshal `data`
  fmt.Printf( "Error: %v\n", json.Unmarshal( data, &john ) );

  // print `john` map
  fmt.Printf( "%#v\n\n", john )

  // iterate through keys and values
  i := 1;
  for k, v := range john {
    fmt.Printf("%d: key (`%T`)`%v`, value (`%T`)`%#v`\n", i, k, k, v, v)
    i++;
  }
}
```

En el ejemplo anterior, hemos creado un `map Student` de tipos de datos los cuales contienen claves de tipo `string` y valores de tipos de datos `interfaz {}`. Si nuestro JSON contiene valores de tipo de datos específicos, entonces podemos personalizar el tipo de datos del `map` de acuerdo con él, por ejemplo `map[string]float64` para valores JSON de tipo `int`.

Hemos creado una variable de `map` vacía john que es nula. Hemos pasado el puntero de john a la función `Unmarshal` y esta función inicializará un mapa para almacenar los datos decodificados JSON.

From the above result, we can see that all fields got populated inside the john map. But take a look at is the data types of the map values. There are certain rules Unmarshal functions follows to store the JSON values in a map.

A partir del resultado anterior, podemos ver que todos los campos se completaron dentro del mapa de John. Pero echemos un vistazo a los tipos de datos de los valores del mapa. Hay ciertas reglas que siguen las funciones de `Unmarshal` para almacenar los valores JSON en un mapa.

1. Un valor JSON `string` es almacenado como `string`.
2. Un valor JSON `number` (`int` or `float`) es almacenado como `float64`.
3. Un valor JSON `boolean` es almacenado como `bool`.
4. Un valor JSON `null` es almacenado como `nil`.
5. Un valor JSON `array` es almacenado como `slice` de tipo `[]interface{}`.
6. Un valor JSON `object` es almacenado como un `map` de tipo `map[string]interface{}`.

Las cosas interesantes para mirar en el resultado son los valores del `array` y el `objeto`. Los valores del `array` se almacenaron en un slice de tipo `[]interfaz{}` y los valores de `objeto` se almacenan en un mapa de tipo `map[string]interfaz{}`.

Como sabemos, un formato JSON válido puede ser un `object` (como en el ejemplo anterior) o un `array`. Dado que `Unmarshal` es capaz de asignar memoria para un puntero, así como también puede crear contenedores para contener datos JSON decodificados por sí solo, podemos almacenar datos JSON complejos sin definir un tipo de contenedor.

```go
package main

import(
  "fmt"
  "encoding/json"
)

func main() {

  // some JSON data
  data := []byte(`
  {
    "id": 123,
    "fname": "John",
    "height": 1.75,
    "male": true,
    "languages": null,
    "subjects": [ "Math", "Science" ],
    "profile": {
      "uname": "johndoe91",
      "f_count": 1975
    }
  }`)

  // create a data container
  var john interface{}
  fmt.Printf( "Before: `type` of `john` is %T and its `value` is %v\n", john, john )

  // unmarshal `data`
  fmt.Printf( "Error: %v\n", json.Unmarshal( data, &john ) );
  fmt.Printf( "After: `type` of `john` is %T\n\n", john )

  // print `john` map
  fmt.Printf( "%#v\n", john )
}
```

En el ejemplo anterior, hemos creado el contenedor `john` de tipo `interfaz{}`. Su valor predeterminado es `nil` porque, por el momento, la interfaz no tiene un valor concreto. Estamos pasando un puntero de la interfaz john como argumento de la función `Unmarshal`.

La función `Unmarshal` extraerá el valor concreto de la interfaz y si lo encuentra nulo, le asignará un tipo de datos adecuado para almacenar datos JSON decodificados. El programa anterior produce el siguiente resultado.

```shell
Before: `type` of `john` is <nil> and its `value` is <nil>
Error: <nil>
After: `type` of `john` is map[string]interface {}
map[string]interface {}{"fname":"John", "height":1.75, "id":123, "languages":interface {}(nil), "male":true, "profile":map[string]interface {}{"f_count":1975, "uname":"johndoe91"}, "subjects":[]interface {}{"Math", "Science"}}
```

Como podemos ver en el resultado anterior, la función `Unmarshal` almacenó un `map` si el tipo `map[string]interfaz{}` es como el valor concreto de la interfaz. Si el JSON fuera un `array` en lugar de un `object`, habría almacenado un `slice` de tipo `[]interface` como valor concreto.

Recuerde, por el momento, `john` es una `interface` y para acceder a los valores del `map`, primero necesitamos extraer el valor concreto de la interfaz. Para hacer eso, necesitamos usar la sintaxis de aserción de tipo de la interfaz.

```go
johnData := john.(map[string]interface{})
```

### Usando Unmarshaler and TextUnmarshaler

Un `struct` puede asumir la responsabilidad de `unmarshaling` los datos JSON por sí solo. En tal caso, el valor del campo debe implementar la interfaz `json.Unmarshaler` que proporciona la declaración del método `UnmarshalJSON`.

```go
type Unmarshaler interface {
    UnmarshalJSON([]byte) error
}
```

Este método se utiliza para delegar la responsabilidad de `unmarshalling` un campo atrás en el propio campo. Si `Unmarshal` encuentra un campo de tipo `Unmarshaler`, llamará a la función `UnmarshalJSON` con datos JSON de ese campo (incluso si es nulo) y será responsabilidad de ese campo inicializar/asignar un valor.

```go
package main

import(
  "fmt"
  "strings"
  "encoding/json"
)

// Profile declares `Profile` structure
type Profile struct {
  Username string
  Followers string
}

// UnmarshalJSON - implement Unmarshaler interface
func ( p *Profile ) UnmarshalJSON( data []byte ) error {

  // unmarshal JSON
  var container map[string]interface{}
  _ = json.Unmarshal( data, &container )

  fmt.Printf( "container: %T / %#v\n\n", container, container )

  // extract interface values
  iuserName, _ := container[ "Username" ]
  ifollowers, _ := container[ "f_count" ]

  fmt.Printf( "iuserName: %T/%#v\n", iuserName, iuserName )
  fmt.Printf( "ifollowers: %T/%#v\n\n", ifollowers, ifollowers )

  // extract concrete values
  userName, _ := iuserName.(string) // get `string` value
  followers, _ := ifollowers.(float64) // get `float64` value

  fmt.Printf( "userName: %T/%#v\n", userName, userName )
  fmt.Printf( "followers: %T/%#v\n\n", followers, followers )

  // assign values
  p.Username = strings.ToUpper(userName)
  p.Followers = fmt.Sprintf( "%.2fk", followers/1000 )

  return nil;
}

// Student declares `Student` structure
type Student struct {
  FirstName string
  Profile Profile
}

func main() {

  // some JSON data
  data := []byte(`
  {
    "FirstName": "John",
    "Profile": {
      "Username": "johndoe91",
      "f_count": 1975
    }
  }`)

  // create a data container
  var john Student

  // unmarshal `data`
  fmt.Printf( "Error: %v\n", json.Unmarshal( data, &john ) );

  // print `john` struct
  fmt.Printf( "%#v\n", john )
}
```

En el programa anterior, el valor del campo `Profile` implementa la interfaz `Unmarshaler`. Este programa produce el siguiente resultado.

```shell
container: map[string]interface {} / map[string]interface {}{"Username":"johndoe91", "f_count":1975}

iuserName: string/"johndoe91"
ifollowers: float64/1975

userName: string/"johndoe91"
followers: float64/1975

Error: <nil>
main.Student{FirstName:"John", Profile:main.Profile{Username:"JOHNDOE91", Followers:"1.98k"}}
```

Si un campo implementa la interfaz `encoding.TextUnmarshaler` y el valor del campo JSON es un `string`, `Unmarshal` llama al método `UnmarshalText` de ese valor con la forma sin comillas de la cadena.

## Codificador y Decodificador

Go proporciona los tipos de estructura json/Encoder y json/Decoder para codificar JSON de un flujo de datos y decodificar JSON en un flujo de datos. Esto es útil para procesar JSON ya que hay algunos datos disponibles.

### Codificador

El tipo de estructura json/Encoder le permite crear una estructura que contiene un objeto `io.Writer` y proporciona el método `Encode()` para codificar JSON desde un objeto y escribir en este objeto `io.Writer`.

```go
func (enc *Encoder) Encode(v interface{}) error
```

Pero primero, necesitamos crear un objeto `*Encoder` a partir de un `io.Writer` usando la función `NewEncoder`.

```go
func NewEncoder(w io.Writer) *Encoder
```

Cada vez que se llama al método `Encode()`, JSON se calcula desde `v` y se agrega a la `w`.

```go
package main

import (
  "bytes"
  "encoding/json"
  "fmt"
)

type Person struct {
  Name string
  Age  int
}

func main() {

  // create a buffer to hold JSON data
  buf := new(bytes.Buffer)
  // create JSON encoder for `buf`
  bufEncoder := json.NewEncoder(buf)

  // encode JSON from `Person` structs
  bufEncoder.Encode(Person{"Ross Geller", 28})
  bufEncoder.Encode(Person{"Monica Geller", 27})
  bufEncoder.Encode(Person{"Jack Geller", 56})

  // print contents of the `buf`
  fmt.Println(buf) // calls `buf.String()` method
}
```

```shell
{"Name":"Ross Geller","Age":28}
{"Name":"Monica Geller","Age":27}
{"Name":"Jack Geller","Age":56}


Program exited.
```

> En el ejemplo anterior, hemos usado un `*Buffer` porque implementa la interfaz `io.Writer` implementando el método `Write`.

### Decodificador

El tipo de estructura json/Decoder le permite crear una estructura que contiene un objeto `io.Reader` y proporciona el método `Decode()` para decodificar JSON de este objeto `io.Writer` y escribir en un objeto.

```go
func (dec *Decoder) Decode(v interface{}) error
```

Si se han leído todas las líneas del `io.Reader`, la siguiente llamada a `Decode` devuelve el `io.EOF error`. Pero primero, necesitamos crear un objeto `*Decoder` desde un `io.Reader` usando la función `NewDecoder`.

```go
func NewDecoder(r io.Reader) *Decoder
```

Cada vez que se llama al método `Encode()`, JSON se desmarca de `r` leyendo una línea (al final con un carácter de nueva línea) y se guarda en `v`.

```go
package main

import (
  "encoding/json"
  "fmt"
  "strings"
)

type Person struct {
  Name string
  Age  int
}

func main() {

  // create a strings reader
  jsonStream := strings.NewReader(`
    {"Name":"Ross Geller","Age":28}
    {"Name":"Monica Geller","Age":27}
    {"Name":"Jack Geller","Age":56}
  `)

  // create JSON decoder using `jsonStream`
  decoder := json.NewDecoder(jsonStream)

  // create `Person` structs to hold decoded data
  var ross, monica Person

  // decode JSON from `decoder` one line at a time
  decoder.Decode(&ross)
  decoder.Decode(&monica)

  // see value of the `ross` and `monica`
  fmt.Printf("ross: %#v\n", ross)
  fmt.Printf("monica: %#v\n", monica)
}
```

```shell
ross: main.Person{Name:"Ross Geller", Age:28}
monica: main.Person{Name:"Monica Geller", Age:27}

Program exited.
```
