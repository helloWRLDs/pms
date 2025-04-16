import { useEffect } from "react";
import useAuth from "./hooks/useAuth";
import { Header } from "./components/ui/Header";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import HomePage from "./pages/HomePage.tsx";
import DashboardPage from "./pages/DashboardPage.tsx";
import LoginPage from "./pages/LoginPage.tsx";
import RegisterPage from "./pages/RegisterPage.tsx";
import ChatPage from "./pages/ChatPage.tsx";
import ProfilePage from "./pages/ProfilePage.tsx";
import ProjectsPage from "./pages/ProjectsPage.tsx";
import Footer from "./components/ui/Footer.tsx";

function App() {
  const { isAuthenticated, user } = useAuth();
  useEffect(() => {
    console.log(isAuthenticated);
    console.log(user);
  }, []);

  return (
    <>
      <BrowserRouter>
        <Header logoURL="./src/assets/logo.png" />
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/dashboard" element={<DashboardPage />} />

          {/* authorization */}
          <Route path="/login" element={<LoginPage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/profile" element={<ProfilePage />} />

          {/* projects */}
          <Route path="/projects" element={<ProjectsPage />} />
          <Route path="/projects/:id" />
          <Route path="/projects/:projectID/backlog" />
          <Route path="/projects/:projectID/sprints/:sprintID" />
          <Route path="/projects/:projectID/sprints/:sprintID/tasks/:taskID" />

          {/* admin */}
          <Route path="/users" />

          <Route path="/test" element={<ChatPage />} />
        </Routes>
      </BrowserRouter>
      {/* <Footer /> */}
    </>
  );
}

export default App;
