import { AxiosResponse } from "axios";
import createAPIClient, { APIConfig } from "./client";

const responseBody = <T>(response: AxiosResponse<T>) => response.data;

const request = (config: APIConfig) => {
  const customApiClient = createAPIClient(config);
  return {
    get: <T>(url: string) => customApiClient.get<T>(url).then(responseBody),
    post: <T>(url: string, body: {}) =>
      customApiClient.post<T>(url, body).then(responseBody),
    put: <T>(url: string, body: {}) =>
      customApiClient.put<T>(url, body).then(responseBody),
    delete: <T>(url: string) =>
      customApiClient.delete<T>(url).then(responseBody),
  };
};

export { request };
