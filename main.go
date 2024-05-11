package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/ditto")
	if err != nil {
		println(err.Error())
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		println(err.Error())
		return
	}
	err = os.WriteFile("pokemon.json", body, 0666)
	if err != nil {
		println(err.Error())
		return
	}
	println(string(body))
}
