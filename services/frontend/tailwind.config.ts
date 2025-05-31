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
      animation: {
        blob: "blob 7s infinite",
        bounce: "bounce 1s infinite",
        scroll: "scroll 2s infinite",
      },
      keyframes: {
        blob: {
          "0%": {
            transform: "translate(0px, 0px) scale(1)",
          },
          "33%": {
            transform: "translate(30px, -50px) scale(1.1)",
          },
          "66%": {
            transform: "translate(-20px, 20px) scale(0.9)",
          },
          "100%": {
            transform: "translate(0px, 0px) scale(1)",
          },
        },
        bounce: {
          "0%, 100%": {
            transform: "translateY(-25%)",
            animationTimingFunction: "cubic-bezier(0.8, 0, 1, 1)",
          },
          "50%": {
            transform: "translateY(0)",
            animationTimingFunction: "cubic-bezier(0, 0, 0.2, 1)",
          },
        },
        scroll: {
          "0%": {
            transform: "translateY(0)",
            opacity: "0",
          },
          "30%": {
            opacity: "1",
          },
          "100%": {
            transform: "translateY(0.5rem)",
            opacity: "0",
          },
        },
      },
    },
  },
  plugins: [],
};

export default config;
