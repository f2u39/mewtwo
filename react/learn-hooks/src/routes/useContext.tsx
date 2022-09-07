import { useState, createContext, useContext } from "react";

const GameContext = createContext("");

export default function UseContextExample() {
    const [game, setGame] = useState("Halo");

    return (
        <GameContext.Provider value={ game }>
            <h2>Xbox</h2>
            <h3>{`Playing ${ game }!`}</h3>
            <PS5 />
        </GameContext.Provider>
    );
}

function PS5() {
    const game = useContext(GameContext);
    return (
        <>
            <h2>PS5</h2>
            <h3>{`Playing ${ game }!`}</h3>
        </>
    );
}