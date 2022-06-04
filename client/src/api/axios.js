import axios from "axios";

const URL = "http://localhost:4000/api"

export const axiosPublic = axios.create({
    baseURL: URL,
    withCredentials: true
})

export const axiosPrivate = axios.create({
    baseURL: URL,
    headers: { "Content-Type": "application/json" },
    withCredentials: true
})

let refreshing_token = null;

axiosPrivate.interceptors.request.use(
    config => {
        if (!config.headers["Authorization"]) {
            config.headers["Authorization"] = `Bearer ${window.accessToken}`
        }
        return config
    }, (error) => Promise.reject(error)
)

axiosPrivate.interceptors.response.use(
    response => response,
    async (error) => {
        const prevRequest = error?.config;
        if (error?.response?.status === 401 && !prevRequest?.sent) {
            prevRequest.sent = true;
            refreshing_token = refreshing_token ? refreshing_token : refreshToken();
            let res = await refreshing_token;
            refreshing_token = null;
            window.accessToken = res
            prevRequest.headers["Authorization"] = `Bearer ${res}`;
            return axiosPrivate(prevRequest);
        }
        return Promise.reject(error);
    }
);

const refreshToken = async () => {
    const response = await axiosPublic.get("/auth/refresh");

    return response.data.accessToken;
}