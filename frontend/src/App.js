// App.js
import React from 'react';
import './App.css';
import Header from './components/Header';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import PagAdmin from './PagAdmin';
import Home from './Home'; // Vamos criar este componente

function App() {
  return (
    <Router>
      <div className="App">
        <Header />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/adm" element={<PagAdmin />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;