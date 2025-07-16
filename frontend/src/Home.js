// components/Home.js
import React from 'react';
import CardSensor from './components/CardSensor';

function Home() {
  const sensores = [
    { id: 1, nome: 'Sensor 1', dados: { temperatura: 20, umidade: 55 } },
    { id: 2, nome: 'Sensor 2', dados: { temperatura: 23, indice_calor: 32 } },
    { id: 3, nome: 'Sensor 3', dados: { deslocamento_vertical: 5 } },
  ];

  return (
    <main>
      <h1>Sensores ativos: </h1>
      <div className="cards-container">
        {sensores.map(sensor => (
          <CardSensor key={sensor.id} nome={sensor.nome} dados={sensor.dados} />
        ))}
      </div>
    </main>
  );
}

export default Home;