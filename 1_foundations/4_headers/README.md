# Cabeceras HTTP
Crearemos una nueva ruta en la API para crear algo y haremos que la ruta responda a una peticion HTTP que utilce un determinado metodo

````go
func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed",  http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Crea algo en MyGoAPI"))
}
````
````go
if r.Method != http.MethodPost { ... }
````
Es la principal novedad dentro de la función. Comprueba el método *HTTP* de la petición ```r.Method``` que contiene el método usado por el cliente (ej: "GET", "POST", "PUT", etc.). ```http.MethodPost``` Es una constante definida en el paquete ```net/http``` cuyo valor es la cadena *"POST"*. Usar la constante en lugar de escribir *"POST"* directamente es más seguro y claro. La condición verifica si el método de la petición no es POST.
````go
w.Header().Set("Allow", http.MethodPost)
````
```w.Header()``` Este método del *ResponseWriter (w)* devuelve un mapa que representa las *cabeceras HTTP (headers)* de la respuesta que se enviará al cliente. ```.Set("Allow", http.MethodPost)``` Se llama al método ```Set``` sobre las cabeceras de la respuesta. Esto añade o modifica la cabecera ```Allow```. La cabecera *Allow* se usa para indicar al cliente qué métodos *HTTP* están permitidos para la *URL* solicitada. Aquí, le decimos al cliente que el único método permitido es *POST*. Es una buena práctica al devolver un error *405*.
````go
http.Error(w, "Method Not Allowed",  http.StatusMethodNotAllowed)
````
Si la condición del ```if``` es verdadera, se llama a esta función. Es una forma conveniente de enviar una respuesta de error HTTP completa. ```http.StatusMethodNotAllowed``` Es una constante que representa el código de estado *HTTP 405 Method Not Allowed*. La función ```http.Error``` se encarga de establecer este código de estado y de escribir el mensaje de error en la respuesta.

````bash
curl.exe -i -X POST http://localhost:9876/create
````
Con el anterior comando podremos probar nuestro servidor.