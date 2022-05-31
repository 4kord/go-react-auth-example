import { Route, Routes } from 'react-router-dom';
import { Layout } from './components/Layout';
import { Signup } from './pages/Signup';
import { Signin } from './pages/Signin';
import { Main } from './pages/Main';
import { RequireAuth } from './auth/RequireAuth';
import { Admin } from './pages/Admin';

function App() {
    return (
        <Routes>
            <Route path="/" element={<Layout />}>
                <Route path="/" element={<Main />} />
                <Route path="signup" element={<Signup />} />
                <Route path="signin" element={<Signin />} />

                <Route element={<RequireAuth allowedRoles={["admin"]} />}>
                    <Route path="admin" element={<Admin />} />
                </Route>
            </Route>
        </Routes>
    );
}

export default App;
