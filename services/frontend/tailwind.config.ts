import type { Config } from "tailwindcss";

const config: Config = {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        primary: "var(--primary-color-1)",
        neutral: "var(--neutral-color-1)",
        accent: "var(--accent-color-1)",
      },
    },
  },
  plugins: [],
};

export default config;
