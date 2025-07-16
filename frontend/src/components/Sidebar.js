// components/Sidebar.js
import React from 'react';
import './Sidebar.css';
import { Link } from 'react-router-dom';

function Sidebar({ isOpen, onClose }) {
  return (
    <div className={`Sidebar ${isOpen ? 'open' : ''}`}>
      <button className="close-button" onClick={onClose}>
        &times;
      </button>
      <nav>
        <Link to="/" onClick={onClose}>Página Inicial</Link>
        <Link to="/adm" onClick={onClose}>Administração</Link>
      </nav>
    </div>
  );
}

export default Sidebar;