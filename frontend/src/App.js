import React from 'react';
import './App.css';
import Header from './components/Header';

function App() {
  return (
    <div className="App">
      <Header />

      <main>
        <h2>Bem-vindo ao painel de sensores</h2>
        <p>Aqui serão exibidos os dados em tempo real.</p>
      </main>
    </div>
  );
}

export default App;
