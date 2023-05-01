Instalação do Go
Acesse o site oficial do Go em https://golang.org/dl/
Faça o download do instalador do Go para o seu sistema operacional
Siga as instruções do instalador para concluir a instalação do Go
Instalação do MariaDB
Acesse o site oficial do MariaDB em https://mariadb.org/download/
Selecione a versão do MariaDB que você deseja baixar e faça o download para o seu sistema operacional
Siga as instruções do instalador para concluir a instalação do MariaDB
Caso você esteja utilizando um sistema operacional Linux, pode seguir as instruções abaixo para instalar o Go e o MariaDB através do terminal:

Instalação do Go no Linux
Abra o terminal e execute o seguinte comando para baixar o pacote do Go:
wget https://golang.org/dl/go1.X.X.linux-amd64.tar.gz
Substitua "X.X" pela versão mais recente do Go.

Descompacte o pacote do Go com o seguinte comando:
sudo tar -C /usr/local -xzf go1.X.X.linux-amd64.tar.gz
Substitua "X.X" pela versão que você baixou.

Adicione o diretório bin do Go ao PATH do sistema:
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
source ~/.profile
Isso permitirá que você execute comandos do Go a partir do terminal.

Instalação do MariaDB no Linux
Abra o terminal e execute o seguinte comando para baixar e instalar o MariaDB:
sudo apt-get update
sudo apt-get install mariadb-server
Siga as instruções do instalador para concluir a instalação do MariaDB.

Com essas instruções, você deve ser capaz de instalar facilmente o Go e o MariaDB em seu sistema operacional.

Se atente ao instalar o mariaDB, você vai precisar passar a senha para o arquivo ConfigDB.toml

Por fim crie no mariaDB um banco chamado "spotless" - create database spotless;
Para rodar a api digite no terminal: go run .\main.go 