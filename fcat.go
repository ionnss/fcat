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

func dividirTexto(texto string, largura int) []string {
	var linhas []string
	for len(texto) > largura {
		quebra := strings.LastIndex(texto[:largura], " ")
		if quebra == -1 {
			quebra = largura
		}
		linhas = append(linhas, texto[:quebra])
		texto = strings.TrimSpace(texto[quebra:])
	}
	linhas = append(linhas, texto)
	return linhas
}

func imprimirGato(citacao string) {
	// Definindo a largura máxima de cada linha do balão de diálogo
	largura := 40
	linhasCitacao := dividirTexto(citacao, largura)

	// Calcula o tamanho do balão baseado na linha mais longa
	tamanhoBalao := 0
	for _, linha := range linhasCitacao {
		if len(linha) > tamanhoBalao {
			tamanhoBalao = len(linha)
		}
	}

	linhaTopo := " " + strings.Repeat("-", tamanhoBalao+2)
	linhaCima := "/" + strings.Repeat(" ", tamanhoBalao+2) + "\\"
	linhaBaixo := "\\" + strings.Repeat(" ", tamanhoBalao+2) + "/"
	linhaBase := " " + strings.Repeat("-", tamanhoBalao+2)

	// Imprime o balão de diálogo com as linhas da citação
	fmt.Println(linhaTopo)
	fmt.Println(linhaCima)
	for _, linha := range linhasCitacao {
		fmt.Printf("| %-*s |\n", tamanhoBalao, linha)
	}
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
