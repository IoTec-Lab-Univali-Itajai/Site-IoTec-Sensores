# 🌐 Site IoTec Sensores – UNIVALI

Este é o novo sistema web desenvolvido para o **Laboratório IoTec da UNIVALI (campus Itajaí)**. O objetivo é substituir o site anterior (`diretório/link X`) e fornecer uma plataforma moderna, dinâmica e segura para **monitoramento e gerenciamento de sensores via MQTT**.

---

## 🧩 Tecnologias Utilizadas

- **Frontend:** React.js
- **Backend:** Golang (Go)
- **Mensageria:** MQTT
- **Banco de Dados:** MongoDB Atlas (cloud)
- **Estilo:** TailwindCSS (ou outro a definir)
- **Hospedagem:** (a definir)
- **Controle de versão:** Git + GitHub

---

## 🧠 Funcionalidades

### 👥 Modos do sistema

#### 🔍 Modo de Exibição (Público)
- Visualização de cards com dados dos sensores em tempo real.
- Gráficos dinâmicos baseados nas informações enviadas.
- Interface limpa e responsiva.

#### 🔐 Modo de Administração (Restrito)
- Login com usuário e senha (armazenados de forma segura).
- Criação de novos cards sensoriais:
  - ID do sensor
  - Tópico MQTT escutado
  - Opções de exibição: valores e gráficos
- Edição e exclusão de cards existentes.
- Os dados são salvos dinamicamente no banco e refletidos no front-end.

---

## 📡 Comunicação MQTT

Os sensores enviam mensagens via tópicos MQTT com o seguinte formato:

temperatura, 23, umidade, 50, pressão, 100

yaml
Copiar
Editar

> Se algum dado não estiver disponível, o sensor envia `NULL` no lugar do valor.

O backend escuta continuamente os tópicos e:
- Armazena os dados no MongoDB Atlas
- Atualiza o front-end em tempo real via WebSocket ou polling (a definir)

---

## 🔐 Segurança

- Senhas do ADM com hash seguro
- Autenticação protegida
- Acesso ao modo ADM restrito
- Banco de dados protegido com regras e usuários próprios no MongoDB Atlas

---

## 🛠️ Como rodar localmente

**Pré-requisitos:**
- Node.js + npm
- Go instalado
- MongoDB Atlas com URI configurada
- Broker MQTT (ex: Mosquitto) local ou remoto

**Passos iniciais:**

```bash
# Clonar o repositório
git clone https://github.com/IoTec-Lab-Univali-Itajai/Site-IoTec-Sensores
cd Site-IoTec-Sensores

# Instalar dependências do front-end
cd frontend
npm install
npm start

# Em outro terminal, rodar o backend
cd ../backend
go run main.go
📁 Estrutura do Projeto
bash
Copiar
Editar
Site-IoTec-Sensores/
├── backend/             # Código Go: MQTT, API, MongoDB
├── frontend/            # React: Interface pública e ADM
├── README.md
└── .gitignore
👨‍🔬 Desenvolvido por
Laboratório IoTec – UNIVALI - Itajaí
Projeto de Inovação em Monitoramento e Controle de Sensores

📝 Licença
Distribuído sob a licença MIT.
