import { Navigate, Outlet } from "react-router-dom";
import { NotFound } from "../pages/NotFound";
import { useAuth } from "./useAuth";

export const RequireAuth = ({allowedRoles}) => {
    const { auth } = useAuth();

    return (
        allowedRoles?.includes(auth?.role)
            ? <Outlet />
            : auth?.username
                ? <NotFound />
                : <Navigate to="/signin" />
    );
}