import React, { useState, useEffect } from 'react';
import './PagAdmin.css';
import CardTopic from './components/CardTopic';
import StatusSensor from './components/StatusSensor';
import BotaoAdd from './components/botaoAdd';

function PagAdmin() {
  const [topicos, setTopicos] = useState([]);
  const [topicoSelecionado, setTopicoSelecionado] = useState(null);
  const [sensoresSelecionados, setSensoresSelecionados] = useState([]);

  // 1. Buscar tópicos ao montar
  useEffect(() => {
    fetch('http://localhost:8080/api/topicsJSON')  
      .then(response => response.json())
      .then(data => setTopicos(data))
      .catch(err => console.error('Erro ao buscar tópicos:', err));
  }, []);

  // 2. Buscar sensores quando topicoSelecionado mudar
  useEffect(() => {
    if (!topicoSelecionado) return;

    fetch(`http://localhost:8080/api/sensorJSON?topic=${topicoSelecionado}`) 
      .then(response => response.json())
      .then(data => setSensoresSelecionados(data))
      .catch(err => console.error('Erro ao buscar sensores:', err));
  }, [topicoSelecionado]);

  const handleSelecionarTopico = (nome) => {
    setTopicoSelecionado(nome);
  };

  const handleVoltar = () => {
    setTopicoSelecionado(null);
    setSensoresSelecionados([]);
  };

  const handleRemoverSensor = async (sensorId) => {
    try {
      const res = await fetch(`http://localhost:8080/api/sensorDelete/${sensorId}`, {
        method: 'DELETE'
      });

      if (!res.ok) throw new Error('Erro ao remover sensor');

      // Atualiza estado local dos sensores
      const novosSensores = sensoresSelecionados.filter(sensor => sensor.mqttID !== sensorId);
      setSensoresSelecionados(novosSensores);

    } catch (error) {
      console.error(error);
      alert('Falha ao remover sensor.');
    }
  };

  const handleAdded = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/topicsJSON');
      const data = await response.json();
      setTopicos(data);
    } catch (err) {
      console.error('Erro ao atualizar tópicos:', err);
    }
  };

  const handleSensorAdded = async () => {
  if (!topicoSelecionado) return;
  
  try {
    const response = await fetch(`http://localhost:8080/api/sensorJSON?topic=${topicoSelecionado}`);
    const data = await response.json();
    setSensoresSelecionados(data);
  } catch (err) {
    console.error('Erro ao atualizar sensores:', err);
  }
};

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
          <BotaoAdd 
            texto="Topico" 
            topico={null} 
            onSuccess={handleAdded} // sem "()"
          />
        </>
      ) : (
        <>
          <button onClick={handleVoltar} className="botao-voltar">← Voltar</button>
          <h3>Sensores do tópico: <em>{topicoSelecionado}</em></h3>
          <StatusSensor 
            sensores={sensoresSelecionados} 
            onRemoveSensor={handleRemoverSensor} 
            topico={topicoSelecionado}
            onSensorAdded={handleSensorAdded} 
          />
        </>
      )}
    </main>
  );
}

export default PagAdmin;
