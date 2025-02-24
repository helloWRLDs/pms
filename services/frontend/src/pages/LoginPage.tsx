import { FC } from "react";
import LoginForm from "../components/forms/LoginForm";

const LoginPage: FC = () => {
  return (
    <div className="flex justify-center items-center min-h-lvh bg-primary-600">
      <div className="bg-primary-500 p-8 rounded-lg shadow-lg w-96">
        <h2 className="text-2xl font-semibold text-neutral-100 mb-6 text-center">
          Welcome Back
        </h2>
        <LoginForm
          onLogin={() => {
            console.log("logged in!");
          }}
        />
      </div>
    </div>
  );
};

export default LoginPage;
