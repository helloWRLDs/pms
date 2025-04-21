import { useNavigate } from "react-router-dom";
import useAuth from "./useAuth";
import { useEffect } from "react";

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
  const { isAuthenticated } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    document.title = title;

    if (requireAuth && !isAuthenticated) {
      navigate("/login");
    }
  }, [title, requireAuth, isAuthenticated, navigate]);

  return {
    showSidebar,
  };
};
