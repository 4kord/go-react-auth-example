import axios from "../api/axios"
import { useAuth } from "./useAuth"

export const useRefreshToken = () => {
    const { setAuth } = useAuth();
    
    const refresh = async () => {
        const response = await axios.get("/auth/refresh");
        setAuth(response.data);

        return response.data.accessToken;
    }

    return refresh
}