package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	jsonFile, err := os.Open("src/shoes.json")
	if err != nil {
		log.Fatal("Não foi possível abrir o arquivo")
	}

	data, _ := ioutil.ReadAll(jsonFile)
	countPairsOfShoes(data)
}

type Shoe struct {
	Number int    `json:"number"`
	Feet   string `json:"feet"`
}

func countPairsOfShoes(data []byte) {
	var shoes []Shoe
	
	if err := json.Unmarshal(data, &shoes); err != nil {
		log.Fatal("Formato inválido")
	}

	var pairsCounter int = 0
	var rights = make(map[int]int)
	var lefts = make(map[int]int)
	min := 0
	max := 0

	for _, v := range shoes {
		if strings.ToLower(v.Feet) == "r" {
			rights[v.Number]++
		}
		if strings.ToLower(v.Feet) == "l" {
			lefts[v.Number]++
		}

		min = chooseMinor(min, v.Number)
		max = chooseGreater(max, v.Number)
	}

	for i := min; i <= max; i++ {
		_, rightHasKey := rights[i]
		_, leftHasKey := lefts[i]

		if rightHasKey && leftHasKey {
			pairsCounter += chooseMinor(rights[i], lefts[i])
		}
	}

	fmt.Println(pairsCounter)
}

func chooseMinor(n1 int, n2 int) int {
	if n1 == 0 || n1 == n2 || n2 < n1 {
		n1 = n2
	}
	return n1
}

func chooseGreater(n1 int, n2 int) int {
	if n2 > n1 {
		n1 = n2
	}
	return n1
}