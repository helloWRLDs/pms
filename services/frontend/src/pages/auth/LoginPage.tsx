import { FC, useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";
import { usePageSettings } from "../../hooks/usePageSettings";
import authAPI from "../../api/authAPI";
import { useAuthStore } from "../../store/authStore";
import LoginForm from "../../components/forms/LoginForm";
import GoogleOAuthButton from "../../components/auth/GoogleOAuthButton";
import { HiLockClosed } from "react-icons/hi";

const LoginPage: FC = () => {
  usePageSettings({
    requireAuth: false,
    title: "Sign in",
  });

  const { isAuthenticated, setAuth } = useAuthStore();
  const navigate = useNavigate();

  useEffect(() => {
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
    navigate("/");
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-primary-600 via-primary-700 to-primary-800 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md w-full space-y-8">
        <div className="text-center">
          <div className="mx-auto h-16 w-16 bg-accent-500 rounded-full flex items-center justify-center">
            <HiLockClosed className="h-8 w-8 text-white" />
          </div>
          <h2 className="mt-6 text-3xl font-bold text-white">Welcome back</h2>
          <p className="mt-2 text-sm text-primary-200">
            Sign in to your account to continue
          </p>
        </div>

        <div className="bg-white rounded-xl shadow-2xl p-8 space-y-6">
          <div className="space-y-4">
            <GoogleOAuthButton />

            <div className="relative">
              <div className="absolute inset-0 flex items-center">
                <div className="w-full border-t border-gray-300" />
              </div>
              <div className="relative flex justify-center text-sm">
                <span className="px-2 bg-white text-gray-500">
                  Or continue with email
                </span>
              </div>
            </div>
          </div>

          <LoginForm onLogin={handleLogin} />

          <div className="text-center space-y-2">
            <p className="text-sm text-gray-600">
              Don't have an account?{" "}
              <Link
                to="/register"
                className="font-medium text-accent-600 hover:text-accent-500 transition-colors"
              >
                Sign up here
              </Link>
            </p>
            <p className="text-xs text-gray-500">
              <a href="#" className="hover:text-accent-500 transition-colors">
                Forgot your password?
              </a>
            </p>
          </div>
        </div>

        <div className="text-center">
          <p className="text-xs text-primary-300">
            By signing in, you agree to our{" "}
            <a
              href="#"
              className="underline hover:text-white transition-colors"
            >
              Terms of Service
            </a>{" "}
            and{" "}
            <a
              href="#"
              className="underline hover:text-white transition-colors"
            >
              Privacy Policy
            </a>
          </p>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
