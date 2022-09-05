import { useState } from "react";

export default function UseStateExample() {
    const [console, setConsole] = useState("Xbox");

    return (
      <>
        <h5>I like { console }.</h5>
        <button type="button" onClick={() => setConsole("Xbox")}>Xbox</button>{' '}
        <button type="button" onClick={() => setConsole("PS5")}>PS5</button>{' '}
        <button type="button" onClick={() => setConsole("Switch")}>Switch</button>
      </>
    );
}