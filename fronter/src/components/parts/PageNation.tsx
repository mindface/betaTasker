"use client";
import React, { useEffect, useState } from "react";

type PageNationProps = {
  page?: number;                    // 初期ページ（省略可）
  limit?: number;                   // 1ページあたり件数
  totalPages: number;              // 総ページ数
  onChange?: (page: number) => void;
};

export default function PageNation({
  page = 1,
  limit = 20,
  totalPages,
  onChange,
}: PageNationProps) {
  // 内部 state
  const [currentPage, setCurrentPage] = useState<number>(page);

  // 親から page が変わったら同期
  useEffect(() => {
    setCurrentPage(page);
  }, [page]);

  const canPrev = currentPage > 1;
  const canNext = currentPage < totalPages;

  const MAX_PAGES = 5;

  const startPage = Math.max(
    1,
    currentPage - Math.floor(MAX_PAGES / 2),
  );
  const endPage = Math.min(
    totalPages,
    startPage + MAX_PAGES - 1,
  );

  const pages: number[] = [];
  for (let i = startPage; i <= endPage; i++) {
    pages.push(i);
  }

  const changePage = (p: number) => {
    setCurrentPage(p);
    onChange?.(p);
  };

  return (
    <div className="pagination" style={{ display: "flex", gap: "6px" }}>
      {/* Prev */}
      <button
        disabled={!canPrev}
        onClick={() => changePage(currentPage - 1)}
      >
        Prev
      </button>

      {/* First */}
      {startPage > 1 && (
        <>
          <button onClick={() => changePage(1)}>1</button>
          <span>…</span>
        </>
      )}
      <p>startPage: {startPage}</p>
      <p>endPage: {endPage}</p>

      {/* Page numbers */}
      {pages.map((p) => (
        <button
          key={p}
          onClick={() => changePage(p)}
          style={{
            fontWeight: p === currentPage ? "bold" : "normal",
            textDecoration: p === currentPage ? "underline" : "none",
          }}
        >
          {p}
        </button>
      ))}

      {/* Last */}
      {endPage < totalPages && (
        <>
          <span>…</span>
          <button onClick={() => changePage(totalPages)}>
            {totalPages}
          </button>
        </>
      )}

      {/* Next */}
      <button
        disabled={!canNext}
        onClick={() => changePage(currentPage + 1)}
      >
        Next
      </button>
    </div>
  );
}
