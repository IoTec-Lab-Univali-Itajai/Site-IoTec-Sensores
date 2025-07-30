import React, { useState } from 'react';
import './PagAdmin.css';
import CardTopic from './components/CardTopic';
import StatusSensor from './components/StatusSensor';

// Dados iniciais (poderia ser movido para um arquivo separado)
let topicosIniciais = [
  {
    nome: "iot/lab/sala1",
    sensores: [
      {
        id: "E301",
        nome: "Sensor Temperatura 1",
        tipoDados: ["temperatura", "umidade"],
        ultimaAtualizacao: "2025-07-30 14:25",
      },
      {
        id: "B700",
        nome: "Sensor Pressão",
        tipoDados: ["pressao"],
        ultimaAtualizacao: "2025-07-28 13:50",
      },
      {
        id: "B701",
        nome: "Sensor Pressão 2",
        tipoDados: ["pressao"],
        ultimaAtualizacao: "2025-07-22 13:50",
      },
    ],
  },
  {
    nome: "iot/lab/sala2",
    sensores: [
      {
        id: "C400",
        nome: "Sensor Vibração",
        tipoDados: ["vibracao", "deslocamento_vertical"],
        ultimaAtualizacao: "2025-07-30 14:10",
      },
      {
        id: "C401",
        nome: "Sensor Vibração",
        tipoDados: ["vibracao", "deslocamento_vertical"],
        ultimaAtualizacao: "2025-07-30 14:10",
      },
    ],
  },
];

function PagAdmin() {
  const [topicoSelecionado, setTopicoSelecionado] = useState(null);
  const [topicos, setTopicos] = useState(topicosIniciais);

  const handleSelecionarTopico = (nome) => {
    setTopicoSelecionado(nome);
  };

  const handleVoltar = () => {
    setTopicoSelecionado(null);
  };

  const handleRemoverSensor = (sensorId) => {
    // Atualiza o array de tópicos
    const novosTopicos = topicos.map(topico => {
      if (topico.nome === topicoSelecionado) {
        return {
          ...topico,
          sensores: topico.sensores.filter(sensor => sensor.id !== sensorId)
        };
      }
      return topico;
    });
    
    // Atualiza o estado e a variável let
    setTopicos(novosTopicos);
    topicosIniciais = novosTopicos;
  };

  const sensoresSelecionados = topicos.find(t => t.nome === topicoSelecionado)?.sensores || [];

  return (
    <main className="pagadmin-main">
      <h2>Área de Administração</h2>
      
      {!topicoSelecionado ? (
        <>
          <p>Selecione um tópico para visualizar os sensores:</p>
          <div className="topicos-container">
            {topicos.map((topico) => (
              <CardTopic
                key={topico.nome}
                nome={topico.nome}
                onSelect={handleSelecionarTopico}
              />
            ))}
          </div>
        </>
      ) : (
        <>
          <button onClick={handleVoltar} className="botao-voltar">
            ← Voltar para tópicos
          </button>
          <h3>Sensores do tópico: <em>{topicoSelecionado}</em></h3>
          <StatusSensor 
            sensores={sensoresSelecionados} 
            onRemoveSensor={handleRemoverSensor}
          />
        </>
      )}
    </main>
  );
}

export default PagAdmin;