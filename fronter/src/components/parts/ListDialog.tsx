import React, { useState, useEffect } from "react";
import CommonDialog from "./CommonDialog";

export interface ListDialogProps<T> {
  onClose?: () => void;
  title?: string;
  btnText?: string;
  viewData?: T[];
  indexType?: string;
  renderItem: (item: T, index: number) => React.ReactNode;
}

export default function ListDialog<T>({
  viewData,
  title,
  btnText = "記録を確認",
  indexType,
  renderItem,
}: ListDialogProps<T>) {
  const [viewDataState, setViewDataState] = useState<T[]>([]);
  const [onDialog, setOnDialog] = useState(false);

  const onDialogHandler = () => {
    setOnDialog(!onDialog);
  };

  useEffect(() => {
    if (viewData) {
      setViewDataState(viewData);
    } else {
      setViewDataState([]);
    }
  }, [viewData]);

  return (
    <>
      <button onClick={onDialogHandler} className="btn btn-secondary">
        {btnText}
      </button>
      <CommonDialog
        isOpen={onDialog}
        onClose={onDialogHandler}
        title={
          title ||
          (viewData && viewData.length > 0
            ? `リストダイアログ ${title}`
            : "リストの確認")
        }
      >
        <div className="language_optimizations-dialog">
          {viewDataState && viewDataState.length === 0 && (
            <p>表示するデータがありません。</p>
          )}
          {(viewData ?? []).map((item: T, index: number) => (
            <div
              key={`${indexType}-${index}`}
              className="language_optimizations-item"
            >
              {renderItem(item, index)}
            </div>
          ))}
        </div>
      </CommonDialog>
    </>
  );
}
