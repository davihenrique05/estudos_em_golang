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
	conectionString := "host=localhost user=root password=root dbname=simple_store port=5432 sslmode=disable"
	db, err := sql.Open("postgres", conectionString)
	if err != nil {
		panic(err.Error())
	}

	return db
}

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// Parse glob busca todos os arquivos que dão match no string entregue via parametro e converte
// html para template
var temp = template.Must(template.ParseGlob("templates/*html"))

func main() {
	//Função para definir a URL que vamos responder e qual a função de resposta
	http.HandleFunc("/", index)

	porta := 8000
	fmt.Println("Servidor inicilizado em porta", porta)
	http.ListenAndServe(fmt.Sprint(":", porta), nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{}
	db := conectWithDb()
	defer db.Close()
	selectProducts, err := db.Query("select * from produtos")

	if err != nil {
		fmt.Println(err)
	}

	for selectProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p := Produto{Nome: nome, Descricao: descricao, Preco: preco, Quantidade: quantidade}

		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
