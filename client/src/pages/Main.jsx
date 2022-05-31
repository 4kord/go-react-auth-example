import { useAuth } from "../auth/useAuth"
import { Typography } from "@mui/material";

export const Main = () => {
    const { auth } = useAuth()    
    return (
        <>
            {auth?.username ? (
                <Typography>You are authorized</Typography>
            ) : (
                <Typography>You are unauthorized</Typography>
            )}
        </>
    );
}