// components/Home.js
import React, { useEffect, useState } from 'react';
import CardSensor from './components/CardSensor';

function Home() {
  const [sensores, setSensores] = useState([]); // sempre inicializa como []

  useEffect(() => {
    const fetchSensores = async () => {
      try {
        const res = await fetch("http://localhost:8080/api/sensorView");
        if (!res.ok) throw new Error("Erro ao buscar sensores ativos");
        const data = await res.json();

        // garante que seja sempre array
        setSensores(Array.isArray(data) ? data : []);
      } catch (err) {
        console.error("Erro ao carregar sensores:", err);
        setSensores([]); // fallback
      }
    };

    fetchSensores();
  }, []);

  return (
    <main>
      <h1>Sensores ativos:</h1>
      <div className="cards-container">
        {(!sensores || sensores.length === 0) ? (
          <p>Nenhum sensor ativo no momento.</p>
        ) : (
          sensores.map(sensor => (
            <CardSensor 
              key={sensor.mqttID || sensor._id?.$oid} 
              nome={sensor.mqttID} 
              dados={sensor.dados || {}} 
            />
          ))
        )}
      </div>
    </main>
  );
}

export default Home;