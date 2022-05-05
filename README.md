

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

// Para declararmos que a função andar() abaixo é um método de Carro, é necessário passar c Carro e o nome dela vai ser andar, por exemplo.

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
}
  

```

Outro exemplo 
```go

func main () {

  variavel := 10

  abc(&variavel) // tem que se colocar o endereço da memoria pois a função pede um parâmetro de ponteiro

  fmt.Ptintln(variavel) // 200
  
  func abc(a *int) {
      *a = 200
    }
}
```

Criamos uma função que nao esta retornando nada, mas ela está alterando um valor que foi passado por referencia atraves do ponteiro! Pois o valor que foi passado por parâmetro é um endereçamento de memória. Portanto, uma função que lida com ponteiro também nao precisa ter o return!


Outro exemplo com Struct
```go

type Carro struct {
  Name string
}

// a variavel c consegue acessar e instanciar o valor da struct
func (c Carro) andou() {
  fmt.Println(c.Name, "andou")
}

func main () {

  carro := Carro {
    Name: "Ka"
  }

  carro.andou() // Ka andou

}
```

Outro exemplo com Struct
```go

type Carro struct {
  Name string
}

// a variavel c consegue acessar e instanciar o valor da struct
func (c Carro) andou() {
  c.Name = "BMW"
  fmt.Println(c.Name, "andou")
}

func main () {

  carro := Carro {
    Name: "Ka"
  }

  carro.andou() // BMW andou
  fmt.Println(carro.Name) // Ka

}
```


Outro exemplo com Struct
```go

type Carro struct {
  Name string
}

// a variavel c consegue acessar e instanciar o valor da struct
func (c Carro) andou() {
  c.Name = "BMW"
  fmt.Println(c.Name, "andou")
}

func main () {

  carro := Carro {
    Name: "Ka"
  }

  carro.andou() // BMW andou
  fmt.Println(carro.Name) // Ka 

  // A função Println apenas retorna o valordas variáveis dentro do escopo da função.

}
```

Outro exemplo com Struct
```go

type Carro struct {
  Name string
}

// a variavel c consegue acessar e instanciar o valor da struct

// Agora, quando dermos c.Name = "BMW", o valor sera alterado na referencia! Isso vai mudar para o Name que está na Struct carro do tipo Carro porque acessa o mesmo local da memória.
func (c *Carro) andou() {
  c.Name = "BMW"
  fmt.Println(c.Name, "andou")
}

func main () {

  carro := Carro {
    Name: "Ka"
  }

  carro.andou() // BMW andou
  fmt.Println(carro.Name) // BMW 

  // A função Println apenas retorna o valor das variáveis dentro do escopo da função.

}
```

Então, quando criamos um método (struct) e SE vc quiser alterar o valor no endereço que estão rodando nessa instância, tem que colocar o *Carro para alterar.




## Structs
Structs nao podem ser confundido com classes.

Go tem o GoWay de programar 

Quando criamos diversos types, os dados ficam soltos. Para ajudar a estruturar essas informações, existe o struct, que cria o esqueleto da informação, 

Arquivo struct.go

```go
package main

import fmt

type Cliente struct {
  Nome string
  Email string
  CPF int
}

// Método da struct Cliente em que c é do tipo Cliente.
func (c Cliente)Exibe() {
  // a variável c é o que dá acesso à struct
  fmt.Println("Exibindo cliente pelo método: ", c.Nome)
}

// outro tipo de cliente muito semelhante ao anterior. Quando estamos utilizando a mesma estrutura anterior, podemos fazer uma COMPOIÇÃO DE STRUCT
// type ClienteInternacional struct {
//   Nome string
//   Email string
//   CPF int
//   Pais string
// }

// Logo:
type ClienteInternacional struct {
    Cliente
    Pais string `json:"pais"`
}

