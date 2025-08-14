import React, { useState, useEffect } from 'react';
import './PagAdmin.css';
import CardTopic from './components/CardTopic';
import StatusSensor from './components/StatusSensor';

function PagAdmin() {
  const [topicos, setTopicos] = useState([]);
  const [topicoSelecionado, setTopicoSelecionado] = useState(null);

  useEffect(() => {
    fetch('http://localhost:8080/api/topics')  // ajuste porta se necessário
      .then(response => response.json())
      .then(data => setTopicos(data))
      .catch(err => console.error('Erro ao buscar tópicos:', err));
  }, []);

  const handleSelecionarTopico = (nome) => {
    setTopicoSelecionado(nome);
  };

  const handleVoltar = () => {
    setTopicoSelecionado(null);
  };

  const handleRemoverSensor = async (sensorId) => {
    try {
      const res = await fetch(`http://localhost:8080/api/sensor/${sensorId}`, {
        method: 'DELETE'
      });

      if (!res.ok) throw new Error('Erro ao remover sensor');

      // Atualiza estado local
      const novosTopicos = topicos.map(topico => {
        if (topico.nome === topicoSelecionado) {
          return {
            ...topico,
            sensores: topico.sensores.filter(sensor => sensor.mqttID !== sensorId)
          };
        }
        return topico;
      });

      setTopicos(novosTopicos);
    } catch (error) {
      console.error(error);
      alert('Falha ao remover sensor.');
    }
  };

  const handleSensorAdded = async () => {
  try {
    const response = await fetch('http://localhost:8080/api/topics');
    const data = await response.json();
    setTopicos(data);
  } catch (err) {
    console.error('Erro ao atualizar tópicos:', err);
  }
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
              <CardTopic key={topico.nome} nome={topico.nome} onSelect={handleSelecionarTopico} />
            ))}
          </div>
        </>
      ) : (
        <>
          <button onClick={handleVoltar} className="botao-voltar">← Voltar</button>
          <h3>Sensores do tópico: <em>{topicoSelecionado}</em></h3>
          <StatusSensor 
            sensores={sensoresSelecionados} 
            onRemoveSensor={handleRemoverSensor} 
            topico={topicoSelecionado}
            onSensorAdded={handleSensorAdded}  // Adicione esta linha
          />
        </>
      )}
    </main>
  );
}

export default PagAdmin;
