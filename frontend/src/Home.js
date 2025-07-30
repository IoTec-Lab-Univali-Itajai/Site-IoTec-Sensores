// components/Home.js
import React from 'react';
import CardSensor from './components/CardSensor';

function Home() {
  let sensores = [
    { id: 1, nome: 'Sensor 1', dados: { temperatura: 20, umidade: 55 } },
    { id: 2, nome: 'Sensor 2', dados: { temperatura: 23, indice_calor: 32 } },
    { id: 3, nome: 'Sensor 3', dados: { deslocamento_vertical: 5 } },
    { id: 4, nome: 'Sensor 4', dados: { temperatura: 24, umidade: 60, pressao: 10, deslocamento_vertical: 90 }},
    { id: 1, nome: 'Sensor 5', dados: { temperatura: 24, umidade: 60, pressao: 10, deslocamento_vertical: 90 } },
    { id: 2, nome: 'Sensor 6', dados: { temperatura: 23, indice_calor: 32 } },
    { id: 3, nome: 'Sensor 7', dados: { deslocamento_vertical: 5 } },
    { id: 4, nome: 'Sensor 8', dados: { temperatura: 23, indice_calor: 32 }},
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