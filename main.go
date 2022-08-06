package main

import (
	"html/template"
	"net/http"
)

type Produtos struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produtos{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 8},
		{"Tenis", "Confortavel", 89, 3},
		{"Fone", "Top", 70, 2},
	}

	temp.ExecuteTemplate(w, "Index", produtos)

}
