

## Setup

```bash
mkdir hello
cd hello

#enable dependency tracking
go mod init github.com/hello
```

## Simple code
### Writing very basic code
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Note: a package is a way to group functions, and it's made up of all the files in the same directory


### Run the code
```bash
go run
# or
go run main.go
```

## Calling external code

Visit https://pkg.go.dev and check for new modules and versions. For now, import rsc.io/quote package 

```go
package main

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
```

Adding new module requirements and sums, Go will add the quote module as a requirement, as well as a go.sum file for use in authenticating the module

```bash
go mod tidy
```

Then,
```bash
go run
# or
go run main.go
```

Adding an specific dependency is simple as:

```bash
# Add all dependencies for a package in te module
go get . 

# or add an specific dependency
go get example.com/newmodule
```



# Functions

```go
func main() {
  resultado, str := soma(10, 20)
  resultado, str := somaTudo(10, 20, 23 ,15 ,17 ,89 ,31)
}

// É necessário declarar o retorno da função para que ela retorne algo.
func soma(a int, b int) (int, string) {
  return a + b, "soma"
}
// or
func soma(a int, b int) (result int, str string) {
    result = a + b
    str = "string"
    return
}

// ...x is an array f integres
func somaTudo(x ...int) int {
  resultado := 0

  for _, v := range x {
      resultado += v
  }

  return resultado
}
```

Funções anonimas e funçoes dentro de funções para se utilizar:

```go

func main() {

  resultado:= func(x ...int) func() int {

    res := 0

    for _, v := range x {
      res += v
    }

    return func() int {
      return res * res
    }
  }
  // Resposta 0x10a6d40. Isso porque estamos aqui retornando uma função, mas deve ser retornado o valor inteiro 
  // fmt.Println(resultado(54,54,54,54)) 

  // Por isso, devemos retornar o resultado e a função:e a chamada da função que deve retornar o inteiro. Resposta 46656
  fmt.Println(resultado(54,54,54,54)()) 
}
```

Transfromando a função em um método (orientado a objetos in the Go way).
Os métodos do Go são basedos em structs utilizando o type

```go
type Carro struct {
  Nome string
}

// Para declararmos que a função abaixo é um método da struct Carro, é necessário passar c Carro e o nome dela vai ser andar, por exemplo.

// Só de colocar (c Carro), esse cara é o relacionamento da função abaixo com o struct acima. Isto é, um BIND! andar() é o nome do método.

func (c Carro) andar() {
  fmt.Println(c.Nome, "andou")
}

// LOgo, esse método pode ser chamado dentro da função principal utilizando o tipo Carro

func main() {

  carro := Carro {
    Nome: "BMW",
  }

  carro.andar()
  // Resultado: BMW andou
}
```


Trabalhando com Ponteiros

```go
package main

import "fmt"

func main() {

  // Quando o valor 50 é atribuido para a variávek a, a variávek a é apontada para um lugar da memória e esse lugar da memoria vai guardar o valor 50.

  // memoria-0xc00004a770(50) <------- a <--------- 50

  // Caso o valor da variável a mude para 40, é necessário mudar o valor do lugar (endereço) da memória que armazenava o valor 50.

  // memoria-0xc00004a770(40) <------- a <--------- 40

  // LOgo, toda a vez que chamarmos a variável a,, a variável a tem um endereço e temos que mudar o valor nesse endereço.

  a := 10
  fmt.println(&a) // 0xc00004a770

  // memoria-0xc00004a770(10) <------- a <--------- 10

  // O endereço da memória fica catalogado nos nossos queridos ponteiros. 

  var ponteito *int = &a
  // Isso significa que o apontamento que está guardando o valor do a para a memória está agora guardado nesse ponteiro. Isto é, a variável ponteiro sabe agora onde a variável a está guardandoos valores dele. O ponteiro detém o endereço da memória onde está guardado o valor de a

  // Uma vez que tenhamos o endereço da memória, é possivel fazer o reference! O reference pega o endereço da memória, consulta o endereço que está lá e traz o resultado para a gente.

  fmt.Println(ponteiro) // 0xc00004a770
  fmt.Println(*ponteiro) // 10

  // Pode-se notar que o asterisco na frente do poneteiro (enedereço) traz o valor que está armazenado no endereço da memória.

  *ponteiro = 50
  fmt.Println(*ponteiro) // 50
  fmt.Println(a) // 50

  // Isso aconteceu porque o valor 50 voi atribuido como valor do enereço da mamória quem que a variál ponteiro está apontando. O valor armazenado muda mas não muda o valor da memória.

  // Então, a variável a aponta para o mesmo endereço que o *ponteiro. Logo todo o valor que for mudar pelo ponteiro, a variável a tb sofrerá as alterações.

  b := &a
  fmt.Println(*b) // 50
  *b = 60
  fmt.Println(*ponteiro) // 60
  fmt.Println(a) // 60

  // Isso acontece pq a, *b e o *ponteiro estão apontando para o mesmo endereço!

  func abc(a *int) int {
    
  }

}

```

