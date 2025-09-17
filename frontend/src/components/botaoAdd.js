import React, { useState } from 'react';
import './botaoAdd.css';
import './ModalForm.css';

function BotaoAdd({ texto, topico, onSuccess }) {
  const [showModal, setShowModal] = useState(false);
  
  // Estado inicial baseado no tipo (sensor ou tópico)
  const initialState = topico !== null 
    ? { 
        mqttID: '', 
        topicName: topico, 
        descricao: '', 
        showOnScreen: true // ← Valor padrão como boolean
      }
    : { nome: '' };
  
  const [formData, setFormData] = useState(initialState);

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    
    // Se for o campo showOnScreen, converter string para boolean
    if (name === "showOnScreen") {
      setFormData(prev => ({
        ...prev,
        [name]: value === "true" // Converte string para boolean
      }));
    } else {
      setFormData(prev => ({
        ...prev,
        [name]: value
      }));
    }
  };

  const handleSubmitSensor = async (e) => {
    e.preventDefault();
    
    try {
        const response = await fetch('http://localhost:8080/api/sensorCreate', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(formData) // Já está correto (showOnScreen é boolean)
        });

        if (!response.ok) throw new Error('Erro ao adicionar sensor');
        
        alert('Sensor adicionado com sucesso!');
        setShowModal(false);
        setFormData(initialState);
        
        if (typeof onSuccess === 'function') {
            onSuccess();
        }
    } catch (error) {
        console.error('Erro ao adicionar sensor:', error);
        alert(`Falha ao adicionar sensor: ${error.message}`);
    }
  };

  const handleSubmitTopic = async (e) => {
    e.preventDefault();
    
    try {
      const response = await fetch('http://localhost:8080/api/topicCreate', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData)
      });

      if (!response.ok) throw new Error('Erro ao adicionar tópico');
      
      alert('Tópico adicionado com sucesso!');
      setShowModal(false);
      setFormData(initialState);
      
      if (typeof onSuccess === 'function') {
        onSuccess();
      }
    } catch (error) {
      console.error('Erro ao adicionar tópico:', error);
      alert(`Falha ao adicionar tópico: ${error.message}`);
    }
  };

  const handleCloseModal = () => {
    setShowModal(false);
    setFormData(initialState);
  };

  return (
    <>
      <button 
        className="botao-add" 
        onClick={() => setShowModal(true)}
        aria-label={`Adicionar ${texto}`}
      >
        <h3>Adicionar {texto} +</h3>
      </button>

      {showModal && (
        <div className="modal-overlay">
          <div className="modal-content">
            <button 
              className="modal-close" 
              onClick={handleCloseModal}
              aria-label="Fechar modal"
            >
              &times;
            </button>
            
            <h3>Adicionar Novo {topico !== null ? 'Sensor' : 'Tópico'}</h3>
            
            <form onSubmit={topico !== null ? handleSubmitSensor : handleSubmitTopic}>
              {topico !== null ? (
                <>
                  <div className="form-group">
                    <label htmlFor="mqttID">ID MQTT:</label>
                    <input 
                      type="text" 
                      id="mqttID"
                      name="mqttID" 
                      value={formData.mqttID}
                      onChange={handleInputChange}
                      required
                      placeholder="Ex: sensor_temperatura_01"
                    />
                  </div>
                  
                  <div className="form-group">
                    <label htmlFor="descricao">Descrição:</label>
                    <textarea 
                      id="descricao"
                      name="descricao" 
                      value={formData.descricao}
                      onChange={handleInputChange}
                      placeholder="Descrição do sensor (opcional)"
                      rows="3"
                    />
                  </div>

                  <div className="form-group">
                    <label htmlFor="showOnScreen">Mostrar na Home:</label>
                    <select
                      id="showOnScreen"
                      name="showOnScreen"
                      value={formData.showOnScreen.toString()} // ← Converte boolean para string
                      onChange={handleInputChange}
                    >
                      <option value="true">Sim</option> {/* ← Strings */}
                      <option value="false">Não</option> {/* ← Strings */}
                    </select>
                  </div>
                </>
              ) : (
                <div className="form-group">
                  <label htmlFor="nome">Nome do Tópico:</label>
                  <input 
                    type="text" 
                    id="nome"
                    name="nome" 
                    value={formData.nome}
                    onChange={handleInputChange}
                    required
                    placeholder="Ex: temperatura_ambiente"
                  />
                </div>
              )}
              
              <div className="form-actions">
                <button 
                  type="button" 
                  onClick={handleCloseModal}
                  className="secondary-button"
                >
                  Cancelar
                </button>
                <button 
                  type="submit"
                  className="primary-button"
                >
                  Adicionar
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </>
  );
}

export default BotaoAdd;