# 🌐 IoTec Dashboard - UNIVALI

Este projeto é um sistema web de monitoramento de sensores desenvolvido para o **laboratório IoTec** da **UNIVALI - Campus Itajaí**. Ele tem como objetivo substituir o sistema anterior hospedado no diretório: `x`.

---

## 📦 Tecnologias Utilizadas

### Back-end
- [Golang](https://go.dev/) – Servidor principal, escutando tópicos MQTT e fornecendo APIs/WebSocket
- [MongoDB Atlas](https://www.mongodb.com/cloud/atlas) – Banco de dados NoSQL em nuvem
- [MQTT (Eclipse Paho)](https://github.com/eclipse/paho.mqtt.golang) – Protocolo leve de mensagens IoT
- [JWT](https://jwt.io/) – Autenticação segura de administradores
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) – Criptografia de senhas

### Front-end
- [React.js](https://reactjs.org/) – Interface dinâmica e responsiva
- [Recharts](https://recharts.org/) – Gráficos dinâmicos em tempo real
- [WebSocket](https://developer.mozilla.org/en-US/docs/Web/API/WebSocket) – Comunicação em tempo real com o back-end

---

## 🚀 Funcionalidades

### 🖥 Modo Público (Visualização)
- Exibe **cards dinâmicos** com dados de sensores
- Gráficos em tempo real para sensores com dados relevantes
- Atualização automática via WebSocket (sem necessidade de recarregar a página)

### 🔐 Modo Administrativo (ADM)
- Acesso via login e senha com autenticação segura
- Permite **criar** e **remover** cards de sensores
- Define:
  - ID do sensor
  - Tópico MQTT
  - Campos que devem ser exibidos (ex: temperatura, umidade)
  - Se o gráfico deve ser exibido ou não

---

## 🔄 Fluxo de Dados

1. Dispositivos IoT enviam mensagens no padrão: temperatura, 23, umidade, 50, pressão, 100. Campos podem conter `NULL` caso o sensor não tenha esse dado.

2. O servidor Go escuta os tópicos MQTT e:
- Salva os dados no MongoDB
- Envia os dados ao front-end via WebSocket

3. O front-end React exibe os dados conforme os cards configurados pelo administrador

---

## 📁 Estrutura do Projeto
mqtt-dashboard/
├── backend/ # Servidor em Golang (API, MQTT, DB, WebSocket)
├── frontend/ # Aplicação React
└── README.md

## 🛡 Segurança

- Senhas de administradores são salvas com hash `bcrypt`
- Autenticação via tokens JWT protegendo endpoints sensíveis
- Acesso ao banco MongoDB Atlas com IP Whitelist e conexão segura
- Validações de entrada e sanitização no backend

---

## ⚙️ Como Rodar Localmente

### Requisitos
- Go instalado (https://go.dev/dl/)
- Node.js (https://nodejs.org/)
- Conta no MongoDB Atlas
- Broker MQTT (local ou externo, como test.mosquitto.org)

📌 Sobre
Desenvolvido para o Laboratório IoTec - UNIVALI

Substituto oficial do sistema anterior hospedado no diretório: x

🤝 Contribuição
Pull Requests são bem-vindos! Para contribuições maiores, por favor abra uma issue primeiro para discutir o que você gostaria de mudar.

📄 Licença
Distribuído sob a licença MIT.

👨‍🔧 Autor
Desenvolvido por Doc – com apoio do laboratório IoTec - UNIVALI

