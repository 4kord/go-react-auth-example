import React from 'react';
import { Route, Routes } from 'react-router-dom';
import { Layout } from './components/Layout';
import { Signup } from './pages/Signup';
import { Signin } from './pages/Signin';
import { Main } from './pages/Main';
import { RequireAuth } from './auth/RequireAuth';
import { PersistLogin } from './auth/PersistLogin';
import { User } from './pages/User';

function App() {
    return (
        <Routes>
            <Route path="/" element={<Layout />}>
            <Route element={<PersistLogin />}>
                <Route path="/" element={<Main />} />
                <Route path="signup" element={<Signup />} />
                <Route path="signin" element={<Signin />} />

                    <Route element={<RequireAuth allowedRoles={["user"]} />}>
                        <Route path="user" element={<User />} />
                    </Route>
                </Route>
            </Route>
        </Routes>
    );
}

export default App;
