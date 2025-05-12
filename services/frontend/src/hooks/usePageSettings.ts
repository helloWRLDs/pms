import { useNavigate } from "react-router-dom";
// import useAuth from "./useAuth";
import { useEffect } from "react";
import { useAuthStore } from "../store/authStore";

export interface PageSettingsConfig {
  title: string;
  requireAuth?: boolean;
  showSidebar?: boolean;
}

export const usePageSettings = ({
  title,
  requireAuth = true,
  showSidebar = true,
}: PageSettingsConfig) => {
  const { isAuthenticated } = useAuthStore();
  const navigate = useNavigate();

  const isLoggedIn = isAuthenticated();
  useEffect(() => {
    document.title = title;

    if (requireAuth && !isLoggedIn) {
      navigate("/login");
    }
  }, [title, requireAuth, isLoggedIn, navigate]);

  return {
    showSidebar,
  };
};
