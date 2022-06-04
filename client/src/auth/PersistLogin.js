import { useEffect, useState } from "react";
import { Outlet } from "react-router-dom";
import { axiosPrivate, axiosPublic } from "../api/axios";
import { Loading } from "../components/Loading";
import { useAuth } from "./useAuth";

export const PersistLogin = () => {
    const [isLoading, setIsLoading] = useState(true);
    const { auth, setAuth } = useAuth();

    useEffect(() => {
        const refresh = async () => {
            const response = await axiosPublic.get("/auth/refresh");
    
            return response.data.accessToken;
        }

        const getUser = async () => {
            const response = await axiosPrivate.get("/user/me");
    
            return response.data;
        }

        const verifyRefreshToken = async () => {
            try {
                const accessToken = await refresh();
                window.accessToken = accessToken;
                const userRes = await getUser();
                setAuth(userRes);
            } catch(err) {
                console.log(err)
            } finally {
                setTimeout(() => {
                    setIsLoading(false)
                }, 500)
            }
        }

        !auth?.username ? verifyRefreshToken() : setTimeout(() => setIsLoading(false), 500);
    }, [auth?.username, setAuth])

    return (
        <>
            {isLoading
                ? <Loading />
                : <Outlet />
            }
        </>
    )
}