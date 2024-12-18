import { toast } from "react-toastify";
import { toastOpts } from "../utils/toast";

export const useLoading = (message: string) => {
  const id = toast.loading(message, toastOpts);
  const start = new Date().getMilliseconds();
  console.log(`started at ${start}`);
  const done = (message: string, success: boolean) => {
    console.log(`ended at ${new Date().getMilliseconds()}`);
    toast.update(id, {
      render: message,
      type: success ? "success" : "error",
      isLoading: false,
      ...toastOpts,
      autoClose: 1500,
    });
  };
  return { done };
};
