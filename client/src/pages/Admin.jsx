import React from "react";
import { useAxiosPrivate } from "../auth/useAxiosPrivate";
import { useAxiosFunction } from "../auth/useAxiosFunction";
import { Loading } from "../components/Loading";
import { Typography } from "@mui/material";
import { Box } from "@mui/material";
import { ErrorAlert } from "../components/ErrorAlert";

export const Admin = () => {
    const axiosPrivate = useAxiosPrivate();
    const [response, error, loading, axiosFetch] = useAxiosFunction();

    React.useEffect(() => {
        axiosFetch({
            axiosInstance: axiosPrivate,
            method: "GET",
            url: "/admin/test"
        });
        
    }, [axiosFetch, axiosPrivate])

    return (
        <Box>            
            {loading && <Loading />}

            {!loading && error && <ErrorAlert message={error} />}

            {!loading && !error && response?.message && <Typography>{response?.message}</Typography>}

            {!loading && !error && !response?.message && <Typography>No content to display</Typography>}
        </Box>
    )
}