import { Route, Routes } from 'react-router-dom';
import { Layout } from './components/Layout';
import { Signup } from './pages/Signup';
import { Signin } from './pages/Signin';

function App() {
    return (
        <Routes>
            <Route path="/" element={<Layout />}>
                <Route path="signup" element={<Signup />} />
                <Route path="signin" element={<Signin />} />
            </Route>
        </Routes>
    );
}

export default App;
