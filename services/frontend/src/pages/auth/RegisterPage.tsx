import { FC } from "react";
import RegisterForm from "../../components/forms/RegisterForm";
import { useNavigate } from "react-router-dom";
import { errorToast } from "../../lib/utils/toast";
import { usePageSettings } from "../../hooks/usePageSettings";
import authAPI from "../../api/authAPI";

const RegisterPage: FC = () => {
  usePageSettings({ title: "Sign up", requireAuth: false });

  const navigate = useNavigate();

  const handleRegister = async (
    email: string,
    password: string,
    firstName: string,
    lastName: string
  ) => {
    try {
      await authAPI.register({
        email: email,
        password: password,
        first_name: firstName,
        last_name: lastName,
      });
      navigate("/");
    } catch (e) {
      errorToast("Failed to registre user");
    }
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
