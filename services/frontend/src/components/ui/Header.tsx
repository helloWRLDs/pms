import { ComponentProps, FC, useEffect } from "react";
import useAuth from "../../hooks/useAuth";

interface HeaderProps extends ComponentProps<"div"> {
  logoURL: string;
}

export const Header: FC<HeaderProps> = ({ logoURL, className }) => {
  const { isAuthenticated } = useAuth();

  useEffect(() => {
    console.log(isAuthenticated);
  }, []);

  return (
    <div className={`bg-primary-500 shadow-md ${className}`}>
      <header className="container mx-auto flex items-center py-4 px-6">
        {/* ✅ Logo */}
        <a href="/" className="mr-auto">
          <img src={logoURL} alt="logo" width={90} />
        </a>

        {/* ✅ Navigation Links */}
        <ul className="hidden md:flex space-x-8 text-neutral-100">
          <a href="/" className="hover:text-neutral-300 transition">
            <li>Home</li>
          </a>
          <a href="/dashboard" className="hover:text-neutral-300 transition">
            <li>Dashboard</li>
          </a>
          <a href="/team" className="hover:text-neutral-300 transition">
            <li>Team</li>
          </a>
          <a href="/backlogs" className="hover:text-neutral-300 transition">
            <li>Backlogs</li>
          </a>
        </ul>

        {/* ✅ Login Status */}
        <div className="ml-auto text-neutral-200 font-semibold">
          {isAuthenticated ? (
            <span className="text-green-400">Logged in</span>
          ) : (
            <a href="/login" className="hover:text-neutral-300 transition">
              Login
            </a>
          )}
        </div>
      </header>
    </div>
  );
};
