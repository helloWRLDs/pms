import { useSelector, useDispatch } from "react-redux";
import { RootState } from "../store/authStore";
import { login, logout } from "../store/authSlice";
import { isTokenValid } from "../utils/jwt";
import { useMemo } from "react";

const useAuth = () => {
  const dispatch = useDispatch();

  const { access_token, user } = useSelector((state: RootState) => state);

  const isAuthenticated = useMemo(
    () => isTokenValid(access_token || ""),
    [access_token]
  );

  const handleLogin = (auth: RootState) => {
    dispatch(login(auth));
  };

  const handleLogout = () => {
    dispatch(logout());
  };

  return {
    access_token,
    user,
    isAuthenticated,
    login: handleLogin,
    logout: handleLogout,
  };
};

export default useAuth;
