import React from 'react';
import './CardSensor.css';

function CardSensor({ nome, dados }) {
  const sensorData = Array.isArray(dados) ? dados : [];

  console.log('CARDSENSOR.JS DIZ: Dados recebidos:', { nome, dados });
  
  return (
    <div className="chart-container">
      <h3>{nome}</h3>

      <div className="sensor-chart">
        {sensorData.length === 0 ? (
          <div className="chart-container sensor-chart-item">
            <div className="conteudo">
              <h1>N/A</h1>
              <p>Sem dados</p>
            </div>
          </div>
        ) : (
          sensorData.map((dado, index) => (
            <div key={index} className="chart-container sensor-chart-item">
              <div className="conteudo">
                <h1>
                  {dado.valor === null || dado.valor === undefined ? 
                    'N/A' : 
                    Number.isInteger(dado.valor) ? 
                    dado.valor : 
                    dado.valor.toFixed(2)
                  }
                  {/* Corrigido para usar camelCase */}
                {dado.unidade && dado.unidade !== "null" && dado.unidade !== null && (
                  <small>{dado.unidade}</small>
                )}
                </h1>
                <p>{dado.infoType || 'Dado'}</p>
              </div>
            </div>
          ))
        )}
      </div>
    </div>
  );
}

export default CardSensor;