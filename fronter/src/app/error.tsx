"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { ErrorDialog } from "@/components/partsError/ErrorDialog";

export default function Error({
  error,
  reset,
}: {
  error: Error & { status?: number };
  reset: () => void;
}) {
  const router = useRouter();

  useEffect(() => {
    console.error(error);
  }, [error]);

  if ((error as any).status === 401) {
    router.replace("/login");
    return null;
  }

  return (
    <ErrorDialog
      message={error.message}
      onRetry={reset}
    />
  );
}
