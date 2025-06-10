import axios, { AxiosError, AxiosInstance } from "axios";
import { LocalStorageKeysMap } from "../lib/consts/localstorage";
import { ErrorResponse } from "../lib/errors";
import { toastOpts } from "../lib/utils/toast";
import { toast } from "react-toastify";
import { AuthData } from "../lib/user/session";
import useMetaCache from "../store/useMetaCache";
import { useAuthStore } from "../store/authStore";

export class API {
  protected baseURL: string;
  protected req: AxiosInstance;

  constructor(baseURL = "http://localhost:8080/api/v1") {
    this.baseURL = baseURL;
    this.req = axios.create({ baseURL: this.baseURL });

    this.req.interceptors.request.use((config) => {
      const { auth } = useAuthStore.getState();
      // const auth: {
      //   state: {
      //     auth: AuthData;
      //   };
      // } | null = localStorage.getItem(LocalStorageKeysMap.AUTH)
      //   ? JSON.parse(localStorage.getItem(LocalStorageKeysMap.AUTH)!)
      //   : null;
      if (auth && auth.access_token) {
        config.headers.Authorization = `Bearer ${auth.access_token}`;
      }
      const metaCache = useMetaCache.getState();
      if (metaCache.metadata.selectedProject?.id) {
        config.headers["X-Project-ID"] = metaCache.metadata.selectedProject.id;
      }
      if (metaCache.metadata.selectedCompany?.id) {
        config.headers["X-Company-ID"] = metaCache.metadata.selectedCompany.id;
      }
      return config;
    });

    this.req.interceptors.response.use(
      (res) => res,
      (error: AxiosError) => {
        if (error.response) {
          const { data, status } = error.response;
          const resBody: ErrorResponse =
            typeof data === "string" ? JSON.parse(data) : data;

          switch (status) {
            case 400:
              toast.warning(resBody?.err || "Bad request", toastOpts);
              break;
            case 401:
              toast.error(
                resBody?.err || "Unauthorized! Please login.",
                toastOpts
              );
              break;
            case 404:
              toast.error(resBody?.err || "Not found", toastOpts);
              break;
            case 500:
              toast.error(
                `Server error: ${resBody.err}! Please try again later.`,
                toastOpts
              );
              break;
            default:
              toast.error("An unexpected error occurred!", toastOpts);
          }
        } else {
          toast.error(
            "Network error! Please check your connection.",
            toastOpts
          );
        }
        return Promise.reject(error);
      }
    );
  }
}
