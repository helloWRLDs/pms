export interface ErrorResponse {
  msg: string;
  status: number;
  err: string;
}

export const parseError = (error: any): ErrorResponse | null => {
  const e = JSON.parse(JSON.stringify(error));
  return e;
};
