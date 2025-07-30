import React from 'react';
import './CardTopic.css';

function CardTopic({ nome, onSelect }) {
  return (
    <div className="card-topic" onClick={() => onSelect(nome)}>
      <h3>{nome}</h3>
    </div>
  );
}

export default CardTopic;