package main

type Name string

const (
	GITSTATUS   Name = "git status"
	GITCOMMIT        = "git commit"
	GITCOMMITM	 = "git commit -m"
	GITINIT          = "git init"
	GITCLONE         = "git clone"
	GITPUSH          = "git push"
	GITPULL          = "git pull"
	GITADD           = "git add"
	GITCHECKOUT      = "git checkout"
)

var CommandMap = map[Name]Command{
	GITSTATUS:   Command{"Comando utilizado para verificar as mudanças nos arquivos de um projeto."},
	GITCOMMIT:   Command{"Comando utilizado para levar as mudanças de um ambiente local para um repositório git, com a opção de inserir uma mensagem descritiva (-m [mensagem])."},
	GITCOMMITM:  Command{"Comando utilizado para salvar as alterações feitas em um repositório Git, o -m permite adicionar uma mensagem de commit junto ao comando, descrevendo assim as alterações feitas."},
	GITINIT:     Command{"Comando utilizado para inicializar um repositório local."},
	GITCLONE:    Command{"Comando utilizado para clonar um repositório remoto git para sua máquina."},
	GITPUSH:     Command{"Comando utilizado para de fato enviar as mudanças locais para o repositório git."},
	GITPULL:     Command{"Comando utilizado para trazer as mudanças do repositório git para seu repositório local."},
	GITADD:      Command{"Comando utilizado para preparar arquivos para o commit."},
	GITCHECKOUT: Command{"Comando utilizado para mudar para uma branch específica."},
}
