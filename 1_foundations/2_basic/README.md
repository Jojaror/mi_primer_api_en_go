# Aplicación Web Básica
````go
import (
	"log"
	"net/http"
)
````
Se comienza importando los paquetes ```log``` para registrar mensajes y ```net/http``` para funcionalidades de servidor y cliente HTTP.
````go
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola desde Mi Primera API con Go"))
}
````
Lo primero que necesitamos es definir una función para el **Request Handler**, en nuestro caso llamada ```home``` para manejar las peticiones HTTP, con los parametros de ```w http.ResponseWriter``` que es una interfaz que representa la respuesta que nuestro servidor enviara al cliente y ```r *http.Request``` que es un puntero que contiene la informacion sobre la petición entrante.

````go
w.Write([]byte("Hola desde Mi Primera API con Go"))
````

```w.Write()``` Es un método para escribir datos en el cuerpo de la respuesta HTTP. ```[]byte("Hola desde Mi Primera API con Go")``` El método Write espera recibir los datos como un "slice de bytes" *([]byte)*. Por eso, convertimos la cadena de texto antes de enviarla. Este será el texto que el usuario verá en su navegador.

Lo segundo es un **Router** (servemux) que almacena el mapeo entre las rutas URL de la app y los respectivos *handlers*.

````go
mux := http.NewServeMux()
````
Se llama a ```http.NewServeMux()``` para crear un nuevo ServeMux. Actúa como un enrutador (o multiplexor). Su trabajo es recibir las peticiones HTTP entrantes y, basándose en la URL solicitada, decidir qué función manejadora (como nuestra función home) debe ejecutarse para esa petición. Se asigna este nuevo enrutador a la variable ```mux```.

````go
mux.HandleFunc("/", home)
````
Se usa el método ```HandleFunc``` del ```mux``` para registrar una ruta y su manejador.
```"/"``` Es el patrón de la URL. En este caso, **/** es la ruta raíz. ```home``` Es el nombre de la función que se ejecutará cuando una petición coincida con el patrón **/**. Le estamos diciendo al **mux** "Cuando llegue una petición para /, llama a la función home". Luego Hacemos una impresión en la consola del server para mostrar que nuestra web app se está ejecutando.

Finalmente, necesitamos un **Web Server**, que afortunadamente **Go** lo establece facilmente.

````go
err := http.ListenAndServe(":9876", mux)
````
```http.ListenAndServe``` Es la función del paquete **net/http** que crea y arranca el servidor.```":9876"``` Es la dirección de red en la que el servidor escuchará las peticiones entrantes. El ```:``` inicial significa "escuchar en todas las interfaces de red disponibles" y ```9876``` es el número de puerto. Le pasamos nuestro enrutador ```mux``` al servidor. ```err := ...``` La función **ListenAndServe** es bloqueante. Esto significa que la ejecución del programa se detiene en esta línea y el servidor se queda escuchando peticiones indefinidamente. Solo retorna si ocurre un error al intentar iniciar el servidor.

````go
log.Fatal(err):
````
Esta línea solo se alcanza si **ListenAndServe** retorna un error. ```log.Fatal()``` Imprime el mensaje en la consola y luego termina inmediatamente la ejecución del programa (llamando a os.Exit(1)).

> **NOTA**: **servemux** trata a la URL **/** como un atrapa todo. Asi que en este momento, todas las peticiones HTTP seran manejadas por la funcion **home** sin importar su ruta.