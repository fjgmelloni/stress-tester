# Stress Tester CLI - Desafio Pós FullCycle

Este projeto foi desenvolvido como parte do **desafio da Pós-Graduação FullCycle**, com o objetivo de criar uma aplicação **CLI em Go** capaz de realizar **testes de carga (stress test)** em serviços web.



Por se tratar de uma aplicação extremamente simples, com uma única responsabilidade e sem dependências externas ou lógica de domínio complexa, **optei por não utilizar DDD nem Clean Architecture** neste projeto.  
O foco foi em manter o código simples, direto e fácil de entender, utilizando apenas pacotes padrão da linguagem Go.

---

## Objetivo do desafio

Criar uma aplicação capaz de:

- Receber via linha de comando:
  - `--url`: URL do serviço a ser testado
  - `--requests`: número total de requisições
  - `--concurrency`: número de requisições simultâneas
- Realizar as requisições com concorrência controlada
- Gerar um **relatório de desempenho** ao final com:
  - Tempo total da execução
  - Total de requisições realizadas
  - Quantidade de respostas `200 OK`
  - Distribuição dos demais códigos HTTP (ex: 404, 500 etc)

---


## Executando com Docker

Para rodar a aplicação utilizando Docker, basta seguir os passos abaixo:

### 1. Build da imagem

Dentro da raiz do projeto:

```bash
docker build -t stress-tester .
```

### 2. Execução do teste

```bash
docker run stress-tester ./stress-tester --url=https://www.google.com --requests=100 --concurrency=10
```

> **Importante:**  
> O comando acima precisa incluir `./stress-tester` explicitamente, pois o Docker substitui o `CMD` padrão quando argumentos são passados.  
> Isso é uma peculiaridade que deve ser observada na execução.

---

## Observações

- Usei `https://www.google.com` apenas como exemplo.
- O número de requisições pode ser alterado:
  - **100** foi utilizado no exemplo por ser mais rápido para testes locais.
  - O desafio mencionava **1000**, que também pode ser executado, porém leva mais tempo.

Exemplo com 1000 requisições:

```bash
docker run stress-tester ./stress-tester --url=https://www.google.com --requests=1000 --concurrency=10
```

---

## Autor

**Felicio Melloni**  
