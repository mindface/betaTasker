"use client";

import CommonDialog from "@/components/parts/CommonDialog";

interface ErrorDialogProps {
  title?: string;
  message: string;
  onRetry: () => void;
}

export function ErrorDialog({
  title = "エラーが発生しました",
  message = "しばらくしてから再度お試しください。",
  onRetry,
}: ErrorDialogProps) {
  return (
    <CommonDialog
      isOpen={true}
      onClose={onRetry}
      title={title}
      returnFocus={false}
    >
      <p className="p-8">{message}</p>
      <div style={{ marginTop: 16, textAlign: "right" }}>
        <button className="btn btn-primary border-radius" onClick={onRetry}>
          再読み込み
        </button>
      </div>
    </CommonDialog>
  );
}
