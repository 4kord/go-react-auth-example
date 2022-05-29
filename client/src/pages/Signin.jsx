import { Avatar, Button, Grid, TextField, Typography } from "@mui/material";
import { Box, Container } from "@mui/system";
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import { Link as RouterLink } from "react-router-dom";
import Link from '@mui/material/Link';

export const Signin = () => {
    return (
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
                <Box component="form" sx={{ mt: 3, width: 1}}>
                    <Grid container spacing={2}>
                        <Grid item xs={12}>
                            <TextField
                                name="username"
                                required
                                id="username"
                                label="Username"
                                autoFocus
                                fullWidth
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
                            <Link component={RouterLink} to="/signup" >
                                Don't have an account? Sign up
                            </Link>
                        </Grid>
                    </Grid>
                </Box>
            </Box>           
        </Container>
    );
}