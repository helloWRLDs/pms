import { ComponentProps, FC, useEffect } from "react";
import useAuth from "../../hooks/useAuth";
import { Menu, MenuButton, MenuItem, MenuItems } from "@headlessui/react";
import { FaCircleUser } from "react-icons/fa6";
import { useNavigate } from "react-router-dom";

interface HeaderProps extends ComponentProps<"div"> {
  logoURL: string;
}

const NAV_ITEMS = [
  { name: "Home", link: "/" },
  { name: "Dashboard", link: "/dashboard" },
  { name: "Team", link: "/team" },
  { name: "Backlogs", link: "/backlogs" },
];

export const Header: FC<HeaderProps> = ({ logoURL, className }) => {
  const { isAuthenticated, logout } = useAuth();

  const navigate = useNavigate();

  useEffect(() => {
    console.log(isAuthenticated);
  }, []);

  return (
    <div className={`bg-primary-400 shadow-md ${className}`}>
      <header className="flex items-center px-8">
        {/* Logo */}
        <a href="/" className="flex items-center">
          <img src={logoURL} alt="logo" width={90} />
        </a>

        {/* Navigation */}
        <nav className="flex ml-12 mr-auto space-x-8 text-soft-100 ">
          {NAV_ITEMS.map((item, i) => (
            <a
              href={item.link}
              key={i}
              className="w-auto rounded-md px-3 py-1 transition ease-in-out duration-300 relative group"
            >
              {item.name}
              <span className="absolute left-0 -bottom-0 w-0 h-0.5 bg-accent-500 transition-all group-hover:w-full"></span>
            </a>
          ))}
        </nav>

        {/* Auth Status */}
        <div className="ml-auto text-soft-200 font-semibold text-sm">
          {isAuthenticated ? (
            <Menu as="div" className="relative inline-block text-left">
              <div>
                <MenuButton className="inline-flex items-center gap-2 rounded-md  px-4 py-2 text-gray-900 text-sm font-medium  transition cursor-pointer">
                  <FaCircleUser className="h-10 w-10" color="white" />
                </MenuButton>
              </div>

              <MenuItems className="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white ring-1 ring-black/5 shadow-lg focus:outline-none">
                <div className="py-1">
                  <MenuItem>
                    {({ active }) => (
                      <a
                        href="/profile"
                        className={`block px-4 py-2 text-sm ${
                          active
                            ? "bg-accent-400 text-gray-900"
                            : "text-gray-700"
                        }`}
                      >
                        Profile
                      </a>
                    )}
                  </MenuItem>

                  <form action="#" method="POST">
                    <MenuItem>
                      {({ active }) => (
                        <button
                          type="submit"
                          onClick={() => {
                            logout();
                            navigate("/login");
                          }}
                          className={`block w-full px-4 py-2 text-left text-sm ${
                            active
                              ? "bg-accent-400 text-gray-900"
                              : "text-gray-700"
                          }`}
                        >
                          Sign out
                        </button>
                      )}
                    </MenuItem>
                  </form>
                </div>
              </MenuItems>
            </Menu>
          ) : (
            <button
              onClick={() => {
                navigate("/login");
              }}
              className="text-primary-700 cursor-pointer rounded-md py-2 px-4 bg-accent-700 hover:bg-accent-200 active:bg-accent-500 transition"
            >
              Login
            </button>
          )}
        </div>
      </header>
    </div>
  );
};
