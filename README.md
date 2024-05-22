# Studio Sol - Prova prática Back-end (05-2024)

## Tabela de conteúdos

- [Introdução](#introdução)
  - [Pontuações possíveis](#pontuações-possíveis)
- [Stack](#stack)
- [Como executar a aplicação](#como-executar-a-aplicação)
  - [Localmente](#localmente)
  - [Dockerfile](#dockerfile)
- [Como executar os testes](#como-executar-os-testes)

## Introdução

Esse projeto é a minha prova prática para o processo seletivo de Back-end da Studio Sol em maio de 2024.

O problema consiste em avaliar quantas combinações de pontuações são possíveis para se obter um determinado placar em uma partida de futebol americano.

### Pontuações possíveis

- Touchdown (6 pontos)
  - Extra touchdown (até 2 pontos obtidos **apenas** após marcar um touchdown)
- Field goal (3 pontos)

## Stack

- Linguagem back-end: Golang
- Biblioteca para criação de servidor GraphQL: [gqlgen](https://github.com/99designs/gqlgen)
- Bibliotecas para testes: [ginkgo](https://github.com/onsi/ginkgo) e [gomega](https://github.com/onsi/gomega)
  > Essas bibliotecas usam sintaxe inspirada em BDD, o que ajuda na clareza e organização dos testes.
- Ferramenta para pipeline: Github Actions
  > O projeto tem um [pipeline do Github Actions](.github/workflows/main.yml) que faz o build do projeto e executa os testes usando _ginkgo_.
- Ferramenta de containerização: Docker

## Como executar a aplicação

### Localmente

Para executar localmente você precisa ter `go 1.20` instalado.

```bash
# Navegue até pasta raiz do projeto
cd <pasta-onde-está-o-projeto>/studio-sol-backend-test-2

# Execute a aplicação
go run cmd/server/server.go
```

A aplicação decide a porta pela variável de ambiente `PORT`. Caso nenhuma seja fornecida, a porta padrão é a 8080. As rotas são as seguintes:

Endpoint graphql: http://localhost:8080/graphql

Playground GraphQL: http://localhost:8080

> Lembre-se de trocar a porta se tiver fornecido um valor para a variável de ambiente `PORT`.

### Dockerfile

Caso não tenha `go 1.20` instalado, pode ser mais simples utilizar um container para testar.

```bash
# Navegue até pasta raiz do projeto
cd <pasta-onde-está-o-projeto>/studio-sol-backend-test-2

# Faça o build da imagem
docker build -t studio-sol-backend-test-2 .

# Rode o container
docker run -d -p 8080:8080 --name studio-sol-backend-test-2 studio-sol-backend-test-2
```

As rotas são as seguintes:

Endpoint graphql: http://localhost:8080/graphql

Playground GraphQL: http://localhost:8080

Você também pode passar outra porta ser usada pela aplicação _dentro do container_:

```bash
docker run -d -e PORT=8000 -p 8080:8000 --name studio-sol-backend-test-2 studio-sol-backend-test-2
```

## Como executar os testes

Esse projeto usa o [ginkgo](https://github.com/onsi/ginkgo) como ferramenta de testes. Então, para executar os testes, você precisa da [Ginkgo CLI](https://onsi.github.io/ginkgo/#installing-ginkgo) instalada em sua máquina.
