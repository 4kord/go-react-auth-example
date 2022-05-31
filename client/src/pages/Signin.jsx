import React from "react";
import { Avatar, Button, Grid, TextField, Typography } from "@mui/material";
import { Box, Container } from "@mui/system";
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import { Link as RouterLink } from "react-router-dom";
import Link from '@mui/material/Link';
import axios from "../api/axios";
import { useAuth } from "../auth/useAuth";
import { useAxiosFunction } from "../auth/useAxiosFunction";
import { useNavigate, useLocation } from "react-router-dom";
import { Loading } from "../components/Loading";
import { ErrorAlert } from "../components/ErrorAlert";

export const Signin = () => {
    const [username, setUsername] = React.useState("");
    const [password, setPassword] = React.useState("");

    const { setAuth } = useAuth()
    const [response, error, loading, axiosFetch] = useAxiosFunction();

    const navigate = useNavigate()
    const location = useLocation();
    const from = location.state?.from?.pathname || "/";

    const submit = async e => {
        e.preventDefault();

        await axiosFetch({
            axiosInstance: axios,
            method: "POST",
            url: "/auth/login",
            requestConfig: {
                data: {
                    username,
                    password
                }
            }
        })

        setUsername("");
        setPassword("");
    }
 
    React.useEffect(() => {
        if (response && !error && !loading) {
            setAuth(response);
            navigate(from, { replace: true });
        }
    }, [error, from, loading, navigate, response, setAuth]);

    return (
        <>
            {loading && <Loading />}

            {!loading && error && <ErrorAlert message={error} />}

            <Container component="main" maxWidth="xs">
                <Box
                    sx={{
                        marginTop: 8,
                        display: "flex",
                        flexDirection: "column",
                        alignItems: "center",
                    }}
                >
                    <Avatar sx={{ m: 1, bgcolor: "secondary.main" }}>
                        <LockOutlinedIcon />
                    </Avatar>
                    <Typography component="h1" variant="h5">
                        Sign in
                    </Typography>
                    <Box component="form" onSubmit={submit} sx={{ mt: 3, width: 1}}>
                        <Grid container spacing={2}>
                            <Grid item xs={12}>
                                <TextField
                                    name="username"
                                    required
                                    id="username"
                                    label="Username"
                                    autoFocus
                                    fullWidth
                                    onChange={e => setUsername(e.target.value)}
                                    value={username}
                                />
                            </Grid>
                            <Grid item xs={12}>
                                <TextField
                                    name="password"
                                    required
                                    id="password"
                                    label="Password"
                                    type="password"
                                    fullWidth
                                    onChange={e => setPassword(e.target.value)}
                                    value={password}
                                />
                            </Grid>
                        </Grid>
                        <Button
                            type="submit"
                            fullWidth
                            variant="contained"
                            sx={{ mt: 3, mb: 2 }}
                        >
                            Sign In
                        </Button>
                        <Grid container justifyContent="flex-end">
                            <Grid item>
                                <Link component={RouterLink} to="/signup" >
                                    Don't have an account? Sign up
                                </Link>
                            </Grid>
                        </Grid>
                    </Box>
                </Box>           
            </Container>
        </>
    );
}