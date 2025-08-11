import React from 'react';
import './botaoAdd.css';

function BotaoAdd({ texto }) {
  return (
    <button className="botao-add">
      <h3>
        Adicionar {texto} +
      </h3>
    </button>
  );
}

export default BotaoAdd;
