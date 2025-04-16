import { FC, useEffect } from "react";
import { PageSettings } from "./page";
import RegisterForm from "../components/forms/RegisterForm";
import authAPI from "../api/auth";
import { register } from "module";
import { useNavigate } from "react-router-dom";

class RegisterPageSettings extends PageSettings {}

const RegisterPage: FC = () => {
  const settings = new RegisterPageSettings("Sign up");

  const navigate = useNavigate();

  useEffect(() => {
    settings.setup();
  }, []);

  const handleRegister = async (email: string, password: string) => {
    // const res = await authAPI().register({ email: email, password: password });
    // register(res);
    navigate("/");
  };
  return (
    <div className="flex justify-center items-center min-h-lvh bg-primary-600">
      <div className="bg-primary-500 p-8 rounded-lg shadow-lg w-96">
        <h2 className="text-2xl font-semibold text-muted-100 mb-6 text-center">
          Welcome Back
        </h2>
        <RegisterForm onRegister={handleRegister} />
      </div>
    </div>
  );
};

export default RegisterPage;
