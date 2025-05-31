import { FC, useState } from "react";
import { errorToast } from "../../lib/utils/toast";

interface Props {
  onRegister: (
    email: string,
    password: string,
    firstName: string,
    lastName: string
  ) => Promise<void>;
}

const RegisterForm: FC<Props> = (props) => {
  const [firstName, setFirstName] = useState<string>("");
  const [lastName, setLastName] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [confirmPassword, setConfirmPassword] = useState<string>("");

  const handleSubmit = async () => {
    try {
      await props.onRegister(email, password, firstName, lastName);
      // toast.success("Register successful!", toastOpts);
    } catch (e) {
      console.log(e);
      errorToast("Register failed!");
    }
  };

  return (
    <div>
      <div className="relative z-0 mb-4">
        <input
          type="text"
          value={firstName}
          id="form-name"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          required={true}
          autoComplete="off"
          onChange={(e) => setFirstName(e.target.value)}
        />
        <label
          htmlFor="form-name"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          First Name
        </label>
      </div>

      <div className="relative z-0 mb-4">
        <input
          type="text"
          value={lastName}
          id="form-name"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          required={true}
          autoComplete="off"
          onChange={(e) => setLastName(e.target.value)}
        />
        <label
          htmlFor="form-name"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Last Name
        </label>
      </div>

      <div className="relative z-0 mb-4">
        <input
          type="text"
          value={email}
          id="form-email"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          required={true}
          autoComplete="off"
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
          type="password"
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
      </div>

      <div className="relative z-0 mb-4">
        <input
          type="password"
          id="form-confirm-password"
          value={confirmPassword}
          onChange={(e) => setConfirmPassword(e.target.value)}
          onBlur={() => {
            if (password !== confirmPassword) {
              errorToast("Passwords do not match!");
            }
          }}
          required={true}
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
        />
        <label
          htmlFor="form-password"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Confirm Password
        </label>
      </div>

      <button
        onClick={handleSubmit}
        className="w-full bg-primary-300 hover:bg-accent-200 hover:text-primary-100 text-neutral-100 cursor-pointer font-semibold py-3 rounded-lg transition-all ease-in-out duration-300"
      >
        Sign Up
      </button>

      <p className="text-neutral-400 text-sm mt-4 text-center">
        Already have an account?{" "}
        <a
          href="/login"
          className="text-[var(--color-neutral-200)] hover:underline hover:text-accent-500 transition-all ease-in-out duration-300"
        >
          Sign in
        </a>
      </p>
    </div>
  );
};

export default RegisterForm;
