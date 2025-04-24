# Composición HTML

````go
import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)
````

```"html/template"``` Este es un paquete crucial de la librería *estándar* de **Go**. Proporciona herramientas para trabajar con plantillas **HTML**. Su función principal es permitirte escribir **HTML** en archivos separados (.html) e inyectar datos dinámicos desde tu código **Go** de forma segura (ayuda a prevenir ataques como Cross-Site Scripting - XSS).

````go
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	ts, err := template.ParseFiles("./ui/html/pages/home.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
````
La línea simple ```w.Write([]byte("Hello from MyGoAPI"))``` ha sido reemplazada por completo por lógica para manejar plantillas HTML.

````go
ts, err := template.ParseFiles("./ui/html/pages/home.html")
````
```template.ParseFiles(...)``` Esta función del paquete ```html/template``` lee uno o más archivos (en este caso, solo uno) desde el disco, los analiza sintácticamente (parsea) como plantillas HTML de Go, y crea un objeto ```*template.Template``` que representa la plantilla lista para usarse.
```"./ui/html/pages/home.html"``` Es la ruta (path) relativa al archivo HTML que contiene la plantilla para la página de inicio. ```ts, err := ...``` La función ```ParseFiles``` devuelve dos valores: ```ts``` (una variable que contendrá la plantilla parseada, comúnmente llamada **ts** por *"template set"*) y ```err``` (un posible error). **err** no será ```nil``` si el archivo no existe, no se puede leer, o contiene errores de sintaxis de plantilla.

````go
http.Error(w, "Internal Server Error", http.StatusInternalServerError)
````
Si falla el parseo de la plantilla, se envía una respuesta genérica **HTTP 500 Internal Server Error** al cliente. Se usa ```500``` porque es un problema del lado del servidor (no pudo cargar/procesar su propio archivo). Es importante no enviar el ```err.Error()``` directamente al cliente, ya que podría revelar información sensible sobre la estructura de archivos del servidor. Usamos la constante ```http.StatusInternalServerError``` en lugar del número **500**.

````go
err = ts.Execute(w, nil)
````
Si el parseo fue exitoso, esta línea ejecuta la plantilla. ```ts.Execute(...)``` Es un método del objeto **template.Template (ts)** que renderiza la plantilla HTML.
```w``` El primer argumento es dónde se escribirá la salida HTML renderizada. Le pasamos nuestro **http.ResponseWriter (w)** para que el HTML resultante se envíe directamente como respuesta al navegador del cliente. ```nil``` El segundo argumento es para pasar datos dinámicos a la plantilla. Aquí se pasa **nil**, lo que significa que esta plantilla en particular (home.html) no necesita ningún dato variable proveniente del código **Go** para renderizarse. ```err = ...``` El método ```Execute``` también puede devolver un error si algo sale mal durante la renderización.

````go
files := []string{
	"./ui/html/base.html",
	"./ui/html/partials/nav.html",
	"./ui/html/pages/home.html",
}
````
Ahora bien, el servidor Go nos permite separar la lógica de nuestro HTML, para lo cual se declara una nueva variable llamada ```files```. ```[]string``` Indica que files es un "slice" (un arreglo dinámico) que contendrá cadenas de texto (string). Se inicializa el slice con tres cadenas de texto. Cada cadena es la ruta a un archivo de plantilla HTML: ```"./ui/html/base.html"``` El archivo que contiene el layout base (estructura HTML principal, <html>, <head>, <body>, etc.). ```"./ui/html/partials/nav.html"``` Un fragmento reutilizable para la barra de navegación. ```"./ui/html/pages/home.html"``` El contenido específico para la página de inicio.

````go
ts, err := template.ParseFiles(files...)
````
La forma de llamar a ParseFiles debe cambiar. ```files...``` En lugar de pasar una sola ruta como antes, ahora pasamos la variable files seguida de ```...```. El operador ```...``` desempaqueta el slice files, pasando cada elemento del slice (cada ruta de archivo) como un argumento individual a ```ParseFiles```. El resultado es que ```ParseFiles``` ahora lee y analiza los tres archivos HTML, creando un conjunto de plantillas (template set) almacenado en ```ts```. Esto es fundamental para que las plantillas puedan incluirse unas a otras usando ```{{template "nombre"}}```.

````go
err = ts.ExecuteTemplate(w, "base", nil)
````
La forma de ejecutar la plantilla debe cambiar de ```ts.Execute(...)``` a ```ts.ExecuteTemplate(...)``` Se usa este método cuando has parseado un conjunto de plantillas (como hicimos con ParseFiles y varios archivos) y quieres empezar la renderización ejecutando una plantilla nombrada específica dentro de ese conjunto.
```"base"``` Es el nombre de la plantilla que queremos ejecutar inicialmente. ```nil``` Sigue siendo el espacio para pasar datos a la plantilla (ninguno en este caso).
