import { useEffect } from "react";
import { Header } from "./components/ui/Header";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import HomePage from "./pages/HomePage.tsx";
import LoginPage from "./pages/LoginPage.tsx";
import RegisterPage from "./pages/RegisterPage.tsx";
import ChatPage from "./pages/ChatPage.tsx";
import ProfilePage from "./pages/ProfilePage.tsx";
import ProjectsPage from "./pages/ProjectsPage.tsx";
import AgileDashboard from "./pages/AgileDashboard.tsx";
import CompaniesPage from "./pages/CompaniesPage.tsx";
import { useAuthStore } from "./store/authStore.ts";
import TestPage from "./pages/TestPage.tsx";
import ProjectOverviewPage from "./pages/ProjectOverviewPage.tsx";
import BacklogPage from "./pages/BacklogPage.tsx";
// import Footer from "./components/ui/Footer.tsx";

function App() {
  const { isAuthenticated, auth } = useAuthStore();
  useEffect(() => {
    console.log(isAuthenticated());
    console.log(auth?.user);
  }, []);

  return (
    <>
      <BrowserRouter>
        <Header logoURL="./src/assets/logo.png" />
        <Routes>
          <Route path="/test" element={<TestPage />} />
          <Route path="/" element={<HomePage />} />
          <Route path="/agile-dashboard" element={<AgileDashboard />} />

          {/* authorization */}
          <Route path="/login" element={<LoginPage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/profile" element={<ProfilePage />} />

          <Route path="/companies" element={<CompaniesPage />} />
          {/* projects */}
          <Route path="/projects" element={<ProjectsPage />} />
          <Route
            path="/companies/:companyID"
            element={<ProjectOverviewPage />}
          />
          <Route
            path="/projects/:projectID/backlog"
            element={<BacklogPage />}
          />
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
