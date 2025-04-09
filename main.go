package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type Result struct {
	status int
}

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 1, "Número total de requisições")
	concurrency := flag.Int("concurrency", 1, "Número de requisições simultâneas")
	flag.Parse()

	if *url == "" {
		fmt.Println("Erro: parâmetro --url é obrigatório")
		os.Exit(1)
	}

	start := time.Now()

	var wg sync.WaitGroup
	results := make(chan Result, *requests)

	sem := make(chan struct{}, *concurrency)

	for i := 0; i < *requests; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()

			resp, err := http.Get(*url)
			if err != nil {
				results <- Result{status: 0}
				return
			}
			defer resp.Body.Close()

			results <- Result{status: resp.StatusCode}
		}()
	}

	wg.Wait()
	close(results)
	elapsed := time.Since(start)

	total := 0
	success := 0
	statusMap := make(map[int]int)

	for r := range results {
		total++
		if r.status == 200 {
			success++
		} else {
			statusMap[r.status]++
		}
	}

	fmt.Println("_________________________")
	fmt.Printf("Tempo total: %s\n", elapsed)
	fmt.Printf("Total de requisições: %d\n", total)
	fmt.Printf("Respostas 200 OK: %d\n", success)
	for code, count := range statusMap {
		if code != 0 {
			fmt.Printf("Status %d: %d\n", code, count)
		} else {
			fmt.Printf("Erros (sem resposta): %d\n", count)
		}
	}
}
