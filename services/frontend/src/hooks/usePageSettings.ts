import { useNavigate } from "react-router-dom";
// import useAuth from "./useAuth";
import { useEffect } from "react";
import { useAuthStore } from "../store/authStore";
import { Layout } from "../lib/layout/layout";

export interface PageSettingsConfig {
  title: string;
  requireAuth?: boolean;
  showSidebar?: boolean;
  layout?: Layout | null;
}

export const usePageSettings = ({
  title,
  requireAuth = true,
}: PageSettingsConfig) => {
  const { isAuthenticated } = useAuthStore();
  const navigate = useNavigate();

  const isLoggedIn = isAuthenticated();
  useEffect(() => {
    document.title = title;

    if (requireAuth && !isLoggedIn) {
      navigate("/login");
    }
  }, [title, requireAuth, isAuthenticated, navigate]);

  return {};
};
