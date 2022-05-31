import { useCallback, useEffect, useState } from "react";
import { useNavigate, useLocation } from "react-router-dom";

export const useAxiosFunction = () => {
    const [response, setResponse] = useState(null);
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);
    const [controller, setController] = useState(null);

    const navigate = useNavigate();
    const location = useLocation();

    const axiosFetch = useCallback(async (configObj) => {
        const {
            axiosInstance,
            method,
            url,
            requestConfig = {}
        } = configObj

        try {
            setLoading(true);
            const ctrl = new AbortController();
            setController(ctrl);
            const res = await axiosInstance.request({
                ...requestConfig,
                method,
                url,
                signal: ctrl.signal
            })
            setResponse(res.data);
            setError("");
        } catch (err) {
            setError(err?.response?.data?.message);
            if (err?.response?.status === 401 || err?.response?.status === 403) navigate("/signin", { state: { from: location }, replace: true} )
        } finally {
            setInterval(() => {
                setLoading(false);
            }, 100); // remove later
        }
    }, [location, navigate])
    useEffect(() => {
        return () => controller && controller.abort();
    }, [controller])

    return [response, error, loading, axiosFetch];
}