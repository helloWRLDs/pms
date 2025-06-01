import { FC, useState } from "react";
import { MdVisibility, MdVisibilityOff } from "react-icons/md";
import { HiMail, HiLockClosed } from "react-icons/hi";
import { toast } from "react-toastify";
import { toastOpts } from "../../lib/utils/toast";

interface Props {
  onLogin: (email: string, password: string) => Promise<void>;
}

const LoginForm: FC<Props> = (props) => {
  const [password, setPassword] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [showPassword, setShowPassword] = useState<boolean>(false);
  const [loading, setLoading] = useState<boolean>(false);

  const isValidEmail = (email: string): boolean => {
    return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!email) {
      toast.error("Email is required!", toastOpts);
      return;
    }
    if (!isValidEmail(email)) {
      toast.error("Invalid email format!", toastOpts);
      return;
    }
    if (!password) {
      toast.error("Password is required!", toastOpts);
      return;
    }

    setLoading(true);
    try {
      await props.onLogin(email, password);
      toast.success("Login successful!", toastOpts);
    } catch (e) {
      toast.error("Login failed. Please check your credentials.", toastOpts);
      console.log(e);
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      {/* Email Input */}
      <div>
        <label
          htmlFor="email"
          className="block text-sm font-medium text-gray-700 mb-1"
        >
          Email address
        </label>
        <div className="relative">
          <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <HiMail className="h-5 w-5 text-gray-400" />
          </div>
          <input
            id="email"
            name="email"
            type="email"
            autoComplete="email"
            required
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-lg 
                     focus:outline-none focus:ring-2 focus:ring-accent-500 focus:border-accent-500
                     placeholder-gray-400 text-gray-900"
            placeholder="Enter your email"
          />
        </div>
      </div>

      {/* Password Input */}
      <div>
        <label
          htmlFor="password"
          className="block text-sm font-medium text-gray-700 mb-1"
        >
          Password
        </label>
        <div className="relative">
          <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <HiLockClosed className="h-5 w-5 text-gray-400" />
          </div>
          <input
            id="password"
            name="password"
            type={showPassword ? "text" : "password"}
            autoComplete="current-password"
            required
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="block w-full pl-10 pr-10 py-2 border border-gray-300 rounded-lg 
                     focus:outline-none focus:ring-2 focus:ring-accent-500 focus:border-accent-500
                     placeholder-gray-400 text-gray-900"
            placeholder="Enter your password"
          />
          <button
            type="button"
            onClick={() => setShowPassword(!showPassword)}
            className="absolute inset-y-0 right-0 pr-3 flex items-center hover:text-gray-600"
          >
            {showPassword ? (
              <MdVisibilityOff className="h-5 w-5 text-gray-400" />
            ) : (
              <MdVisibility className="h-5 w-5 text-gray-400" />
            )}
          </button>
        </div>
      </div>

      {/* Submit Button */}
      <button
        type="submit"
        disabled={loading}
        className="w-full flex justify-center py-2.5 px-4 border border-transparent rounded-lg 
                 shadow-sm text-sm font-medium text-white bg-accent-600 hover:bg-accent-700 
                 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-accent-500
                 disabled:opacity-50 disabled:cursor-not-allowed transition-colors duration-200"
      >
        {loading ? (
          <div className="flex items-center">
            <div className="animate-spin rounded-full h-4 w-4 border-b-2 border-white mr-2"></div>
            Signing in...
          </div>
        ) : (
          "Sign in"
        )}
      </button>
    </form>
  );
};

export default LoginForm;
