import React from 'react';
import logo from './logo.svg';
import { Link } from "react-router-dom";
import './App.css';
import DarkMode from './DarkMode';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <DarkMode />
        <nav
          style={{
            borderBottom: "solid 1px",
            paddingBottom: "1rem",
          }}
        >
          <Link to="/useState">UseState</Link> |{" "}
          <Link to="/useEffect">UseEffect</Link> |{" "}
          <Link to="/useContext">UseContext</Link>
        </nav>

        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
