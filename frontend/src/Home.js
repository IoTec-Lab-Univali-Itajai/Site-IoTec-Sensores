import React, { useEffect, useState } from 'react';
import CardSensor from './components/cardSensor';

function Home() {
  const [sensores, setSensores] = useState([]);

  useEffect(() => {
    const fetchSensores = async () => {
      try {
        const res = await fetch("http://10.1.203.113:8080/api/sensorView");
        if (!res.ok) throw new Error("Erro ao buscar sensores ativos");
        const data = await res.json();

        // Mapeia a estrutura do backend para o que o frontend espera
        const sensoresFormatados = data.map(sensor => ({
          mqttID: sensor.SensorID, // SensorID do backend vira mqttID no frontend
          dados: sensor.SensorData // SensorData do backend vira dados no frontend
        }));

        setSensores(sensoresFormatados);
        console.log('HOME.JS DIZ: Dados recebidos:', sensoresFormatados);
      } catch (err) {
        console.error("Erro ao carregar sensores:", err);
        setSensores([]);
      }
    };

    fetchSensores();
    // Opcional: atualizar a cada X segundos
    const interval = setInterval(fetchSensores, 30000); // Atualiza a cada 30 segundos
    return () => clearInterval(interval);
  }, []);

  return (
    <main>
      <h1>Sensores ativos:</h1>
      <div className="cards-container">
        {sensores.length === 0 ? (
          <p>Nenhum sensor ativo no momento.</p>
        ) : (
          sensores.map(sensor => (
            <CardSensor 
              key={sensor.mqttID} 
              nome={sensor.mqttID} 
              dados={sensor.dados || []} // Garante que dados seja array
            />
          ))
        )}
      </div>
    </main>
  );
}

export default Home;