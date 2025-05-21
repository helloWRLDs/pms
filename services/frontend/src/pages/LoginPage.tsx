import { FC, useEffect } from "react";
import LoginForm from "../components/forms/LoginForm";
import { useNavigate } from "react-router-dom";
import { usePageSettings } from "../hooks/usePageSettings";
import authAPI from "../api/auth";
import { useAuthStore } from "../store/authStore";

const LoginPage: FC = () => {
  usePageSettings({
    requireAuth: false,
    title: "Sign in",
  });

  // const { isAuthenticated, login } = useAuth();
  const { isAuthenticated, setAuth } = useAuthStore();
  const navigate = useNavigate();

  useEffect(() => {
    // console.log(isAuthenticated);
    if (isAuthenticated()) {
      navigate("/");
    }
  }, []);

  const handleLogin = async (email: string, password: string) => {
    const res = await authAPI.login({ email: email, password: password });
    if (!res) {
      return;
    }
    setAuth(res);
    // const res = await authAPI().login({ email: email, password: password });
    // login(res);
    navigate("/");
  };

  return (
    <div className="flex justify-center min-h-lvh bg-primary-600">
      <div className="bg-primary-500 p-8 rounded-lg shadow-lg w-96 mt-13 h-fit">
        <h2 className="text-2xl font-semibold text-muted-100 mb-6 text-center">
          Welcome Back
        </h2>
        <LoginForm onLogin={handleLogin} />
      </div>
    </div>
  );
};

export default LoginPage;
