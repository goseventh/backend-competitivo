# :heart_on_fire: GoSeventh - Competição de Backend 

## :thinking: O que é isso?
Este repositório contem o código-fonte do servidor http para a [competição de backends](https://github.com/zanfranceschi/rinha-de-backend-2023-q3/blob/main/INSTRUCOES.md)

## :star2: Recursos e Conteúdo
Neste repositório, você encontrará o código-fonte completo da api rest

# :books: Arquitetura do proejto:

:unlock: **cmd**: Esta pasta contém os arquivos principais (main) em linguagem Go. Cada arquivo main pode representar um componente independente do servidor ou uma funcionalidade específica. Geralmente, esses arquivos são os pontos de entrada da aplicação, iniciando a execução do servidor e gerenciando configurações.

:lock: **internal**: Aqui, estão localizados os pacotes internos do projeto. Esses pacotes contêm a lógica de negócios e a implementação central do servidor GoSeventh. Eles são projetados para serem utilizados internamente dentro do projeto e podem ser importados por outros componentes do servidor, como os arquivos main na pasta "cmd".

:loudspeaker: **pkg**: Nesta pasta, você encontrará pacotes públicos. Esses pacotes são bibliotecas ou componentes reutilizáveis que podem ser compartilhados com outros projetos ou utilizados por desenvolvedores externos. Eles oferecem funcionalidades complexas que podem ser facilmente integradas em outros sistemas.egócio 
