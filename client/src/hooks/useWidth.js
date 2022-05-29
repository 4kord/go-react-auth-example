import React from "react";

export const useWidth = () => {
    const [width, setWidth] = React.useState(window.innerWidth)

    React.useEffect(() => {
        const handleResizeWindow = () => setWidth(window.innerWidth);
        window.addEventListener("resize", handleResizeWindow);
        return () => {
            window.removeEventListener("resize", handleResizeWindow);
        };
    }, []);

    return width
}