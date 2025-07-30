import React from 'react';
import './ConfirmationModal.css';

function ConfirmationModal({ isOpen, onClose, onConfirm, message }) {
  if (!isOpen) return null;

  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <h3>Confirmar Exclusão</h3>
        <p>{message}</p>
        <div className="modal-buttons">
          <button onClick={onClose} className="cancel-button">Cancelar</button>
          <button onClick={onConfirm} className="confirm-button">Confirmar</button>
        </div>
      </div>
    </div>
  );
}

export default ConfirmationModal;