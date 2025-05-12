import { FC, useState } from "react";
import { MdVisibility, MdVisibilityOff } from "react-icons/md";
import { toast } from "react-toastify";
import { toastOpts } from "../../lib/utils/toast";

interface Props {
  onLogin: (email: string, password: string) => Promise<void>;
}

const LoginForm: FC<Props> = (props) => {
  const [password, setPassword] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [showPassword, setShowPassword] = useState<boolean>(false);

  const isValidEmail = (email: string): boolean => {
    return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
  };

  const handleSubmit = async () => {
    if (!email) {
      toast.error("Email is required!", toastOpts);
      return;
    }
    if (!isValidEmail(email)) {
      toast.error("Invalid email format!", toastOpts);
      return;
    }

    try {
      await props.onLogin(email, password);
      toast.success("Login successful!", toastOpts);
    } catch (e) {
      console.log(e);
    }
  };

  return (
    <div>
      {/* Email Input */}
      <div className="relative z-0 mb-4">
        <input
          type="text"
          value={email}
          id="form-email"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          required={true}
          onChange={(e) => setEmail(e.target.value)}
        />
        <label
          htmlFor="form-email"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Email
        </label>
      </div>

      <div className="relative z-0 mb-4">
        <input
          autoComplete="off"
          type={showPassword ? "text" : "password"}
          id="form-password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required={true}
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
        />
        <label
          htmlFor="form-password"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Password
        </label>
        <button
          type="button"
          className="absolute right-3 top-1/2 transform -translate-y-1/2 text-neutral-400 cursor-pointer hover:text-accent-500 transition-all ease-in-out duration-300"
          onClick={() => setShowPassword(!showPassword)}
        >
          {showPassword ? (
            <MdVisibility size={22} />
          ) : (
            <MdVisibilityOff size={22} />
          )}
        </button>
      </div>

      {/* Submit Button */}
      <button
        onClick={handleSubmit}
        className="w-full bg-primary-300 hover:bg-accent-200 hover:text-primary-100 text-neutral-100 cursor-pointer font-semibold py-3 rounded-lg transition-all ease-in-out duration-300"
      >
        Sign In
      </button>

      {/* Signup Link */}
      <p className="text-neutral-400 text-sm mt-4 text-center">
        Don't have an account?{" "}
        <a
          href="/register"
          className="text-[var(--color-neutral-200)] hover:underline hover:text-accent-500 transition-all ease-in-out duration-300"
        >
          Sign up
        </a>
      </p>
    </div>
  );
};

export default LoginForm;
