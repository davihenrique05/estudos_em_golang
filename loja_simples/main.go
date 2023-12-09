package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// Parse glob busca todos os arquivos que dão match no string entregue via parametro e converte
// html para template
var temp = template.Must(template.ParseGlob("templates/*html"))

func main() {
	porta := 8000

	//Função para definir a URL que vamos responder e qual a função de resposta
	http.HandleFunc("/", index)
	fmt.Println("Servidor inicilizado em porta", porta)
	http.ListenAndServe(fmt.Sprint(":", porta), nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39.90, Quantidade: 5},
		{"Tênis", "Ortopédico e confortável", 89.60, 3},
		{"Fone", "Sem fio e muito bom", 59, 2},
		{"Produto novo", "Muito legal", 1.99, 2}}

	temp.ExecuteTemplate(w, "Index", produtos)
}
