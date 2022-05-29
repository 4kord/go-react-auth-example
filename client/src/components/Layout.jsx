import { Outlet } from "react-router-dom";
import { Navbar } from "./Navbar";
import CssBaseline from '@mui/material/CssBaseline';

export const Layout = () => {
    return (
        <>
            <CssBaseline />
            <Navbar />
            <Outlet />
        </>
    );
}