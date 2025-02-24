import { useEffect } from "react";
import useAuth from "./hooks/useAuth";
import { Header } from "./components/ui/Header";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import HomePage from "./pages/HomePage.tsx";
import DashboardPage from "./pages/DashboardPage.tsx";
import LoginPage from "./pages/LoginPage.tsx";
import RegisterPage from "./pages/RegisterPage.tsx";

function App() {
  const { isAuthenticated } = useAuth();
  useEffect(() => {
    console.log(isAuthenticated);
  }, []);

  return (
    <>
      <Header logoURL="./src/assets/logo.png" />
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/dashboard" element={<DashboardPage />} />
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
