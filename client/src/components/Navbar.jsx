import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import AccountCircle from '@mui/icons-material/AccountCircle';
import MenuItem from '@mui/material/MenuItem';
import Menu from '@mui/material/Menu';
import { Box, Typography } from '@mui/material';
import { Link } from "react-router-dom";
import { useWidth } from '../hooks/useWidth';

export function Navbar() {
    const [auth] = React.useState({});
    const [menu1AnchorEl, setMenu1AnchorEl] = React.useState(null);
    const [menu2AnchorEl, setMenu2AnchorEl] = React.useState(null);

    const width = useWidth()

    const handleMenu1 = (event) => {
        setMenu1AnchorEl(event.currentTarget);
    };

    const handleMenu2 = (event) => {
        setMenu2AnchorEl(event.currentTarget);
    };
    
    const handleClose = () => {
        setMenu1AnchorEl(null);
        setMenu2AnchorEl(null);
    };

    return (
    <AppBar position="static">
        <Toolbar sx={{ display: "flex", justifyContent: "space-between"}}>
            {auth?.username && (
                <>
                    {width < 600 ? (
                        <Box sx={{display: "flex"}}>
                            <IconButton
                                size="large"
                                onClick={handleMenu1}
                                color="inherit"
                            >
                                <MenuIcon />
                            </IconButton>
                            <Menu
                                id="menu-appbar"
                                anchorEl={menu1AnchorEl}
                                anchorOrigin={{
                                    vertical: 'top',
                                    horizontal: 'left',
                                }}
                                keepMounted
                                transformOrigin={{
                                    vertical: 'top',
                                    horizontal: 'left',
                                }}
                                open={Boolean(menu1AnchorEl)}
                                onClose={handleClose}
                            >
                                <MenuItem onClick={handleClose}>El 1</MenuItem>
                                <MenuItem onClick={handleClose}>El 2</MenuItem>
                            </Menu>
                        </Box>
                    ) : (
                        <Box sx={{ display: "flex", mr: 2 }}>
                            <MenuItem>El 1</MenuItem>
                            <MenuItem>El 2</MenuItem>
                        </Box>
                    )}

                    <Box sx={{display: "flex", alignItems: "center"}}>
                        <Typography>
                            {auth.username}
                        </Typography>
                        <IconButton
                            size="large"
                            onClick={handleMenu2}
                            color="inherit"
                        >
                            <AccountCircle />
                        </IconButton>
                        <Menu
                            id="menu-appbar"
                            anchorEl={menu2AnchorEl}
                            anchorOrigin={{
                                vertical: 'top',
                                horizontal: 'right',
                            }}
                            keepMounted
                            transformOrigin={{
                                vertical: 'top',
                                horizontal: 'right',
                            }}
                            open={Boolean(menu2AnchorEl)}
                            onClose={handleClose}
                        >
                            <MenuItem onClick={handleClose}>Profile</MenuItem>
                            <MenuItem onClick={handleClose}>My account</MenuItem>
                        </Menu>
                    </Box>
                </>
            )}

            {!auth?.username && (
                <>
                    <Box sx={{display: "flex"}}>
                        <MenuItem component={Link} to="/" >Main</MenuItem>
                    </Box>
                    <Box sx={{display: "flex"}}>
                        <MenuItem component={Link} to="/signup" >Signup</MenuItem>
                        <MenuItem component={Link} to="/signin" >Signin</MenuItem>
                    </Box>
                </>
            )}
            
        </Toolbar>
    </AppBar>
  );
}
