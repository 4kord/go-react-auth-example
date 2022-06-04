import React from "react";
import { Avatar, Button, Grid, TextField, Typography } from "@mui/material";
import { Box, Container } from "@mui/system";
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import { Link as RouterLink, useNavigate } from "react-router-dom";
import Link from '@mui/material/Link';
import { useAxiosFunction } from "../auth/useAxiosFunction";
import { axiosPublic } from "../api/axios";
import { Loading } from "../components/Loading";
import { ErrorAlert } from "../components/ErrorAlert";

export const Signup = () => {
    const [username, setUsername] = React.useState("");
    const [password, setPassword] = React.useState("");

    const [response, error, loading, axiosFetch] = useAxiosFunction();

    const navigate = useNavigate();
    
    const submit = async e => {
        e.preventDefault();

        await axiosFetch ({
            axiosInstance: axiosPublic,
            method: "POST",
            url: "auth/register",
            requestConfig: {
                data: {
                    username,
                    password
                }
            }
        });

        setUsername("");
        setPassword("");
    }

    React.useEffect(() => {
        if (response && !error && !loading) {
            navigate("/signin");
        }
    }, [error, loading, navigate, response])

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
                        Sign up
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
                            Sign Up
                        </Button>
                        <Grid container justifyContent="flex-end">
                            <Grid item>
                                <Link component={RouterLink} to="/signin" >
                                    Already have an account? Sign in
                                </Link>
                            </Grid>
                        </Grid>
                    </Box>
                </Box>           
            </Container>
        </>
    );
}