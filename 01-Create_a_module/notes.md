# Anotações

The go mod init command creates a go.mod file to track your code's dependencies. So far, the file includes only the name of your module and the Go version your code supports. But as you add dependencies, the go.mod file will list the versions your code depends on. This keeps builds reproducible and gives you direct control over which module versions to use.

O comando ```go mod init``` cria o go.mod e rastreia as dependências. A medida que você adiciona dependências, o go.mod vai listar as versões que seu código depende. Isso mantém a reprodutilidade e entrega controle direto sobre qual versão do modulo utilizar.