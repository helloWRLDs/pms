import { jwtDecode } from "jwt-decode";

export const isTokenValid = (token: string): boolean => {
  try {
    const decoded = jwtDecode(token);
    return decoded?.exp && decoded.exp * 1000 > Date.now() ? true : false;
  } catch (e) {
    console.log(`failed decoding token ${token}`);
    return false;
  }
};
