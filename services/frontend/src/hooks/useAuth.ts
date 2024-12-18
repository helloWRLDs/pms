import { useSelector, useDispatch } from "react-redux";
import { RootState } from "../store/authStore";
import { login, logout } from "../store/authSlice";
import { isTokenValid } from "../utils/jwt";
import { useMemo } from "react";

const useAuth = () => {
  const dispatch = useDispatch();

  const { token, user } = useSelector((state: RootState) => state);

  const isAuthenticated = useMemo(() => isTokenValid(token || ""), [token]);

  const handleLogin = (auth: RootState) => {
    dispatch(login(auth));
  };

  const handleLogout = () => {
    dispatch(logout());
  };

  return {
    token,
    user,
    isAuthenticated,
    login: handleLogin,
    logout: handleLogout,
  };
};

export default useAuth;
