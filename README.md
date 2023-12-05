# Informações do desafio

## Objetivo

Criar um programa em Go que lê informações de um arquivo. Realizar ordenação dos dados lidos por diferentes critérios (por exemplo, ordenação por nome e por idade). Salvar o resultado ordenado em outro arquivo.

## Requisitos Técnicos

Use a biblioteca padrão do Go para manipulação de arquivos. Considere trabalhar com dados estruturados, como linhas de um arquivo CSV.

## Estrutura do Programa

Crie um arquivo de entrada contendo dados (por exemplo, CSV, cada linha representando uma entidade com algumas informações).
Implemente funções que leem o arquivo, realizam a ordenação dos dados por nome e por idade e retornam os resultados.
Salve os resultados ordenados em dois novos arquivos (um para cada critério de ordenação).

## Exemplo de Saída: Arquivo de Entrada (dados.csv)

- Nome,Idade,Pontuação
- Joao,25,80
- Maria,30,95
- Carlos,22,75

## Processamento e Arquivo de Saída (ordenado_por_nome.csv)

- Nome,Idade,Pontuação
- Carlos,22,75
- Joao,25,80
- Maria,30,95

## Processamento e Arquivo de Saída (ordenado_por_idade.csv)

- Nome,Idade,Pontuação
- Carlos,22,75
- Joao,25,80
- Maria,30,95

## Observações

- Utilize a lógica de manipulação de arquivos em Go.
- Implemente a ordenação por nome e por idade.
- Se desejar, crie funções separadas para leitura, processamento e escrita de arquivos.
- Lide com erros de maneira apropriada.

Precisamos rodar o comando go run main.go ./arquivo-origem.csv ./arquivo-destino.csv. Será usado um arquivo CSV aleatório nosso para testar sua aplicação.
