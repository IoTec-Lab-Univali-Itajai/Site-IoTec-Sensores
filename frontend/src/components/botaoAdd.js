import React, { useState } from 'react';
import './botaoAdd.css';
import './ModalForm.css'; // Você precisará criar este CSS

function BotaoAdd({ texto, topico, onSuccess }) {
  const [showModal, setShowModal] = useState(false);
  const [formData, setFormData] = useState({
    mqttID: '',
    topicName: topico, // Agora enviamos o nome do tópico
    descricao: '',
    showOnScreen: false
  });

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    try {
      const response = await fetch('http://localhost:8080/api/sensor', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData)
      });

      if (!response.ok) throw new Error('Erro ao adicionar sensor');

      alert('Sensor adicionado com sucesso!');
      setShowModal(false);
      // Chame o callback de sucesso
      if (typeof onSuccess === 'function') {
        onSuccess();
      }
    } catch (error) {
      console.error(error);
      alert('Falha ao adicionar sensor: ' + error.message);
    }
};

  return (
    <>
      <button className="botao-add" onClick={() => setShowModal(true)}>
        <h3>Adicionar {texto} +</h3>
      </button>

      {showModal && (
        <div className="modal-overlay">
          <div className="modal-content">
            <h3>Adicionar Novo Sensor</h3>
            <form onSubmit={handleSubmit}>
      
              <div className="form-group">
                <label>ID MQTT:</label>
                <input 
                  type="text" 
                  name="mqttID" 
                  value={formData.mqttID}
                  onChange={handleInputChange}
                  required
                />
              </div>
              
              <div className="form-group">
                <label>Descrição:</label>
                <textarea 
                  name="descricao" 
                  value={formData.descricao}
                  onChange={handleInputChange}
                />
              </div>

              <div className="form-group">
                <label>Mostrar na Home:</label>
                <select
                  name="showOnScreen"
                  value={formData.showOnScreen}
                  onChange={handleInputChange}
                >
                  <option value="true">Sim</option>
                  <option value="false">Não</option>
                </select>
              </div>
              
              <div className="form-actions">
                <button type="button" onClick={() => setShowModal(false)}>
                  Cancelar
                </button>
                <button type="submit">Adicionar</button>
              </div>
            </form>
          </div>
        </div>
      )}
    </>
  );
}

export default BotaoAdd;