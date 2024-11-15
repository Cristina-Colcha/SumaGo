package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Estructura que se pasa al template
type SumaData struct {
	Num1      int
	Num2      int
	Resultado int
}

func sumaHandler(w http.ResponseWriter, r *http.Request) {
	// Creamos la estructura para almacenar los números y el resultado
	data := SumaData{}

	// Verificamos si el formulario fue enviado (POST)
	if r.Method == http.MethodPost {
		// Leemos los números desde el formulario
		num1 := r.FormValue("num1")
		num2 := r.FormValue("num2")

		// Convertimos los números de string a enteros
		num1Int, err1 := strconv.Atoi(num1)
		num2Int, err2 := strconv.Atoi(num2)

		// Si no podemos convertir, dejamos el resultado como 0
		if err1 == nil && err2 == nil {
			data.Num1 = num1Int
			data.Num2 = num2Int
			data.Resultado = num1Int + num2Int
		}
	}

	// Cargamos y mostramos la plantilla HTML
	tmpl := template.Must(template.New("suma").Parse(`
		<!DOCTYPE html>
		<html lang="es">
		<head>
			<meta charset="UTF-8">
			<title>Suma de Números</title>
		</head>
		<body>
			<h1>Realizar una Suma</h1>
			<form method="post">
				<div>
					<label for="num1">Número 1:</label>
					<input type="number" name="num1" value="{{.Num1}}" required>
				</div>
				<div>
					<label for="num2">Número 2:</label>
					<input type="number" name="num2" value="{{.Num2}}" required>
				</div>
				<button type="submit">Sumar</button>
			</form>

			{{if .Resultado}}
			<h2>Resultado: {{.Resultado}}</h2>
			{{end}}
		</body>
		</html>
	`))

	// Renderizamos la plantilla con los datos
	tmpl.Execute(w, data)
}

func main() {
	// Definimos la ruta para el formulario de la suma
	http.HandleFunc("/", sumaHandler)

	// Iniciamos el servidor web en el puerto 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
