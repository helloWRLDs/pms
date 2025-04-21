import { Company } from "../lib/company/company";
import { List } from "../lib/utils";
import { buildQuery, Pagination } from "../lib/utils/pagination";
import { request } from "./api";
import { APIConfig } from "./client";

const companyAPI = (access_token?: string) => {
  const config: APIConfig = {
    baseURL: "http://localhost:8080/api/v1",
    headers: {
      "Content-Type": "application/json",
    },
  };
  if (access_token) {
    config.headers.Authorization = `Bearer ${access_token}`;
  }

  const req = request(config);

  const listCompanies = (pagination?: Pagination) => {
    return req.get<List<Company>>(buildQuery("/companies", pagination));
  };

  return { listCompanies };
};

export default companyAPI;
