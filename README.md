- Estrutura básica: 
    - package main.
    - func main: é aqui que tudo começa, é aqui que tudo acaba.
    - import.
- Packages:
    - Pacotes são coleções de funções pré-prontas (ou não) que você pode utilizar.
    - Notação: pacote.Identificador. Exemplo: fmt.Println()
    - Documentação: fmt.Println.
- Variáveis: "uma variável é um objeto (uma posição na memória) capaz de reter e representar um valor ou expressão."
- Variáveis não utilizadas? Não pode: _ nelas.
- ...funções variádicas.
- Lição principal: package main, func main, pacote.Identificador.

====================================================================

- := parece uma marmota (gopher) ou o punisher.
- Uso:
    - Tipagem automática
    - Só pode repetir se houverem variáveis novas 
    - != do assignment operator (operador de atribuição)
    - Só funciona dentro de codeblocks
- Terminologia:
    - keywords (palavras-chave) são termos reservados
    - operadores, operandos
    - statement (declaração, afirmação) → uma linha de código, uma instrução que forma uma ação, formada de expressões 
    - expressão → qualquer coisa que "produz um resultado"
    - scope (abrangência)
        - package-level scope
- Lição principal:
    - := utilizado pra criar novas variáveis, dentro de code blocks
    - = para atribuir valores a variáveis já existentes

==============================================================================================

- Variável declarada em um code block é undefined em outro
- Para variáveis com uma abrangência maior, package level scope, utilizamos `var`
- Funciona em qualquer lugar

==============================================================================================

- Tipos em Go são extremamente importantes. (Veremos mais quando chegarmos em métodos e interfaces.)
- Tipos em Go são estáticos.
- Ao declarar uma variável para conter valores de um certo tipo, essa variável só poderá conter valores desse tipo.
- O tipo pode ser deduzido pelo compilador:
    - x := 10
    - var y = "a tia do batima"
- Ou pode ser declarado especificamente:
    - var w string = "isso é uma string"
    - var z int = 15
    - na declaração var z int com package scope, atribuição z = 15 no codeblock (somente)
- Tipos de dados primitivos: disponíveis na linguagem nativamente como blocos básicos de construção
    - int, string, bool
- Tipos de dados compostos: são tipos compostos de tipos primitivos, e criados pelo usuário
    - slice, array, struct, map
- O ato de definir, criar, estruturar tipos compostos chama-se composição. Veremos muito disso futuramente.