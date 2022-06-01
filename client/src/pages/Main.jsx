import { useAuth } from "../auth/useAuth"
import { Typography } from "@mui/material";
import { Box } from "@mui/system";
import { Link } from "react-router-dom";

export const Main = () => {
    const { auth } = useAuth()    
    return (
        <>
            {auth?.username ? (
                <Box>
                    <Typography>You are authorized</Typography>
                    <Link to="/admin">AFsdfs</Link>
                </Box>
            ) : (
                <Typography>You are unauthorized</Typography>
            )}
        </>
    );
}