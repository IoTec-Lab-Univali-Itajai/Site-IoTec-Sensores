import React, { useState } from 'react';
import './Header.css';
import Sidebar from './Sidebar';
import { FaBars } from 'react-icons/fa';

function Header() {
  const [sidebarOpen, setSidebarOpen] = useState(false);

  const toggleSidebar = () => {
    setSidebarOpen(!sidebarOpen);
  };

  return (
    <>
      <header className="AppHeader">
        {/* Esquerda */}
        <button className="menu-button" onClick={toggleSidebar}>
          <FaBars />
        </button>

        {/* Centro absoluto */}
        <div className="header-center">
          <h1>IoTec Lab</h1>
          <img src="/logo-univali.png" alt="Logo UNIVALI" />
        </div>
      </header>

      <Sidebar isOpen={sidebarOpen} onClose={toggleSidebar} />
    </>
  );
}

export default Header;