func main() {
  cliente := Cliente{
    Nome: "Rogerio",
    Email: "r@r.com",
    CPF: 123345,
  }

  fmt.Println(cliente) //{Rogerio r@r.com 123345}

  cliente2 := Cliente{"Maricota", "cota@cota.com", 987456}
  fmt.Printf("Nome: %s. Email: %s, CPF: %d", cliente2.Nome, cliente2.Email, cliente2.CPF)
  //Maricota cota@cota.com 987456%}

  cliente3 := ClienteInternacional{
    Cliente: Cliente{
      Nome: "Davi"
      Email: "d@d",
      CPF: 465454
    },
    Pais: "EUA",
  }

  fmt.Printf("Nome: %s. Email: %s, CPF: %d", cliente3.Nome, cliente3.Email, cliente3.Pais)

  // MUITO IMPORTANTE: Mesmo fazendo uma composição de uma struct com outra, na utilização das informações, nao precisamos declarar cliente3.Cliente.Nome e isso acaba funcionando como uma espécie de herança como se Nome fizesse parte diretamente do ClienteInternacional e não no ClienteInternacional.Cliente.Nome.

  // Em Go, falando em métodos,podemos atachar (bind) uma função em uma struct. Logo, em Golang, uma struct passa a ter um método. O comum é ouvir dizer que uma classe tem um método. 

  //No caso deste exemplo, a função Exibe vai estar atrelada à struct Cliente. Portanto, Exibe() é um método da struct Cliente.
  cliente.Exibe()
  cliente2.Exibe()
  cliente3.Exibe()
  //Aqui conseguimos chamar o cliente3 como se o método Exibe() estivesse atrelado ao ClienteInternacional, sendo que ele está atrelado apenas ao Cliente. Por isso que se parece muito com uma herança, apesar de estarmos compondo uma struct dentro da outra.

  // JSON: pegamos a informação que está em uma classe e transformamos isso então em JSON para trabalharmos com APIs. Por conta disso, podemos pegar uma struct e converter para JSON.

  // Marshal é quando pegamos algo que está em uma struct e transformamos para JSON.
  clienteJson, err := json.Marshal(cliente3)
  if err != nil {
    log.Fatal(err.Error())
  }

  fmt.Println(clienteJson) // [123 34 78 111 ... 34 125]

  // Toda a vez que damos o Marshal, o Marshal pega esses dados em um formato de Slice de byte. entao cada numero presentado acima é um byte. O que devemos fazer é pegar os bytes e converter no formato de string conforme demonstrado abaixo,
  fmt.Println(string(clienteJson)) 
  // Agora sim temos o Json gerado a partir da struct:
  // {"Nome": "Davi", "Email": "d@d.com", "CPF": 465454, "Pais": "EUA"}

  //MUITO IMPORTANTE:  Para trabalharmos com o nome dos campos em minusculo, pois os oscampos estão sendo gerados a partir da struct, podemos trabalhar com as Tags!

  // Essas tags, conseguimos associar em uma struct e toda a vez que rodamos alguma função, essa funçao tem acesso a essas Tags, ela ve se alguma dessas Tags interessa para ela e se interessar ela consegue transformar em tempo real o valor que está naquela Tag e rodar alguma regra que ela queira. Para json, existe a mais famosa que é `json: pais`. Dessa maneira, essa Tag (anotação) fala que quando rodar a funçao de json, a função vai transformar a string "Pais" para "pais"

  // Agora, tb podemos querer fazer o contrário: pegar o Json e hidratar a struct. Ter uma struct vazia e preencher essa struct com os valores do Json. 
  jsoncliente4 := `{"Nome": "Davi", "Email": "d@d.com", "CPF": 465454, "Pais": "EUA"}`
  cliente4 := ClienteInternacional{}

  // Unmarshal pede para passarmos um slice de bytes json.Unmarshal([]byte, v interface{})
  // []byte(jsoncliente4) -> transforma de string para bytes
  // json.Unmarshal([]byte(jsoncliente4), cliente4) // Dessa maneira, o valor de cliente4  iria ser alterado apenas neste escopo, logo, precisamos passar esse valor por referencia.
  
  json.Unmarshal([]byte(jsoncliente4), &cliente4)
  // Dessa forma isso irá alterar o mesmo endereço da memoria que o cliente4 está colocando. Isso acontece pq a função Unmarshal está esperando receber um enedereço, e dentro dela deve ter algo do tipo *ponteiro = [dado em byte] para preencher a struct cliente4.

  fmt.Println(cliente4) // {{Davu d@d 465454 EUA}}
}
```

## Go Routines, paralelism, multithread, channels

Concorrencia e Paralelismo

Concorrencia é conseguirmos LIDAR com diversas coisas ao mesmo tempo 
Paralelismo é conseguir FAZER diversas coisas ao mesmo tempo

Cada programa iniciado gera um Processo no OS que é independente de outro Processo. Cada processo nao compartilha memoria entre si. Por padrao, toda vez que se inicia um processo, é iniciado uma tarefa que chamamos de Thread (fio de execução). 

Quando estamos trabalhando com Threads, o espaço de memória é compartilhado, independentemente se for concorrente ou paralelo.

Um ponto importante é notar e saber que quando se está trabalhando com Threads, o valor armazenado em um espaço da memória pode ser modificado pela outra Thread se nao prestarmos atenção.

O Golang trabalha com channels!

Toda a vez que o programa começc a trabalhar, o sistema inicia um Thread, e é alocado 1mb de memoria por Thread. Cada vez que acontece uma nova Thread, acontece algo que chama syscall, isto é, o programa chama o OS para que o OS gere uma nova Thread. 

Quando trabalhamos com MultiThreads, existe um cara que chama Scheduler que define quando cada Thread deve ser executada. 

Scheduler Pre emptivo -> Coloca um tempo específico e aquele processo/Thead que está sendo executada é interrompido independente do que ela esteja fazendo para que a outra comece a ser executada. O custo é que cada vez que se muda da Thread A para a B, existe algo que se chama contexto. Entao muda-se o contexto para executar uma nova Thread.Muda todo o set up para rodar o ontexto e as coisas de cada Thread para que cada uma possa ser executada com o seu próprio contexto.

Scheduler cooperativa -> Pede para a Thread trabalhar e a outra esperar. Só depois que terminar a tarefa, virá a outra Thread/Processo para que possa ser executado. Porém existem algumas condições. No exemplo da pizza, espera-se a massa crescer e bloquia-se o tempo esperando algo acontecer. Nesse momento, o scheduler agenda para que a outra trabalhe enquanto uma espera.

A disvantagem do cooperatio é que uma thead pode monopolizar o posto e fazer com que outras tb com muita prioridade tenham que esperar.


O Golang trabalha, por padrao, de forma cooperativa e um pouco diferente com as Threads. Pq as Threads dos exemplos acima sao gerados pelo OS. Já no Go, trabalha-se com runtimes. Runtimes é como uma biblioteca que tem todo o código do Go ali. Nesse runtime, acontece o seguinte: o Go junta o runtime de si com o código escrito e formasse um binário. 

O grande ponto, é que ao invés de o Go chamar o OS, ele tem o próprio Scheduler e cria as proprias Threads. Entao o Go trabalha com Threads mas quem gerencia essas Threads é o próprio Runtime do Go. Essas Threads sao chamadas de Threads in user land or GreenThreads. O importante é entendermos que esse tipo de Thread é gerenciado pelo runtime. E nesse caso, nao eh realizada uma chamada para o OS. Melhor. Outra vantagem é que uma Thead no Go ocupa muito menos do que 1mb.

As threads no Go nao ocupa o OS e elas ocupam 2kb de memória!

Entao consegumos trabalhar com milhares  milhares dessas Threads. E raramente vamos querer trabalhar com multThreads pois o Go funciona com channels.

Apesar de o Go ter o carater de Scheduler cooperativo. Para que nao haja monopolio de Threads, existem algumas regras: 
- Se tiver que acessar disco e fazer qualquer operação bloqueante ou nao bloqueante, vai pra próxima,
- Se tiver que fazer uma chamada externa, uma API ou qualquer coisa desse tipo que bloquei, vai pra próxima.
- Se tiver que dar um timeout ou sleep, vai para a proxima.

Embora isso já aconteça, o Go tb consegue trabalhar de forma pre emptiva, Threads por tempo. 

Exemplo de Go Routine:

```go
package main

