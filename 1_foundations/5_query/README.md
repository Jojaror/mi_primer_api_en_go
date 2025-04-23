# Busqueda por parametro URL

````go
import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)
````
```"fmt"``` Se ha añadido a la importación. Este paquete proporciona funciones para formatear e imprimir datos. Lo necesitaremos para enviar una respuesta formateada al cliente. ```"strconv"``` Este paquete ofrece funciones para convertir cadenas de texto (strings) a otros tipos de datos (como números) y viceversa.

````go
func view(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Muestra algo en especifico con ID %d...", id)
}
````
````go
id, err := strconv.Atoi(r.URL.Query().Get("id")):
````
```r.URL``` Accedemos a la información de la URL de la petición.```.Query()``` Este método analiza la parte de la query string de la URL (lo que viene después del ?, por ejemplo, ?id=123&lang=es) y la devuelve como un mapa especial (url.Values).```.Get("id")``` Sobre ese mapa de parámetros, usamos el método Get para obtener el valor asociado a la clave ```"id"```. Esto devuelve el valor como una string (cadena de texto). Si no existe el parámetro "id" en la URL, devuelve una cadena vacía "". ```strconv.Atoi(...)``` Tomamos la cadena obtenida (el valor del parámetro "id") y usamos la función ```Atoi``` ("ASCII to Integer") del paquete ```strconv``` para intentar convertirla en un número entero (int).
```id, err := ...``` La función *Atoi* devuelve dos valores: el número entero resultante (si la conversión fue exitosa), que asignamos a la variable id, y un valor de error, que asignamos a la variable err. Si la conversión falla (por ejemplo, si el parámetro "id" no estaba, estaba vacío, o contenía letras como "abc"), err contendrá un valor de error; si la conversión es exitosa, err será nil.

````go
if err != nil || id < 1 { ... }
````
Se comprueban dos condiciones para validar el ID ```err != nil``` Verifica si ocurrió un error durante la conversión de string a entero. ```||``` Operador lógico "O". ```id < 1``` Verifica si el número id (asumiendo que la conversión fue exitosa) es menor que 1. Estamos asumiendo que los IDs válidos deben ser números enteros positivos.

````go
fmt.Fprintf(w, "Muestra algo en especifico con ID %d...", id)
````
```fmt.Fprintf(...)``` Es una función del paquete *fmt* similar a *Printf*, pero en lugar de imprimir en la consola, escribe la cadena formateada en un io.Writer. ```w``` El primer argumento es el *ResponseWriter*, que actúa como un io.Writer. ```"Muestra algo en especifico con ID %d..."``` Es la cadena de formato. ```%d``` es un "verbo" de formato que indica que ahí se debe insertar el valor de una variable entera. ```id``` Es la variable entera (el ID validado) cuyo valor se insertará en lugar de ```%d```.
