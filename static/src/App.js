import React from 'react';
import './App.css';

import Division from './components/Division';

function App() {
  return (
    <div className="App">
      <Division divName="Header"/> 
      <Division divName="Body"/> 
      <Division divName="Footer"/> 
    </div>
  );
}

export default App;