import (
  fmt
  time
)

func contador(tipo string) {
  for i := 0; i < 5; i++ {
    fmt.Println(tipo, i)
  }
}

func main() {
  contador(tipo: "sem go routine")
  go contador(tipo: "com go routine")

  fmt.Println("Hello 1")
  fmt.Println("Hello 2")
}

// sem go routine 0
// sem go routine 1
// sem go routine 2
// sem go routine 3
// sem go routine 4
// Hello 1
// Hello 2
```

Quando colocamos go contador(), o comando go indica que deve ser iniciada uma Thread. Mas como o programa é muito rápido, nao deu tempo de iniciar o Print!

Colocando o time.Sleep(time.Second), espera-se m segundo para finalizar o programa

```go
package main

import (
  fmt
  time
)

func contador(tipo string) {
  for i := 0; i < 5; i++ {
    fmt.Println(tipo, i)
  }
}

func main() {
  contador(tipo: "sem go routine")
  go contador(tipo: "com go routine")

  fmt.Println("Hello 1")
  fmt.Println("Hello 2")

  time.Seleep(time.Second)
  fmt.Println("fim")

}

// sem go routine 0
// sem go routine 1
// sem go routine 2
// sem go routine 3
// sem go routine 4
// Hello 1
// Hello 2
// com go routine 0
// com go routine 1
// com go routine 2
// com go routine 3
// com go routine 4
// fim
```

Dessa forma podemos ver que o programa rodou o sem go routine e iniciou uma nova therad com go routine de forma concorrente!

UM exemplo mais claro a seguir:

```go
package main

