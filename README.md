# API DEV BOOK

O Dev Book é um projeto desenvolvido para o curso de Go do professor Otávio Augusto Gallego. 
O projeto consiste em uma API para uma rede social, chamada Dev Book.

## Funcionalidades

A API possui as seguintes funcionalidades:

### Segurança
- /login
- /update-password

### Usuários
- Buscar usuários por nome
- Buscar usuário por id
- Atualizar usuário
- Deletar usuário
- Atualizar 
- Seguir e deixar de seguir 
- Listar seguidores
- Listar seguindo

### Posts
- Criar
- Atualizar
- Deletar
- Curtir
- Descurtir
- Buscar posts por usuário
- Buscar post por Id
- Buscar todos os posts


## Tecnologias

- Go
- Docker
- Docker Compose
- Mysql

## Como executar o projeto

1. Clone o repositório
2. Preencha as variáveis de ambiente conforme o arquivo `.env.example`
3. Execute o comando `docker-compose up` na raiz do projeto
4. Acesse a API
