import { useQuery } from "@tanstack/react-query";
import { useEffect } from "react";

interface Options<TData, TArgs extends any[]> {
  keys: string[];
  fn: (...args: TArgs) => Promise<TData>;
  fnargs?: TArgs;
  isRefetch?: boolean;
  isEnabled?: boolean;
}

export function useFetch<TData = unknown, TArgs extends any[] = []>(
  opts: Options<TData, TArgs>
) {
  const { keys, fn, fnargs, isRefetch = false, isEnabled = true } = opts;

  const safeArgs = (fnargs ?? []) as TArgs;

  const { data, isLoading, isSuccess, isError, refetch } = useQuery({
    queryKey: keys,
    queryFn: () => fn(...safeArgs),
    enabled: isEnabled,
    staleTime: 1000,
  });

  useEffect(() => {
    if (isRefetch) refetch();
  }, [isRefetch, refetch]);

  useEffect(() => {
    if (isSuccess) console.log("Data fetched successfully");
  }, [isSuccess]);

  useEffect(() => {
    if (isError) console.log("Error fetching data");
  }, [isError]);

  return {
    responseData: data,
    isLoading,
    isSuccess,
    isError,
  };
}
