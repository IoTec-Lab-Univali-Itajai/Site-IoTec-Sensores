import React from 'react';
import './App.css';
import Header from './components/Header';
import CardSensor from './components/CardSensor';

function App() {

  const sensores = [
    { id: 1, nome: 'Sensor 1', dados: { temperatura: 20, umidade: 55 } },
    { id: 2, nome: 'Sensor 2', dados: { temperatura: 23, indice_calor: 32 } },
    { id: 3, nome: 'Sensor 3', dados: { deslocamento_vertical: 5 } },
    // Adicione mais sensores
  ];

  return (
    <div className="App">
      <Header />

      <main>
        <h1>Sensores ativos: </h1>
        <div className="cards-container">
          {sensores.map(sensor => (
            <CardSensor key={sensor.id} nome={sensor.nome} dados={sensor.dados} />
          ))}
        </div>
      </main>
    </div>
  );
}

export default App;
