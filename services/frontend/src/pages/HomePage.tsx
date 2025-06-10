import { FC } from "react";
import { usePageSettings } from "../hooks/usePageSettings";
import { TbDeviceAnalytics, TbBrandSpeedtest } from "react-icons/tb";
import {
  MdOutlineSecurity,
  MdOutlineIntegrationInstructions,
} from "react-icons/md";
import { RiTeamFill } from "react-icons/ri";
import { BsKanbanFill } from "react-icons/bs";
import { useNavigate } from "react-router-dom";
import { useAuthStore } from "../store/authStore";

const HomePage: FC = () => {
  usePageSettings({ requireAuth: false, title: "Home" });
  const navigate = useNavigate();
  const { isAuthenticated } = useAuthStore();

  const handleGetStarted = () => {
    if (isAuthenticated()) {
      navigate("/companies");
    } else {
      navigate("/register");
    }
  };

  const handleLogin = () => {
    navigate("/login");
  };

  const FEATURES = [
    {
      title: "Advanced Analytics",
      body: "Real-time insights and data visualization to track project progress and team performance.",
      icon: TbDeviceAnalytics,
      color: "bg-blue-500",
    },
    {
      title: "Enterprise Security",
      body: "Bank-grade encryption and security measures to protect your business data.",
      icon: MdOutlineSecurity,
      color: "bg-green-500",
    },
    {
      title: "Team Collaboration",
      body: "Seamless communication and collaboration tools to keep your team aligned.",
      icon: RiTeamFill,
      color: "bg-purple-500",
    },
    {
      title: "Agile Kanban Boards",
      body: "Flexible task management with customizable workflows and automation.",
      icon: BsKanbanFill,
      color: "bg-orange-500",
    },
    {
      title: "Fast Performance",
      body: "Lightning-fast response times and optimized workflows for maximum efficiency.",
      icon: TbBrandSpeedtest,
      color: "bg-red-500",
    },
    {
      title: "Easy Integration",
      body: "Connect with your favorite tools through our extensive API and webhooks.",
      icon: MdOutlineIntegrationInstructions,
      color: "bg-indigo-500",
    },
  ];

  return (
    <div className="min-h-screen bg-gradient-to-b from-gray-50 to-white">
      {/* Hero Section */}
      <div className="relative overflow-hidden bg-transparent min-h-[85vh]">
        {/* Video Background */}
        <div className="absolute inset-0 w-full h-full">
          <div className="absolute inset-0 bg-black/50 z-10"></div>
          <video
            autoPlay
            loop
            muted
            playsInline
            className="w-full h-full object-cover"
          >
            <source
              src={
                "https://videos.pexels.com/video-files/3254066/3254066-uhd_2560_1440_25fps.mp4"
              }
              type="video/mp4"
            />
            Your browser does not support the video tag.
          </video>
        </div>

        <div className="relative z-20">
          <div className="max-w-7xl mx-auto">
            <div className="relative pb-8 sm:pb-16 md:pb-20 lg:pb-28 xl:pb-32 min-h-[85vh] flex items-center">
              <main className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
                <div className="text-center">
                  {/* <img
                    src={logo}
                    alt="Taskflow Logo"
                    className="w-24 h-24 mx-auto mb-8"
                  /> */}
                  <h1 className="text-4xl tracking-tight font-extrabold text-white sm:text-5xl md:text-6xl">
                    <span className="block text-accent-200">
                      Streamline Your Workflow
                    </span>
                    <span className="block text-white">Boost Productivity</span>
                  </h1>
                  <p className="mt-3 text-base text-gray-100 sm:mt-5 sm:text-lg sm:max-w-xl sm:mx-auto md:mt-5 md:text-xl">
                    Transform your team's efficiency with our intuitive project
                    management platform. Track tasks, collaborate seamlessly,
                    and deliver results faster.
                  </p>
                  <div className="mt-5 sm:mt-8 flex justify-center gap-4">
                    <button
                      onClick={handleGetStarted}
                      className="px-8 py-3 text-base cursor-pointer font-medium rounded-lg text-white bg-accent-500 hover:bg-accent-600 transition-all duration-200 transform hover:scale-105 shadow-lg hover:shadow-xl"
                    >
                      Get Started
                    </button>
                    <button
                      onClick={handleLogin}
                      className="px-8 py-3 text-base cursor-pointer font-medium rounded-lg text-white bg-transparent border-2 border-white hover:bg-white/10 transition-all duration-200 transform hover:scale-105 shadow-lg hover:shadow-xl"
                    >
                      Sign In
                    </button>
                  </div>
                </div>
              </main>
            </div>
          </div>

          {/* Scroll Indicator */}
          <div className="absolute bottom-10 left-1/2 transform -translate-x-1/2 animate-bounce">
            <div className="w-6 h-10 border-2 border-white rounded-full flex justify-center">
              <div className="w-1 h-3 bg-white rounded-full mt-2 animate-scroll"></div>
            </div>
          </div>
        </div>
      </div>

      {/* Features Section */}
      <div className="py-24 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center">
            <h2 className="text-3xl font-extrabold text-gray-900 sm:text-4xl">
              Powerful Features for Modern Teams
            </h2>
            <p className="mt-4 text-xl text-gray-500">
              Everything you need to manage projects effectively in one place
            </p>
          </div>

          <div className="mt-20">
            <div className="grid grid-cols-1 gap-8 sm:grid-cols-2 lg:grid-cols-3">
              {FEATURES.map((feature) => (
                <div
                  key={feature.title}
                  className="relative group bg-white p-6 rounded-2xl shadow-lg hover:shadow-xl transition-all duration-300 transform hover:-translate-y-1"
                >
                  <div
                    className={`${feature.color} rounded-lg p-3 inline-flex`}
                  >
                    <feature.icon className="h-6 w-6 text-white" />
                  </div>
                  <h3 className="mt-4 text-lg font-semibold text-gray-900">
                    {feature.title}
                  </h3>
                  <p className="mt-2 text-gray-500">{feature.body}</p>
                  <div className="absolute bottom-0 left-0 h-1 w-0 group-hover:w-full bg-accent-500 transition-all duration-300 rounded-b-2xl"></div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>

      {/* CTA Section */}
      <div className="bg-secondary-700">
        <div className="max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:py-16 lg:px-8 lg:flex lg:items-center lg:justify-between">
          <div>
            <h2 className="text-3xl font-extrabold tracking-tight text-white sm:text-4xl">
              <span className="block">Ready to get started?</span>
              <span className="block text-accent-500">
                Join thousands of successful teams today.
              </span>
            </h2>
            <p className="mt-4 text-lg text-gray-300">
              Start your free trial now. No credit card required.
            </p>
          </div>
          <div className="mt-8 flex lg:mt-0 lg:flex-shrink-0 gap-4">
            <button
              onClick={handleGetStarted}
              className="px-8 py-3 text-base font-medium rounded-lg text-gray-900 bg-white hover:bg-gray-100 transition-all duration-200 transform hover:scale-105 shadow-lg hover:shadow-xl"
            >
              Get Started
            </button>
            <button
              onClick={handleLogin}
              className="px-8 py-3 text-base font-medium rounded-lg text-white bg-accent-500 hover:bg-accent-600 transition-all duration-200 transform hover:scale-105 shadow-lg hover:shadow-xl"
            >
              Learn More
            </button>
          </div>
        </div>

        {/* Footer Links */}
        <div className="max-w-7xl mx-auto pt-8 pb-12 px-4 sm:px-6 lg:px-8 border-t border-gray-800">
          <div className="grid grid-cols-2 md:grid-cols-4 gap-8">
            <div>
              <h3 className="text-sm font-semibold text-gray-400 tracking-wider uppercase">
                Product
              </h3>
              <ul className="mt-4 space-y-4">
                <li>
                  <a
                    href="#"
                    className="text-base text-gray-300 hover:text-white"
                  >
                    Features
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="text-base text-gray-300 hover:text-white"
                  >
                    Pricing
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="text-base text-gray-300 hover:text-white"
                  >
                    Security
                  </a>
                </li>
              </ul>
            </div>
            <div>
              <h3 className="text-sm font-semibold text-gray-400 tracking-wider uppercase">
                Company
              </h3>
              <ul className="mt-4 space-y-4">
                <li>
                  <a
                    href="#"
                    className="text-base text-gray-300 hover:text-white"
                  >
                    About
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="text-base text-gray-300 hover:text-white"
                  >
                    Blog
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="text-base text-gray-300 hover:text-white"
                  >
                    Careers
                  </a>
                </li>
              </ul>
            </div>
            <div>
              <h3 className="text-sm font-semibold text-gray-400 tracking-wider uppercase">
                Resources
              </h3>
              <ul className="mt-4 space-y-4">
                <li>
                  <a
                    href="#"
                    className="text-base text-gray-300 hover:text-white"
                  >
                    Documentation
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="text-base text-gray-300 hover:text-white"
                  >
                    Help Center
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="text-base text-gray-300 hover:text-white"
                  >
                    API Reference
                  </a>
                </li>
              </ul>
            </div>
            <div>
              <h3 className="text-sm font-semibold text-gray-400 tracking-wider uppercase">
                Legal
              </h3>
              <ul className="mt-4 space-y-4">
                <li>
                  <a
                    href="#"
                    className="text-base text-gray-300 hover:text-white"
                  >
                    Privacy
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="text-base text-gray-300 hover:text-white"
                  >
                    Terms
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="text-base text-gray-300 hover:text-white"
                  >
                    Cookie Policy
                  </a>
                </li>
              </ul>
            </div>
          </div>
          <div className="mt-8 pt-8 border-t border-gray-800 text-center">
            <p className="text-base text-gray-400">
              &copy; 2024 Taskflow. All rights reserved.
            </p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default HomePage;
