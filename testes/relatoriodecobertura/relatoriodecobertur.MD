
comando para exibir cobertur de testes 
λ go test -cover
PASS
coverage: 100.0% of statements
ok      github.com/albertoribeiro/matematica    0.164s



comando para imprimir cobertur de testes em um arquivo
λ go test -coverprofile=resultado.out
PASS
coverage: 100.0% of statements
ok      github.com/albertoribeiro/matematica    0.181s

Mostra o conteudo do relatorio de cobertura de testes
C:\Users\alberto\go\src\github.com\albertoribeiro\matematica  
λ go tool cover -func=resultado.out
github.com/albertoribeiro/matematica/matematica.go:12:  Media           100.0%
total:                                                  (statements)    100.0%

mostra um html com o conteudo do relatorio de cobertura de testes
λ go tool cover -html=resultado.out