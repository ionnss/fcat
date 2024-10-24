package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Função para carregar citações do arquivo quotes.txt
func carregarCitations() ([]string, error) {
	var citas []string

	arquivo, err := os.Open("quotes.txt")
	if err != nil {
		return nil, err
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linha := scanner.Text()
		citas = append(citas, linha)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return citas, nil
}

func getRandomCitation(citations []string) string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(citations))
	return citations[randomIndex]
}

func imprimirGato(citacao string) {
	tamanhoBalao := len(citacao)
	linhaTopo := " " + strings.Repeat("-", tamanhoBalao+2)
	linhaCima := "/" + strings.Repeat(" ", tamanhoBalao+2) + "\\"
	linhaCitacao := fmt.Sprintf("| %s |", citacao)
	linhaBaixo := "\\" + strings.Repeat(" ", tamanhoBalao+2) + "/"
	linhaBase := " " + strings.Repeat("-", tamanhoBalao+2)

	// Imprime o balão de diálogo
	fmt.Println(linhaTopo)
	fmt.Println(linhaCima)
	fmt.Println(linhaCitacao)
	fmt.Println(linhaBaixo)
	fmt.Println(linhaBase)

	// Imprime o desenho do gato em ASCII com barras escapadas corretamente
	fmt.Println(`        \   
             \  
               ,_     _
               |\\_,-~/
               / _  _ |    ,--.
              (  @  @ )   / ,-'
               \  _T_/-._( (
               /         \. \
              |         _  \ |
               \ \ ,  /      |
                || |-_\__   /
               ((_/` + "`(____,-'`)")
}

func main() {
	citations, err := carregarCitations()
	if err != nil {
		fmt.Println("Erro ao carregar citações:", err)
		return
	}

	randomCitation := getRandomCitation(citations)
	imprimirGato(randomCitation)
}
