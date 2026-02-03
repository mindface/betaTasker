import React, { useState, useEffect } from 'react';
import CommonDialog from "./CommonDialog";

export interface ListDialogProps<T> {
  onClose?: () => void;
  viewData?:T[];
  indexType?: string;
  renderItem: (item: T, index: number) => React.ReactNode;
}

export default function ListDialog<T>({
  viewData,
  indexType,
  renderItem
}: ListDialogProps<T>) {
  const [viewDataState, setViewDataState] = useState<(T[])>([]);
  const [onDialog, setOnDialog] = useState(false)

  const onDialogHandler = () => {
    setOnDialog(!onDialog)
  }

  useEffect(() => {
    if (viewData) {
      setViewDataState(viewData);
    }else {
      setViewDataState([]);
    }
    console.log("viewData",viewData);
  }, [viewData]);
  return (
    <>
      <button onClick={onDialogHandler} className="btn btn-secondary">
        記録を確認
      </button>
      <CommonDialog
        isOpen={onDialog}
        onClose={onDialogHandler}
        title={viewData && viewData.length > 0 ? '記録を確認' : '新規メモ'}
      >
        <div className="language_optimizations-dialog">
          {viewDataState && viewDataState.length === 0 && (
            <p>表示するデータがありません。</p>
          )}
          {(viewData ?? []).map((item: T, index: number) => (
            <div key={`${indexType}-${index}`} className="language_optimizations-item">
              {renderItem(item, index)}
            </div>
          ))}
        </div>
      </CommonDialog>
    </>
  );
};