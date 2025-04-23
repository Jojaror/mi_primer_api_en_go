# Creando un módulo

El propósito principal de la ruta del módulo es servir como un identificador único para tu código en el ecosistema Go.

## ¿Cómo se logra la Unicidad Global?

La convención aceptada es usar una ruta que corresponda a la ubicación donde alojas tu código fuente, típicamente una URL de un repositorio.


````bash
go mod init github.com/Jojaror/mi_primer_api_en_go
````

El comando creará un archivo go.mod en la carpeta proyecto.

## ¡ Hola mundo !

````bash
package main
````
```package``` Es una palabra clave en Go que se usa para declarar a qué paquete pertenece el archivo actual. ```main``` Es un nombre de paquete especial. Cuando le dices a Go que un archivo pertenece al paquete main, le estás indicando que este código debe compilarse para ser un programa ejecutable.

````bash
import "fmt"
````
```import``` Es una palabra clave se usa para traer código de otros paquetes a tu archivo actual. ```"fmt"``` Es el nombre del paquete que estamos importando. Este es un paquete estándar de Go.

````bash
func main() { ... }
````
```func``` Es la palabra clave en Go para definir una función. ```main``` Es el nombre de la función. La función llamada main dentro del paquete main es muy especial. Es el punto de entrada de tu programa. ```()``` Los paréntesis después del nombre de la función son para los parámetros. En este caso, la función main no recibe ningún parámetro. ```{ ... }``` Las llaves definen el cuerpo de la función.

````bash
fmt.Println("Hola mundo!")
````
Accedemos al paquete ```fmt``` que importamos y usamos la función ```Println``` que viene de *Print Line (imprimir linea)* pasandole como argumento un string ```"Hola mundo!"```.

## Ejecutando la App

````bash
go run 1_foundations/1_setup/main.go
````