import (
  fmt
  time
)

func contador(tipo string) {
  for i := 0; i < 5; i++ {
    fmt.Println(tipo, i)
    time.Sleep(time.Second)
  }
}

func main() {

  go contador(tipo: "a")
  go contador(tipo: "b")

}

// a0
// b0
// a1
// b1
// a2
// b2
// a3
// b3
// a4
// b4

```

Dessa forma, ambas rodam de forma concorrente como 2 Threads dentro do runtime do Go.


Um exemplo para nao ficar tetando advinhar qual tipo de processamento de Threads Utilizar. O Go é multicore e, eventualmente, ele consegue rodar o código em parallo tb!

Para isolarmos o processamento do Go, pediremos a ele para que ele rode em apenas 1 core.

```go

package main

import "runtime"

func main() {
  runtime.GOMAXPROCS(n:1)
  fmt.Println("Começou")
  // Routine com loop infinito
  go func() {
    for {

    }
  }()
  time.Sleep(time.Second)
  fmt.Println("Terminou")
}
```

A função main, é uma Thread no Go. Mas a função go func() tb é uma outra Thread que roda um loop infinito. Logo, se ela é uma Thread de loop infinito e o Go trabalha de forma cooperativa, o que acontece é que o go func() vai travar o programa até que a função do loop infinito termine e entao seja continuada a execução da func main().

Na versao 1.13, o programa go func() não está fazendo nenhuma para na função para dar uma desculpa para que o Schedular do Go possa chamar e voltar para a Thread do main() pq essa Thread nunca vai terminar. Pq dessa forma cooperativa vai travar.

Na versão 1.17, nao travou. Isso pq a partir da versao 1.14 o Go adicionou alguns recursos no Scheduler que quando acontecem situações como essa, ele tb consiga trabalhar de forma pre emptiva. Quem está enrolando cai fora e deixa o outro terminar de rodar. Isso foi uma GRANDE MUDANÇCA!

### Channels

Os channels permitem que uma Thread se comuniquem com a outra.

No Go, ao invés de ter esse tipo de conflito, cada programa/processo/Thread nao precise compartilhar a memoria para se comunicar. ELE SE COMUNICA COMPARTILHANDO MEMÓRIA!

Exemplo:

func main é uma Thread e go func() é uma outra Thread.

```go
package main

// Thread1
func main() {

  // Thread1 <-> // Thread2
  hello := make(chan string)

  // Thread2
  go func() {
    // Isto fala que o valor que se está passando por esse canal é "Hello World"
    hello <- "Hello World"

  }()

  result := <-hello
  fmt.Println(result)


}
```

1. Cria um canal chamado hello
2. O canal hello recebeu o valor "Hello World" em uma Thread
3. Toda vez em que tiver um valor no canal hello, deixe a variável result pegar este valor



Outro exemplo com Channels em Go:

```go
package main

