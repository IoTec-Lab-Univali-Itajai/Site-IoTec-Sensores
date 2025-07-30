import React, { useState } from 'react';
import { FaTrash } from 'react-icons/fa';
import './StatusSensor.css';

const getStatusColor = (ultimaAtualizacao) => {
  const agora = new Date();
  const dataAtualizacao = new Date(ultimaAtualizacao);
  const diffDias = Math.floor((agora - dataAtualizacao) / (1000 * 60 * 60 * 24));

  if (diffDias <= 1) return 'green';
  if (diffDias <= 7) return 'yellow';
  return 'red';
};

function StatusSensor({ sensores, onRemoveSensor }) {
  const [sensorParaRemover, setSensorParaRemover] = useState(null);

  const confirmarRemocao = (sensor) => {
    if (window.confirm(`Tem certeza que deseja remover o sensor "${sensor.nome}" (ID: ${sensor.id})?`)) {
      onRemoveSensor(sensor.id);
    }
  };

  if (!sensores || sensores.length === 0) {
    return <p>Nenhum sensor registrado neste tópico.</p>;
  }

  return (
    <div className="sensor-container">
      <div className="sensor-list">
        {sensores.map((sensor) => (
          <div key={sensor.id} className="sensor-card">
            <div className="sensor-info">
              <h4>{sensor.nome}</h4>
              <p><strong>ID:</strong> {sensor.id}</p>
              <p><strong>Tipo:</strong> {sensor.tipoDados.join(', ')}</p>
              <div className="status-dot-container">
                <span className={`status-dot ${getStatusColor(sensor.ultimaAtualizacao)}`}></span>
              </div>
            </div>
            <button 
              onClick={() => confirmarRemocao(sensor)}
              className="delete-button"
              aria-label="Remover sensor"
            >
              <FaTrash />
            </button>
          </div>
        ))}
      </div>
      
      <div className="status-legend">
        <h4>Legenda de Status:</h4>
        <div className="legend-items">
          <div className="legend-item">
            <span className="status-dot green"></span>
            <span>Atualizado (últimas 24h)</span>
          </div>
          <div className="legend-item">
            <span className="status-dot yellow"></span>
            <span>Recente (última semana)</span>
          </div>
          <div className="legend-item">
            <span className="status-dot red"></span>
            <span>Desatualizado (+1 semana)</span>
          </div>
        </div>
      </div>
    </div>
  );
}

export default StatusSensor;