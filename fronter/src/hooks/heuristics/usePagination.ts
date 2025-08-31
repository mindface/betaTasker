import { useState, useCallback, useMemo } from 'react';

export interface PaginationConfig {
  initialPage?: number;
  initialPageSize?: number;
  pageSizeOptions?: number[];
  showSizeChanger?: boolean;
  showQuickJumper?: boolean;
  total?: number;
}

export interface PaginationResult {
  currentPage: number;
  pageSize: number;
  total: number;
  totalPages: number;
  hasNext: boolean;
  hasPrevious: boolean;
  offset: number;
  
  // Actions
  setPage: (page: number) => void;
  setPageSize: (size: number) => void;
  setTotal: (total: number) => void;
  nextPage: () => void;
  previousPage: () => void;
  firstPage: () => void;
  lastPage: () => void;
  
  // Computed values
  startIndex: number;
  endIndex: number;
  pageNumbers: number[];
  
  // Validation
  isValidPage: (page: number) => boolean;
  
  // Reset
  reset: () => void;
}

export const usePagination = (config: PaginationConfig = {}): PaginationResult => {
  const {
    initialPage = 1,
    initialPageSize = 10,
    total: initialTotal = 0,
    pageSizeOptions = [10, 20, 50, 100]
  } = config;

  const [currentPage, setCurrentPage] = useState(initialPage);
  const [pageSize, setCurrentPageSize] = useState(initialPageSize);
  const [total, setCurrentTotal] = useState(initialTotal);

  const totalPages = useMemo(() => {
    return Math.max(1, Math.ceil(total / pageSize));
  }, [total, pageSize]);

  const hasNext = useMemo(() => {
    return currentPage < totalPages;
  }, [currentPage, totalPages]);

  const hasPrevious = useMemo(() => {
    return currentPage > 1;
  }, [currentPage]);

  const offset = useMemo(() => {
    return (currentPage - 1) * pageSize;
  }, [currentPage, pageSize]);

  const startIndex = useMemo(() => {
    return total === 0 ? 0 : offset + 1;
  }, [offset, total]);

  const endIndex = useMemo(() => {
    return Math.min(offset + pageSize, total);
  }, [offset, pageSize, total]);

  const pageNumbers = useMemo(() => {
    const maxPagesToShow = 7;
    const pages: number[] = [];

    if (totalPages <= maxPagesToShow) {
      // Show all pages
      for (let i = 1; i <= totalPages; i++) {
        pages.push(i);
      }
    } else {
      // Show selected pages with ellipsis logic
      const halfRange = Math.floor(maxPagesToShow / 2);
      let start = Math.max(1, currentPage - halfRange);
      let end = Math.min(totalPages, currentPage + halfRange);

      // Adjust if we're near the beginning
      if (currentPage - halfRange < 1) {
        end = Math.min(totalPages, end + (halfRange - currentPage + 1));
      }

      // Adjust if we're near the end
      if (currentPage + halfRange > totalPages) {
        start = Math.max(1, start - (currentPage + halfRange - totalPages));
      }

      // Always show first page
      if (start > 1) {
        pages.push(1);
        if (start > 2) {
          pages.push(-1); // Represents ellipsis
        }
      }

      // Add page range
      for (let i = start; i <= end; i++) {
        pages.push(i);
      }

      // Always show last page
      if (end < totalPages) {
        if (end < totalPages - 1) {
          pages.push(-1); // Represents ellipsis
        }
        pages.push(totalPages);
      }
    }

    return pages;
  }, [currentPage, totalPages]);

  const isValidPage = useCallback((page: number): boolean => {
    return page >= 1 && page <= totalPages;
  }, [totalPages]);

  const setPage = useCallback((page: number) => {
    if (isValidPage(page)) {
      setCurrentPage(page);
    }
  }, [isValidPage]);

  const setPageSize = useCallback((size: number) => {
    if (pageSizeOptions.includes(size)) {
      setCurrentPageSize(size);
      // Reset to first page when changing page size
      setCurrentPage(1);
    }
  }, [pageSizeOptions]);

  const setTotal = useCallback((newTotal: number) => {
    setCurrentTotal(Math.max(0, newTotal));
    // Adjust current page if it's beyond the new total pages
    const newTotalPages = Math.max(1, Math.ceil(newTotal / pageSize));
    if (currentPage > newTotalPages) {
      setCurrentPage(newTotalPages);
    }
  }, [currentPage, pageSize]);

  const nextPage = useCallback(() => {
    if (hasNext) {
      setCurrentPage(currentPage + 1);
    }
  }, [currentPage, hasNext]);

  const previousPage = useCallback(() => {
    if (hasPrevious) {
      setCurrentPage(currentPage - 1);
    }
  }, [currentPage, hasPrevious]);

  const firstPage = useCallback(() => {
    setCurrentPage(1);
  }, []);

  const lastPage = useCallback(() => {
    setCurrentPage(totalPages);
  }, [totalPages]);

  const reset = useCallback(() => {
    setCurrentPage(initialPage);
    setCurrentPageSize(initialPageSize);
    setCurrentTotal(initialTotal);
  }, [initialPage, initialPageSize, initialTotal]);

  return {
    currentPage,
    pageSize,
    total,
    totalPages,
    hasNext,
    hasPrevious,
    offset,
    startIndex,
    endIndex,
    pageNumbers,
    setPage,
    setPageSize,
    setTotal,
    nextPage,
    previousPage,
    firstPage,
    lastPage,
    isValidPage,
    reset
  };
};

// 追加のヘルパーフック：サーバーサイド・ページング用
export interface ServerPaginationConfig extends PaginationConfig {
  onPageChange?: (page: number, pageSize: number) => void;
  loading?: boolean;
}

export const useServerPagination = (config: ServerPaginationConfig = {}) => {
  const { onPageChange, loading = false } = config;
  const pagination = usePagination(config);

  const handlePageChange = useCallback((page: number) => {
    if (!loading && pagination.isValidPage(page)) {
      pagination.setPage(page);
      onPageChange?.(page, pagination.pageSize);
    }
  }, [loading, onPageChange, pagination]);

  const handlePageSizeChange = useCallback((size: number) => {
    if (!loading) {
      pagination.setPageSize(size);
      onPageChange?.(1, size);
    }
  }, [loading, onPageChange, pagination]);

  return {
    ...pagination,
    setPage: handlePageChange,
    setPageSize: handlePageSizeChange,
    loading
  };
};

// ページング情報のフォーマット用ヘルパー
export const formatPaginationInfo = (pagination: PaginationResult): string => {
  if (pagination.total === 0) {
    return '0件のアイテム';
  }

  return `${pagination.startIndex}〜${pagination.endIndex}件 / ${pagination.total}件中`;
};

export const formatPageInfo = (pagination: PaginationResult): string => {
  return `${pagination.currentPage} / ${pagination.totalPages}`;
};