func main() {
  forever := make(chan string)

  go func() {
    x := true
    for {
      if x == true {
        continue
      }
    }
  }()

  fmt.Println("Aguardamdo para sempre")
  <-forever
}
```

O channel forever fica aguardando alguem colocar um valor nele. Enquanto ninguem coloca algum valor no forever ele fica aguardando mesmo que o Go possa vir a tabalhar na forma pre emptiva.

Isso significa quem enquanto o Go nao receber nenhum valor, ele vai travar a execução. 

O CHANNEL PARA A EXECUÇÃO ENQUANTO NAO CHEGA NENHUM VALOR PARA ELE!


Outro exemplo:
```go
package main

func main() {

  hello := make(chan string)

  go func() {
    time.Sleep(time.Second * 2)
    hello <- "hello world"

  }()


  select {
    // Caso venha algum valor em hello que seja atriuido pro x, imprime o valor de x
    case x := <-hello:
      fmt.Println(x)
    default:
      fmt.Println("default")
  }

}
```
O resultado do programa foi default. Pois o select recebeu a instrução default antes que o channel pudesse receber o valor de hello world.


Outro exemplo:

```go
package main

func main() {
  queue := make(chan int)

  go func() {
    i := 0
    for {
      time.Sleep(time.Second)
      queue <- i
      i++
      //S O programa só irá conseguir voltar novamente nesse loop quando o queue tiver sido esvaziado. oU SEJA, SE NINGUÉM LER ESSE QUE EM UMA OUTRA THREAD, O QUEUE <- 1 AI PARAR A EXECUÇÃO DO PROGRAMA ATÉ QUE ACONTEÇA A SUA LEITURA.
    }
  }()

  // pARA LER A QUEUR:
  for x := range queue {
    Pega o valor que está na queue, envia para x e imprime o valor na tela.
    fmt.Printls(x)
  }

  // A cada vez que a fila queue é lida, o for faz uma nova iteração.pois enquanto o channel não é esvaziado, ele fica parado aguardando a interação.
  // 0
  // 1
  // 2
  // 3
  // 4
}
```

No exemplo acima, estamos compartilhando a variável channel entre o processo principal main() e a outra Thread go func(). A cada vez que o for debaixo iterar por x, o for loop roda novamente.



Exemplo com go web server

Cada vez que o Go recebe uma nova requisição, ele gera uma nova Thread, por isso ele é tao rapido.

Exemplo:

```go

package main

// função worker conta
func worker(workerId int, msg chan int) {
  for res := range msg {
    fmt.Println("Worker", workerId, " Msg:", res)
    time.Sleep(time.Second)
  }
}

func main() {
  msg := make(chan int)

  // Go routine aqui. Enão o channel msg é iniciado e então o worker roda em background.
  go worker(workerId: 1, msg)


  for i := 0; i < 10; i++ {
    //passaro valor do i para a msg
    msg <- i
  }
}

// Toda vez que rodar um loop do for e o valor de i for escrito no chanel msg, a função worker pea esse valor e imprime a msg.
  
```

Da form demonstrada acima, é necessário que o worker processe a página para que então outro possa acessá-la.

Adicionando um novo Worker, workerId: 2, é possível 

xemplo:

```go

package main

func worker(workerId int, msg chan int) {
  for res := range msg {
    fmt.Println("Worker", workerId, " Msg:", res)
    time.Sleep(time.Second)
  }
}

func main() {
  msg := make(chan int)

  go worker(workerId: 1, msg)
  go worker(workerId: 2, msg)
  go worker(workerId: 3, msg)
  go worker(workerId: 4, msg)


  for i := 0; i < 10; i++ {
    msg <- i
  }
}
  
