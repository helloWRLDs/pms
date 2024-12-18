import { toast, ToastOptions } from "react-toastify";

export const toastOpts: ToastOptions = {
  theme: "dark",
  autoClose: 5000,
  closeOnClick: true,
  position: "bottom-right",
};

export const infoToast = (message: string) => {
  toast(message, {
    type: "info",
    ...toastOpts,
  });
};

export const errorToast = (message: string) => {
  toast(message, {
    type: "error",
    ...toastOpts,
  });
};

export const loadingToast = (message: string, loading: boolean) => {
  toast.loading(message, {
    isLoading: loading,
    ...toastOpts,
  });
};
