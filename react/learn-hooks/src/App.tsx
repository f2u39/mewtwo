import React from 'react';
import logo from './logo.svg';
import { Link } from "react-router-dom";
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
      <nav
        style={{
          borderBottom: "solid 1px",
          paddingBottom: "1rem",
        }}
      >
        <Link to="/useState">UseState</Link> |{" "}
        <Link to="/useEffect">UseEffect</Link>
      </nav>
      </header>
    </div>
  );
}

export default App;