```
Dessa forma, enquanto um worker espera, o próximo worker pega outra mensagem!

Esse é o grande poder do go routines pq não é necessário esperar um terminar para começar outra Thread! E cada Thread ocupa apenas 2kb!

Utilizando a linguagem tradicional, cada Thread iria ocupar 1mb. 

No caso do Apache2 webserver, declaramos que o numero máximo de workers sao 5. Logo, vai acontecer de ele faer um fork ali e dividir em 5 threads para ele rodar. E confome os usuários vao acessando, ele vai lidando de 5 em 5 e vai gerando uma fila atrás dele.

No caso do GO, a gente pode colocar milhares de workers, uma vez que cada Thread do worker é muito leve para processar uma fila de páginas!


Quando estamos trabalhando com channels, conseguimos pegar os valores que estao rodando em um processo e ir lidando com esses valores em outros processos e fazer um canal que manda pro outro e que manda poro outro... 

Isso é ótimo pq vao greando várias Threads e deixando o processamento ainda mais rápido.

 Go ainda trabalha com garbage collector que fica de tempos em tempos esperando para trabalhar. O Garbage collector tb é uma go routine que fica rodando. 




## Generics
Go 1.18

Quase a mesma função, mas tem que escrever outra apenas por causa da tipagem, por exemplo.

```go
package main

import "fmt"

func SomaInteiros(m map[string]int64) int64 {
  var soma int64
  for _, v := range m {
    soma += v
  }
  return soma
}

func SomaFloat(m map[string]float64) float64 {
  var soma float64
  for _, v := range m {
    soma += v
  }
  return soma
}


func main() {
  fmt.Println(SomaInteiros(map[string]int64{"a":1, "b":2, "c":3}))
  fmt.Println(SomaFloat(map[string]float64{"a":1.1, "b":22.2, "c":3.2}))
}
```


Então:


```go
package main

import (
  fmt
  constraints  
)

type Number interface {
  ~int | ~int64 | ~float64
}

type MyNumber int

// Quando falamos de Generics, Number int64 ou float64 sao chamado sde contraints.
func SomaGenerica[T Number] (m map[string]T) T {
  var soma T
  for _, v := range m {
    soma += v
  }
  return soma
}


// o tipo comparable sugere que os numeros possam ser comparáveis em relação ao tipo. Aqui o tipo any nao funcionaria pois sugere-se que o valor recebido poderia ser uma string e isso nao está dentro do type de T
func Soma[T comparable] (number1 T, number2 T) T {
  if number1 == number2 {
    return number1
  }
  return number2
}


// COm o comparable, apenaconseguimos comparar igualdade. Neste caso existe um pacote oficial de COntraints do Go que nos auxiliam neste tipo de coisa. Esse pacote chama-se constraints.
// go get golang.org/x/exp/constraints
func Max[T constraints.Ordered] (number1 T, number2 T) T {
  if number1 > number2 {
    return number1
  }
  return number2
}

// As generics tb funcionam em métodos!
// Não conseguimos garantir, a pricipio que o tipo any tem um método .String
// Nesse momento, podemos criar um tipo stringer que tem um método que retorne uma string
type stringer inteface {
  String() string
}

//Então um outro tipo MyString do tipo string
type MyString string

// E um método atrelado ao Mystring que retorna uma string convertida em string
func (m MyString) String() string {
  return string(m)
}

func concat[T stringer] (vals []T) string {
  result := ""
  for _, val := range vals {
    result += val.String()
  }
  return result
}



func main() {

  // concat espera receber um array de strings e retorna o tipo MyString
  // Dentro da função concat espera-se receber parâmetros do tipo stringer
  // dentro do tipo stringer foi declado um método String() que está atrelado à MyString e retorna uma string
  fmt.Prinln(concat[](MyString{"a", "b", "c"}))

  var x, y, z MyNumber
  x = 1
  y = 2
  z = 3

  fmt.Println(SomaGenerica(map[string]MyNumber{"a":x, "b":y, "c":z}))
  // No caso acima, o Go impede que isso aconteça pq a Soma Gnerica possui um tipo definido como Number. 
  // Entretanto, dentro do tipo Number e MyNumber, existe o tipo int. Para podermos prosseguir com isso, podemos fazer uma aproximação. ~, Portanto, se tiver algum tipo, que o tipo dele é int, por exemplo, o Go passa a aceitar e fazer parte da interface MyNumber

  fmt.Println(SomaGenerica(map[string]int64{"a":1, "b":2, "c":3}))
  fmt.Println(SomaGenerica(map[string]float64{"a":1.1, "b":22.2, "c":3.2}))
}
```

Em geral, com Generics é possível reaproveitar mais códigos e criar coisas que antigamente eram um pouco mais dolorosas na linguagem Go.