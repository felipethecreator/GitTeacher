## <img src="https://skillicons.dev/icons?i=git" /> Comandos Git 

# Respondendo as perguntas sobre os comandos git.

Primeiramente, qual é a diferença entre git clone & git init?

Basicamente, o comando "git clone" clona o repositório para sua máquina.
Já o "git init" cria um novo repositório git na pasta selecionada, assim permitindo que o usuário use comandos git.

# Exemplos com os 2:

```sh
git clone [url]
|||
git init 
```

## GIT REMOVE ADD ORIGIN

Para que serve? 

Esse comando permite criar, ver e excluir conexões com outros repositórios

# Exemplo:

```sh
git remote add origin https://gitlab.com/grupo/myproject.git
```

## Explicando a função de alguns comandos GIT

# a. git branch: 

Usado para listar, criar e gerenciar branches (ramificações) em um repositório Git.

Exemplo:

```sh
git branch [nome-da-branch]
```
# b. git branch -d:

Usado para deletar uma branch localmente no seu repositório.

Exemplo:

```sh
git branch -d [nome-da-branch]
```

# c. git checkout

Usado para alternar entre branches no repositório Git.

Exemplo:

```sh
git checkout [nome-da-branch]
```

# d. git branch -a

Esse comando lista todos os branches que estão disponíveis.

Exemplo:

```sh
git branch -a
```

# e. git checkout -b

Esse comando cria um novo branch com um nome específico de sua escolha.

Exemplo:

```sh
git checkout -b [nome-da-branch]
```

## Comando "git pull" e as suas variações

Para que serve o git pull?

Basicamente, sua função é, trazer as mudanças do repositório git para a sua máquina, ao dar o git pull, irá literalmente puxar os arquivos que estão na núvem para sua máquina, assim, sendo possível fazer as devidas alterações.

Exemplos:

```sh
git pull <repositório-remoto> <branch-remota>
```
