import { ThemeProvider } from '@emotion/react';
import { createTheme } from '@mui/material/styles';
import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import App from './App';

const darkTheme = createTheme({
    palette: {
        mode: 'dark',
    },
});

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <React.StrictMode>
        <BrowserRouter>
            <ThemeProvider theme={darkTheme}>
                <Routes>
                    <Route path="/*" element={<App />} />
                </Routes>
            </ThemeProvider>
        </BrowserRouter>
    </React.StrictMode>
);
