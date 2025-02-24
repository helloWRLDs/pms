import { FC, useState } from "react";
import { MdEmail, MdLock, MdVisibility, MdVisibilityOff } from "react-icons/md";
import { toast } from "react-toastify";
import { toastOpts } from "../../utils/toast";

interface Props {
  onLogin: (email: string, password: string) => void;
}

const LoginForm: FC<Props> = (props) => {
  const [password, setPassword] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [showPassword, setShowPassword] = useState<boolean>(false);

  const isValidEmail = (email: string): boolean => {
    return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
  };

  // ✅ Handle Form Submission
  const handleSubmit = () => {
    if (!email) {
      toast.error("Email is required!", toastOpts);
      return;
    }
    if (!isValidEmail(email)) {
      toast.error("Invalid email format!", toastOpts);
      return;
    }

    // ✅ Call `onLogin` prop
    props.onLogin(email, password);
    toast.success("Login successful!", toastOpts);
  };

  return (
    <div>
      {/* Email Input */}
      <div className="mb-4">
        <label htmlFor="email" className="relative block">
          <MdEmail
            size={20}
            className="absolute left-3 top-1/2 transform -translate-y-1/2 text-[var(--color-neutral-400)]"
          />
          <input
            type="email"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="Email"
            className="w-full h-12 pl-10 pr-4 border border-[var(--color-neutral-400)] rounded-lg bg-[var(--color-primary-400)] text-[var(--color-neutral-100)] focus:ring-2 focus:ring-[var(--color-primary-200)] focus:border-[var(--color-primary-200)] outline-none transition"
          />
        </label>
      </div>

      {/* Password Input */}
      <div className="mb-4">
        <label htmlFor="password" className="relative block">
          <MdLock
            size={20}
            className="absolute left-3 top-1/2 transform -translate-y-1/2 text-[var(--color-neutral-400)]"
          />
          <input
            type={showPassword ? "text" : "password"}
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Password"
            className="w-full h-12 pl-10 pr-10 border border-[var(--color-neutral-400)] rounded-lg bg-[var(--color-primary-400)] text-[var(--color-neutral-100)] focus:ring-2 focus:ring-[var(--color-primary-200)] focus:border-[var(--color-primary-200)] outline-none transition"
          />
          {/* Toggle Password Visibility */}
          <button
            type="button"
            className="absolute right-3 top-1/2 transform -translate-y-1/2 text-neutral-400 cursor-pointer"
            onClick={() => setShowPassword(!showPassword)}
          >
            {showPassword ? (
              <MdVisibility size={22} />
            ) : (
              <MdVisibilityOff size={22} />
            )}
          </button>
        </label>
      </div>

      {/* Submit Button */}
      <button
        onClick={handleSubmit}
        className="w-full bg-primary-300 hover:bg-primary-200 text-neutral-100 cursor-pointer font-semibold py-3 rounded-lg transition"
      >
        Sign In
      </button>

      {/* Signup Link */}
      <p className="text-neutral-400 text-sm mt-4 text-center">
        Don't have an account?{" "}
        <a href="#" className="text-[var(--color-neutral-200)] hover:underline">
          Sign up
        </a>
      </p>
    </div>
  );
};

export default LoginForm;
