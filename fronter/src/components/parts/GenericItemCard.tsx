"use client";
import React, { useState } from "react";
import Image from "next/image";
import CommonModal from "./CommonModal";

export interface GenericItemCardProps<T> {
  item: T;
  itemType: "task" | "assessment" | "memory";
  onDelete: (id: number) => Promise<void>;
  onUpdate: (item: T) => Promise<void>;
  onView?: (item: T) => void;
  renderContent: (item: T) => React.ReactNode;
  renderEditForm?: (item: T, onChange: (item: T) => void) => React.ReactNode;
  getTitle: (item: T) => string;
  getDescription: (item: T) => string;
  getId: (item: T) => number;
}

export default function GenericItemCard<T>({
  item,
  onDelete,
  onUpdate,
  onView,
  renderContent,
  renderEditForm,
  getTitle,
  getDescription,
  getId,
}: GenericItemCardProps<T>) {
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [isViewModalOpen, setIsViewModalOpen] = useState(false);
  const [editingItem, setEditingItem] = useState<T>(item);

  const handleDelete = async () => {
    const title = getTitle(item);
    const id = getId(item);

    if (confirm(`「${title}」を削除しますか？この操作は取り消せません。`)) {
      await onDelete(id);
    }
  };

  const handleUpdate = async () => {
    await onUpdate(editingItem);
    setIsEditModalOpen(false);
  };

  const handleView = () => {
    if (onView) {
      onView(item);
    } else {
      setIsViewModalOpen(true);
    }
  };

  const handleEditClick = () => {
    setEditingItem(item);
    setIsEditModalOpen(true);
  };

  return (
    <>
      <div className="card">
        <h4 className="card__title">{getTitle(item)}</h4>
        <div className="card__body">
          <div className="description">
            <p className="description-text">{getDescription(item)}</p>
          </div>
          {renderContent(item)}
        </div>
        <div className="card__btns">
          <div className="btns">
            <button className="btn" onClick={handleView} title="表示">
              <Image src="/image/look.svg" alt="表示" width={20} height={20} />
            </button>
            {renderEditForm && (
              <button className="btn" onClick={handleEditClick} title="編集">
                <Image
                  src="/image/edit.svg"
                  alt="編集"
                  width={20}
                  height={20}
                />
              </button>
            )}
            <button className="btn" onClick={handleDelete} title="削除">
              <Image
                src="/image/delete.svg"
                alt="削除"
                width={20}
                height={20}
              />
            </button>
          </div>
        </div>
      </div>

      {/* 編集モーダル */}
      {renderEditForm && (
        <CommonModal
          isOpen={isEditModalOpen}
          onClose={() => setIsEditModalOpen(false)}
          title={`${getTitle(item)}を編集`}
        >
          <div className="modal-content">
            {renderEditForm(editingItem, setEditingItem)}
            <div className="form-actions">
              <button
                type="button"
                onClick={() => setIsEditModalOpen(false)}
                className="btn btn-secondary"
              >
                キャンセル
              </button>
              <button
                type="button"
                onClick={handleUpdate}
                className="btn btn-primary"
              >
                保存
              </button>
            </div>
          </div>
        </CommonModal>
      )}

      {/* 表示モーダル */}
      <CommonModal
        isOpen={isViewModalOpen}
        onClose={() => setIsViewModalOpen(false)}
        title={getTitle(item)}
      >
        <div className="modal-content">
          <div className="view-content">
            <p className="description">{getDescription(item)}</p>
            {renderContent(item)}
          </div>
        </div>
      </CommonModal>
    </>
  );
}
