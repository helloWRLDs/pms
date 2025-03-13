import { FC } from "react";
import LoginForm from "../components/forms/LoginForm";
import authAPI from "../api/auth";
import useAuth from "../hooks/useAuth";
import { useNavigate } from "react-router-dom";

const LoginPage: FC = () => {
  const { login } = useAuth();
  const navigate = useNavigate();

  const handleLogin = async (email: string, password: string) => {
    const res = await authAPI().login({ email: email, password: password });
    login(res);
    navigate("/");
  };

  return (
    <div className="flex justify-center items-center min-h-lvh bg-primary-600">
      <div className="bg-primary-500 p-8 rounded-lg shadow-lg w-96">
        <h2 className="text-2xl font-semibold text-neutral-100 mb-6 text-center">
          Welcome Back
        </h2>
        <LoginForm onLogin={handleLogin} />
      </div>
    </div>
  );
};

export default LoginPage;
