package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	//Recurso para permitir importações que são utilizadas em runtime
	_ "github.com/lib/pq"
)

func conectWithDb() *sql.DB {
	conectionString := "user=root dbname=loja_simples password=root host=postgre sslmode=disable"
	db, err := sql.Open("postgres", conectionString)
	if err != nil {
		panic(err.Error())
	}

	return db
}

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
	db := conectWithDb()
	defer db.Close()
	//Função para definir a URL que vamos responder e qual a função de resposta
	http.HandleFunc("/", index)

	porta := 8000
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
