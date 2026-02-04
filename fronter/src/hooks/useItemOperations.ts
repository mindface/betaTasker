import { useCallback } from "react";
import { useApiCall } from "./useApiCall";
import { apiFacade, EntityType, ApiFacade } from "../facade/apiFacade";
import { Task, AddTask } from "../model/task";
import { Memory, AddMemory } from "../model/memory";
import { Assessment, AddAssessment } from "../model/assessment";

type EntityTypeMap = {
  task: { item: Task; create: AddTask };
  memory: { item: Memory; create: AddMemory };
  assessment: { item: Assessment; create: AddAssessment };
};

interface UseItemOperationsOptions {
  onDeleteSuccess?: () => void;
  onDeleteError?: (error: Error) => void;
  onUpdateSuccess?: () => void;
  onUpdateError?: (error: Error) => void;
}

export function useItemOperations<T extends EntityType>(
  itemType: T,
  options?: UseItemOperationsOptions,
) {
  type ItemType = EntityTypeMap[T]["item"];
  type CreateType = EntityTypeMap[T]["create"];

  const strategy = apiFacade[itemType] as ApiFacade<ItemType, CreateType>;

  const { execute: deleteItem, loading: deleteLoading } = useApiCall(
    (id: number) => strategy.delete(id),
    {
      onSuccess: () => {
        options?.onDeleteSuccess?.();
      },
      onError: (error) => {
        options?.onDeleteError?.(error);
      },
    },
  );

  const { execute: updateItem, loading: updateLoading } = useApiCall(
    (item: ItemType) => strategy.update(item),
    {
      onSuccess: () => {
        options?.onUpdateSuccess?.();
      },
      onError: (error) => {
        options?.onUpdateError?.(error);
      },
    },
  );

  const {
    execute: fetchItems,
    loading: fetchLoading,
    data: items,
  } = useApiCall(() => strategy.getAll() as Promise<ItemType[]>, {});

  const handleDelete = useCallback(
    async (id: number, itemTitle?: string) => {
      const message = itemTitle
        ? `「${itemTitle}」を削除しますか？この操作は取り消せません。`
        : "このアイテムを削除しますか？この操作は取り消せません。";

      if (confirm(message)) {
        await deleteItem(id);
      }
    },
    [deleteItem],
  );

  return {
    items: items as ItemType[] | null,
    deleteItem: handleDelete,
    updateItem,
    fetchItems,
    loading: {
      delete: deleteLoading,
      update: updateLoading,
      fetch: fetchLoading,
    },
  };
}
