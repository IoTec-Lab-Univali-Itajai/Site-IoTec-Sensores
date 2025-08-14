import React, { useState } from 'react';
import { FaTrash } from 'react-icons/fa';
import './StatusSensor.css';
import BotaoAdd from './botaoAdd';

const getStatusColor = (ultimaAtualizacao) => {
  const DATA_PADRAO = new Date("1977-11-18T12:00:00Z");
  const agora = new Date();
  const dataAtualizacao = new Date(ultimaAtualizacao);
  const diffDias = Math.floor((agora - dataAtualizacao) / (1000 * 60 * 60 * 24));

  if (dataAtualizacao.getTime() === DATA_PADRAO.getTime()) {
    return 'black';
  }
  if (diffDias <= 1) return 'green';
  if (diffDias <= 7) return 'yellow';
  return 'red';
};

function StatusSensor({ sensores, onRemoveSensor, topico, onSensorAdded }) {

  const confirmarRemocao = (sensor) => {
    if (window.confirm(`Tem certeza que deseja remover o sensor "${sensor.nome}" (ID: ${sensor.mqttID})?`)) {
      onRemoveSensor(sensor.mqttID);
    }
  };

  if (!sensores || sensores.length === 0) {
    return <p>Nenhum sensor registrado neste tópico.</p>;
  }

  return (
    <div className="sensor-container">
      <div className="sensor-list">
        {sensores.map((sensor) => (
          <div key={sensor._id?.$oid || sensor.mqttID} className="sensor-card">
            <div className="sensor-info">
              <h4>{sensor.nome}</h4>
              <p><strong>ID:</strong> {sensor.mqttID}</p>
              <div className="status-dot-container">
                <span className={`status-dot ${getStatusColor(sensor.lastUpdate || sensor.ultimaAtualizacao)}`}></span>
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

      <div className="botao-add-wrapper">
        <BotaoAdd 
          texto="Sensor" 
          topico={topico} 
          onSuccess={() => {
            // Isso será tratado pelo componente pai
            if (typeof onSensorAdded === 'function') {
              onSensorAdded();
            }
          }} 
        />
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
           <div className="legend-item">
            <span className="status-dot black"></span>
            <span>Sem dados recebidos</span>
          </div>
        </div>
      </div>
    </div>
  );
}

export default StatusSensor;