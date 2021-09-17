# Mini-guía Go

## Nociones básicas

- -> Editor y compilador on-line: [play.golang.org](https://play.golang.org)

- -> Todos los programas de go tienen un `package main` y una función `main()`.  
Ahi es donde estaría el "inicio" del programa.

- -> No necesitas poner `;` por que el compilador los incluye, pero se pueden poner.

- **COMPILAR**: `go build [*.go]`

- **LIBRERÍA "ESTÁNDAR"**: `import fmt`

## Declaración de variables

->*Lenguaje fuertemente tipado*

Estructura declaración variable:

```go
    var [nombre] [tipo]

    //ejemplos:
    var x,y,z int
    var booleano bool
    var nombre string
```

-> todas las variables tienen un valor inicial, No esta a null.

- En caso de números `int`=0
- En caso `string`= cadena vacía
- En caso de `bool`=`false`

-> NO USAR VARIABLES ES MOTIVO DE ERROR DE COMPILACIÓN.

Forma abreviada de declarar una variable (sin especificar tipo). Tipado dinámico:

```go
func main(){
    x := 23
    nombre := "Manuel"
}
```

### Conversiones de tipos

-> Cuando intentamos concatenar un numero con una cadena, nos da un error de compilación debemos hacer **una conversion de tipos**.

- Usamos paquete `strconv`
- `Itoa`: int -> string
- `Atoi`: string -> int, **DEVUELVE DOS VALORES**.

```go
import(
    "fmt"
    "strconv"
)

edad := 42
edad_str := strconv.Itoa(edad)
fmt.println("mi edad es "+edad_str)
//o fmt.println("mi edad es "+ strconv.Itoa(edad))
```

Atoi:

```go
edad_str:="42"
edad_int,_:=strconv.Atoi(edad)
//Atoi retorna múltiples valores
//Podemos declarar una variable en ese mismo momento PERO DEBEMOS USARLA
//O DESECHAMOS EL VALOR CON _
//En este caso el 2º valor es un código de err
```

## Input y output de datos

- Usa el paquete `fmt`
IMPRIMIR
- `Print`,`Println`, 
- `Printf`(usa formato como en C %)
    - %d int.
    - %v, %s string
    - %t bool
    - %b %e %f float, double etc..


```go
    var name string
    name="Manuel"
    fmt.println(name)
    fmt.println("mi nombre es "+name)//concatenación con +
```

LEER DATOS

- `Scanf`
```go
var edad int
fmt.Scanf(%d\n, &edad)
```
### Usando bufio y os

```go
import(
    "fmt"
    "bufio"
    "os"
)
func main(){
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Ingresa tu nombre: ")
    nombre,err := reader.ReadString('\n')
    if(err != nil){
        fmtPrintln(err)
    }else{
        fmt.println("Hola "+nombre)
    }
}
```

## Condicionales

- ==, !=, <, >, <=, >=

```go
if {

}else if{

}else
```

- no es obligatorio poner ()
- Siempre se debe poner `{}`
- el `{}` debe de estar en la misma linea que el `if`.

## Bucles

- Solo existe For 
    - `for [initialize];[condición];[ejecución]`
    - Se puede quitar esos elementos

```go
for i:=0;i<10;i++{}

for i<10{}

for{break}
```

- sentencia `continue` comienza de nuevo el bucle.
- sentencia `break`termina el bucle. 

### iterador

- `for index, dato := range Array_de_datos`

```GO
    palabras := strings.Split("hola buen señor", "")
    for _, letra := range palabras
        Print(letra)
```

## Arrays, vectores o arreglos

- `var vector [10]int`
- funciona `fmt.Println(vector)`

Formas de declaración:

```go
    vector := [3]int{1,2,3}
    vector2 := [3]int{1,2}

    //tamaño de un vector
    tam := len(vector)  //3

    var matriz [2][2]int
    fmt.Println(matriz)
```

## Slice

Es como un vector dinámico. Estructura basada en array.

- Declaración: `var sli []int`, `var sli []int{1,2,3}`
- Cuando tenemos un slice sin inicializar su valor es `nil`

La estructura del slice tiene 3 datos importantes

- **Puntero al array referenciado**
- **Longitud del array al que apunta**
- **Capacidad**

### Crear un slice con un vector

Se puede crear el slice con un fragmento del vector.

HACER SLICER
```go
arr:=[3]int{1,2,3}

sl1:=arr[:2]    // 1 2
sl2:=arr[1:2]   // 2
sl3:=arr[0:]     //todo el vector 1 2 3

```
### Crear slices con make

- `make([]Tipo, tamaño, capacidad(opcional))`

```go
//make()
slice := make([]int, 3, 5)

len(slice) //3
cap(slice) //5

```

- `append()` añade un elemento

```go
slice = append(slice, 2)
```

### Funcion `copy()`

Copia un array en otro. Copia el minimo de elementos de ambos vectores.

- `copy(destino, fuente)`

```go
    arr:=[]int{1,2,3,4}
    copia:=make([]int, len(arr))

    copy(copia, arr)

    //El tamaño de la copia es el tamaña menor de los dos vectores
    //len = min(copia,fuente)
```

## Punteros

1. una dirección de memoria
2. En lugar del valor, tenemos la dirección en la que esta el valor

DECLARACIÓN:

- `*T` => `*int`, `*string`...
- el valor cero => `nil`
- `&` devuelve la dirección de memoria de una variable
- `*` devuelve el valor de una dirección de memoria

Ejemplo

```go
var x,y *int
num:=5

x = &num
```

## Structs

```go
type User struct{
    edad int
    nombre, apellido string
}

func main(){
    var manu User

    fmt.Println(manu)

    //Asignación indicando nombre de atributo
    manu := User{nombre: "manu", apellido:"HC"}

    fmt.Println(manu)

    //Asignación según orden de declaración de atributos 
    //Se incluyen todos los atributos
    manu := User{20,"",""}

    //new devuelve un puntero
    Manuel:=new(User)
}
```

## Métodos para Structs

- `func (id(this) Struct) nombreFunc() (datoDeveulto){}`

id: identificador del struct en la función
Struct: struct de la función

```go
func (usuario User) nombre_completo() string{
    return usuario.nombre + " " + usuario.apellido
}

func (this *User) set_name(n string) string{
    this.nombre = n
}
```

## Campos anónimos

- Se asemeja mucho a la herencia en lenguajes OO

```go
type Persona struct{
    name string
}

type tutor struct{
    Persona
}

func main(){
    tutor :=Tutor{Persona{"manu"}}
    //HEREDA el nombre de Persona
    fmt.Println(tutor.name)
    //también se puede con:
    fmt.Println(tutor.Persona.name)
}
```

- También podemos usar los métodos y sobrescribir los métodos con `this.campo.func()`

## Interfaces

- Es un tipo de dato
- se define con `type nombre interface{//definiciones de métodos}`
- se pueden usar en un array
- Si los métodos de la interfaz están implementados en una estructura esa interfaz ya es usada por la estructura
- cualquier estructura que implemente la interfaz, sera la interfaz

```go
type User interface{
    Permisos() int
    Nombre() string
}

type Admin struct{
    nombre string
}

func (this Admin) Permisos() int{
    return 5
}

func (this Admin) Nombre() string{
    return this.nombre
}

func auth(user User) string{
    if user.Permisos() > 4{
        return user.Nombre() + " es admin"
    }
}

func main(){
    admin := Admin{"Uriel"}
    fmt.Println(auth(admin))
}
```
---
## Concurrencia `go rotine`

- paquetes -> `time`
- `time.Sleep(1000 * time.Millisecond)`
- con `go` indicamos que una función se ejecute en una hebra distinta a la principal. Con go creamos tantas hebras como queramos, el compilador se encarga de gestionarlas, borrado y eliminado. solo la definimos

- para ejecutar un bloque de código:
```go
go func(){
    //codigo concurrente
}()
```

### channels

- Permite comunicar gorutines unas con otras.  
Definición:

- `channel := make(chan tipo_dato)`

```go
channel := make(chan string) 
go func (channel chan string){
    var name string
    for{
        Scanln(&name)
        channel <- channel
    }
}(channel)

msg := <- channel

print(msg)
```

----

## leer archivos

Varias bibliotecas:

1. -> importar `io/ioutil`

```go
file_data, err := ioutil.ReadFile("./file.txt")//tiene que estar en el mismos dir donde se ejecuta.

if err != nil{
    Print("Error")
}else{
    Print(string(file_data)) //imprime todo el archivo
}
```

2. -> importar `os`, y `bufio`

Permite leer un archivo linea por linea

```go
import(
    "fmt"
    "bufio"
)

func main(){
    file,err := os.Open("./file.txt")

    if err != nil{
        Print("Error")
    }

    scanner := bufio.NewScanner(archivo)

    for scanner.Scan(){
        linea := scanner.Text()
        Println(linea)
    }

    file.Close()
}
```

## Tratar errores con Defer, panic y recover

- `Defer`: cuando un proceso finaliza de forma abrupta o la ejecución llega al final, con `Defer` nos aseguramos que ese código o función se ejecute al terminar.

- `panic`: imprime a detalle el error que ocurrió. ademas, termina el programa sacando de la pila todas las funciones o proceso por encima de el.

```go
if err!=nil{
    panic(err)
}
//esto ya no se ejecuta
```

terminaría la ejecución finalizando main sin que llegase al final.

- `recover`: con `recover` recupera el programa después de un panic.

```go
func main(){
    executeReadFile()
    recover()
    //esto se puede seguir ejecutando
}

func executeReadFile(){
    execute := readFile()
    print(execute)
    //esto no se ejecuta después del panic
}

func readFile() bool
    file_data, err := ReadFile("./file.txt")
    if err != nil{
        panic(err)
    }
```

- `recover` también devuelve el error que proboco panic

---

## Web con Paquete http

- importar `net/http`
- `función http.ListenAndServe(":puerto", )`: lebanta un servidor
- `http.handleFunc("/", handler)`: indica para una url que hacer `en la funcion handler`

- Funcion handler tiene la sigueinte estructura:
    - `func handler(w http.ResponseWriter, r *http.Request)`
    - donde `http.ResponseWriter` es una estructura de string de datos, que usamos apra debolver una `request` http determinado.

```go
import(
    "net/http"
    "io"
)

func main(){
    http.HandleFunc("/paht", handler)
    http.ListenAndServe(":80000", nil) // nil informacion de routing??
}

func hander(w http.ResponseWriter, r *http.Request){
    //recive petición y responde
    io.WriteString(w, "Hola mundo") // escrive el string que se le va a pasar
}
```

### Servir archivos estaticos

Forma facil, pasar el archivo

- `http.ServeFile(w,r,"nomnre del archivo")` para que encuentre el archivo se debe ejecuatr en el mismo sitio donde esta el archivo.

Para que un cliente pueda requerir todos los archivos de un dir:
- `r.URL.Path[1:]`
    - donde `r` es el request, `URL` extrae la url de r despues del `/` indicando con [1:] que empieze desde el caracter 1.
    - Esto tienen un problema, se pueden ver todos los archivo incluso los que no queremos.

Para evitarlo, crear parte publica:

- creamod un directorio `public`
- indicamos que el servidor de archivos comienza en public, asi se combierte en el nuevo espacio de archivos:
```go
    fileServer := http.FileServer(http.Dir("public"))
```
     - ahroa en usamos la funcion Handle indicando en path desde donde empieza y el espacio que usa

```go
func main(){
    fileServer := http.FileServer(http.Dir("public"))

    http.Handle("/", http.StringPrefix("/",fileServer))
    http.ListenAndServer(":8000",nil)
}
```

## Mostrar Estructura en forma de JSON

Go incluye librerias para tratar con json

- import `encodig/json`
- trasforma estructuras a json
- los atributos tienen que tener mayuscular si no los coje. Si queremos que la structura json se vea en con otro nombre, lo modemos indicar con:
    - `Nombre string` ``json: "nombre"`

```go
//....
type Asignatura struct{
    Nombre string ``json: "nombre"
    Horas int ``json: "nombre"
}

type Asignaturas []Asignatura


func main(){
    http.HanleFunc("/", func(w http.ResponseWriter, r *http.Request)){

        asignaturas := Asignaturas{
            Asignatura("spsi", 30),
            Asignatura("dai", 30),
            Asignatura("iv", 30),
        }
        json.NewEncodig(W).Encode(asignaturas)
    }
    //....
}
```

## Creación de paquetes

Para crear paquetes que podamos importar debemos crear un directorio con el nombre del paquete, por ejemplo si queremos crear el paquete demo debemos crear un archivo .go dentro de una carpeta demo, donde el archivo comienza con `package demo`. El archivo no tiene por que llamarse igual.  
no deve contener ninguna funcion `main`

- Las funciones que comienzan por **MINUSCULAS** se tratan como **privadas**
- Las funciones que comienzan por **MAYUSCULAS** se tratan como **publicas**

- Exixte una funcion de configuarion que se ejecuta cuando comienza el paquete -> `init`

```go
var nombre string

func init(){
    nombre = "paquete de prueba"
}
```

- Los **atributos o variables** tambien se pueden exportar, con la mismo conbención.
