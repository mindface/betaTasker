import { useState, useCallback } from 'react';

interface UseApiCallOptions<T> {
  onSuccess?: (data: T) => void;
  onError?: (error: Error) => void;
}

interface UseApiCallResult<T, Args extends any[]> {
  execute: (...args: Args) => Promise<T | undefined>;
  loading: boolean;
  error: Error | null;
  data: T | null;
}

export function useApiCall<T, Args extends any[]>(
  apiCall: (...args: Args) => Promise<T>,
  options?: UseApiCallOptions<T>
): UseApiCallResult<T, Args> {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<Error | null>(null);
  const [data, setData] = useState<T | null>(null);

  const execute = useCallback(async (...args: Args): Promise<T | undefined> => {
    setLoading(true);
    setError(null);

    try {
      const result = await apiCall(...args);
      setData(result);
      options?.onSuccess?.(result);
      return result;
    } catch (err) {
      const errorObj = err instanceof Error ? err : new Error(String(err));
      setError(errorObj);
      options?.onError?.(errorObj);
      return undefined;
    } finally {
      setLoading(false);
    }
  }, [apiCall, options]);

  return { execute, loading, error, data };
}