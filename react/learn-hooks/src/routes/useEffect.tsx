import { useEffect, useState } from "react";

export default function useEffectExample() {
    const [console, setConsole] = useState("Xbox");

    useEffect(()=> {
        const interval = setInterval(() => {
            setConsole("PS5");
          }, 1000);
    }, [])

    return (
        <>
            <h5>I like { console }.</h5>
        </>
    );
}