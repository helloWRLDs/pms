import { useSelector, useDispatch } from "react-redux";
import { RootState, login, logout } from "../store/authStore";
import { isTokenValid } from "../utils/jwt";
import { useMemo } from "react";
import { AuthData } from "../lib/user";

const useAuth = () => {
  const dispatch = useDispatch();

  const { access_token, user } = useSelector((state: RootState) => state.auth);

  const isAuthenticated = useMemo(
    () => isTokenValid(access_token || ""),
    [access_token]
  );

  const handleLogin = (state: AuthData) => {
    dispatch(login(state));
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
