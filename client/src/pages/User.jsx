import React from "react";
import { useAxiosFunction } from "../auth/useAxiosFunction";
import { Loading } from "../components/Loading";
import { Typography } from "@mui/material";
import { Box } from "@mui/material";
import { ErrorAlert } from "../components/ErrorAlert";
import { axiosPrivate } from "../api/axios";

export const User = () => {
    const [response, error, loading, axiosFetch] = useAxiosFunction();



    React.useEffect(() => {
        axiosFetch({
            axiosInstance: axiosPrivate,
            method: "GET",
            url: "/user/test"
        });
        axiosFetch({
            axiosInstance: axiosPrivate,
            method: "GET",
            url: "/user/test"
        });
        axiosFetch({
            axiosInstance: axiosPrivate,
            method: "GET",
            url: "/user/test"
        });
        axiosFetch({
            axiosInstance: axiosPrivate,
            method: "GET",
            url: "/user/test"
        });
        
    
    }, [axiosFetch])

    return (
        <Box>            
            {loading && <Loading />}

            {!loading && error && <ErrorAlert message={error} />}

            {!loading && !error && response?.message && <Typography>{response?.message}</Typography>}
        </Box>
    )
}