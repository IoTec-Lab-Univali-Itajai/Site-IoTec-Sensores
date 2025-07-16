import React from 'react';
import './Sidebar.css';

function Sidebar({ isOpen, onClose }) {
  return (
    <div className={`Sidebar ${isOpen ? 'open' : ''}`}>
      <button className="close-button" onClick={onClose}>
        &times;
      </button>
      <nav>
        <a href="/">Página Inicial</a>
        <a href="/adm">Administração</a>
        {/* Adicione mais links aqui */}
      </nav>
    </div>
  );
}

export default Sidebar;
