import { FC, useEffect } from "react";
import { TbDeviceAnalytics } from "react-icons/tb";
import { MdOutlineSecurity } from "react-icons/md";
import { RiTeamFill } from "react-icons/ri";
import { usePageSettings } from "../hooks/usePageSettings";

const HomePage: FC = () => {
  usePageSettings({ requireAuth: false, title: "Home" });

  const FEATURES = [
    {
      title: "Advanced Analytics",
      body: "Get deep insights into your business performance with our powerful analytics tools.",
      icon: TbDeviceAnalytics,
    },
    {
      title: "Enterprise Security",
      body: "Industry-leading security measures to protect your sensitive business data.",
      icon: MdOutlineSecurity,
    },
    {
      title: "Team Collaboration",
      body: "Seamless collaboration tools to keep your team connected and productive.",
      icon: RiTeamFill,
    },
  ];
  return (
    <div className="min-h-screen bg-white">
      {/* Hero Section */}
      <div className="relative overflow-hidden">
        <div className="max-w-7xl ">
          <div className="relative z-10 pb-8 bg-white sm:pb-16 md:pb-20 lg:max-w-2xl lg:w-full lg:pb-28 xl:pb-32">
            <main className="mt-10 mx-auto max-w-7xl px-4 sm:mt-12 sm:px-6 md:mt-16 lg:mt-20 lg:px-8 xl:mt-28">
              <div className="sm:text-center lg:text-left">
                <h1 className="text-4xl tracking-tight font-extrabold text-gray-900 sm:text-5xl md:text-6xl">
                  <span className="block">Transform your</span>
                  <span className="block text-[rgb(41,43,41)]">
                    business workflow
                  </span>
                </h1>
                <p className="mt-3 text-base text-gray-500 sm:mt-5 sm:text-lg sm:max-w-xl sm:mx-auto md:mt-5 md:text-xl lg:mx-0">
                  Streamline your operations, boost productivity, and achieve
                  better results with our innovative platform. Join thousands of
                  successful businesses already using our solution.
                </p>
                <div className="mt-5 sm:mt-8 sm:flex sm:justify-center lg:justify-start">
                  <div className="rounded-md shadow">
                    <button className="w-full flex items-center justify-center px-8 py-3 border border-transparent text-base font-medium rounded-md text-white bg-[rgb(41,43,41)] hover:bg-[rgb(31,33,31)] md:py-4 md:text-lg md:px-10 cursor-pointer whitespace-nowrap !rounded-button">
                      Get started
                    </button>
                  </div>
                  <div className="mt-3 sm:mt-0 sm:ml-3">
                    <button className="w-full flex items-center justify-center px-8 py-3 border border-transparent text-base font-medium rounded-md text-[rgb(41,43,41)] bg-gray-100 hover:bg-gray-200 md:py-4 md:text-lg md:px-10 cursor-pointer whitespace-nowrap !rounded-button">
                      Live demo
                    </button>
                  </div>
                </div>
              </div>
            </main>
          </div>
        </div>
        <div className="lg:absolute lg:inset-y-0 lg:right-0 lg:w-1/2">
          <img
            className="h-56 w-full object-cover sm:h-72 md:h-96 lg:w-full lg:h-full"
            src="https://readdy.ai/api/search-image?query=modern%20minimalist%20workspace%20with%20sleek%20computer%20setup%2C%20ambient%20lighting%2C%20and%20clean%20design%20elements%20creating%20a%20professional%20and%20innovative%20atmosphere%20perfect%20for%20showcasing%20technology%20solutions&width=800&height=600&seq=9&orientation=landscape"
            alt="Modern workspace"
          />
        </div>
      </div>

      {/* Features Section */}
      <div className="py-12 bg-white">
        <div className="max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="text-center">
            <h2 className="text-3xl font-extrabold text-gray-900 sm:text-4xl">
              Why choose us?
            </h2>
            <p className="mt-4 text-xl text-gray-500">
              Everything you need to manage your business effectively
            </p>
          </div>

          <div className="mt-20">
            <div className="w-full grid grid-cols-1 gap-8 sm:grid-cols-2 lg:grid-cols-3">
              {FEATURES.map((feature) => (
                <div key={feature.title} className="pt-6">
                  <div className="flow-root rounded-lg px-6 pb-8">
                    <div className="-mt-6">
                      <div>
                        <span className="inline-flex items-center justify-center p-3 bg-[rgb(41,43,41)] rounded-md shadow-lg">
                          <feature.icon className="text-accent-300 aspect-square w-10 h-10" />
                        </span>
                      </div>
                    </div>
                    <h3 className="mt-8 text-lg font-medium text-gray-900 tracking-tight">
                      {feature.title}
                    </h3>
                    <p className="mt-5 text-base text-gray-500">
                      {feature.body}
                    </p>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>

      {/* CTA Section */}
      <div className="bg-[rgb(41,43,41)]">
        <div className="max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:py-16 lg:px-8 lg:flex lg:items-center lg:justify-between">
          <h2 className="text-3xl font-extrabold tracking-tight text-white sm:text-4xl">
            <span className="block">Ready to dive in?</span>
            <span className="block text-gray-300">
              Start your free trial today.
            </span>
          </h2>
          <div className="mt-8 flex lg:mt-0 lg:flex-shrink-0">
            <div className="inline-flex rounded-md shadow">
              <button className="inline-flex items-center justify-center px-5 py-3 border border-transparent text-base font-medium rounded-md text-[rgb(41,43,41)] bg-white hover:bg-gray-50 cursor-pointer whitespace-nowrap !rounded-button">
                Get started
              </button>
            </div>
            <div className="ml-3 inline-flex rounded-md shadow">
              <button className="inline-flex items-center justify-center px-5 py-3 border border-transparent text-base font-medium rounded-md text-white bg-[rgb(31,33,31)] hover:bg-[rgb(26,28,26)] cursor-pointer whitespace-nowrap !rounded-button">
                Learn more
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default HomePage;
