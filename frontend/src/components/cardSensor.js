import React from 'react';
import './CardSensor.css';

function CardSensor({ nome, dados }) {
  return (
    <div className="chart-container">
      <h3>{nome}</h3>

      <div className="sensor-chart">
        {Object.entries(dados).map(([tipo, valor]) => (
          <div key={tipo} className="chart-container sensor-chart-item">
            <div className="conteudo">
              <h1>{valor === null ? 'N/A' : valor}</h1>
              <p>{tipo.replace(/_/g, ' ')}</p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default CardSensor;
