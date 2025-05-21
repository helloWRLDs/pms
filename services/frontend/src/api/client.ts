import axios, { AxiosError, AxiosInstance } from "axios";
import { toast } from "react-toastify";
import { toastOpts } from "../lib/utils/toast";
import { ErrorResponse } from "../lib/errors";
// import useAuth from "../hooks/useAuth";

export interface APIConfig {
  baseURL: string;
  headers: Record<string, string>;
}

const defaultConfig: APIConfig = {
  baseURL: "https://localhost:8080/api/v1",
  headers: { "Content-Type": "application/json" },
};

const createAPIClient = (config: APIConfig) => {
  const apiConfig: APIConfig = {
    baseURL: config.baseURL || defaultConfig.baseURL,
    headers: config.headers || defaultConfig.headers,
  };
  const apiClient: AxiosInstance = axios.create(apiConfig);

  apiClient.interceptors.response.use(
    (res) => res,
    (error: AxiosError) => {
      if (error.response) {
        const { data, status } = error.response;
        const resBody: ErrorResponse =
          typeof data === "string" ? JSON.parse(data) : data;

        switch (status) {
          case 400:
            toast.warning(resBody?.msg || "Bad request", toastOpts);
            break;
          case 401:
            toast.error("Unauthorized! Please login.", toastOpts);
            break;
          case 404:
            toast.error(resBody?.msg || "Not found", toastOpts);
            break;
          case 500:
            toast.error(
              `Server error: ${resBody.msg}! Please try again later.`,
              toastOpts
            );
            break;
          default:
            toast.error("An unexpected error occurred!", toastOpts);
        }
      } else {
        toast.error("Network error! Please check your connection.", toastOpts);
      }
      return Promise.reject(error);
    }
  );

  return apiClient;
};

export default createAPIClient;
