import { useEffect, useState } from "react";

export default function useEffectExample() {
    const [console, setConsole] = useState("PS5");

    useEffect(()=> {
        const interval = setInterval(() => {
            setConsole("Xbox");
          }, 1000);
    }, [])

    return (
        <>
            <h5>I like { console }.</h5>
        </>
    );
}