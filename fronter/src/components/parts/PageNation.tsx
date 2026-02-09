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
  const [currentPage, setCurrentPage] = useState<number>(page);

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
    if(currentPage === p) return;
    setCurrentPage(p);
    onChange?.(p);
  };

  return (
    <div className="pagenation p-16" style={{ display: "flex", gap: "6px" }}>
      <button
        disabled={!canPrev}
        onClick={() => changePage(currentPage - 1)}
      >
        &lt; &lt;
      </button>

      {startPage > 1 && (
        <>
          <button onClick={() => changePage(1)}>1</button>
          <span>…</span>
        </>
      )}
      {pages.map((p) => (
        <button
          key={p}
          className={currentPage === p ? "active" : ""}
          onClick={() => changePage(p)}
          style={{
            fontWeight: p === currentPage ? "bold" : "normal",
            textDecoration: p === currentPage ? "underline" : "none",
          }}
        >
          {p}
        </button>
      ))}

      {endPage < totalPages && (
        <>
          <span>…</span>
          <button onClick={() => changePage(totalPages)}>
            {totalPages}
          </button>
        </>
      )}

      <button
        disabled={!canNext}
        onClick={() => changePage(currentPage + 1)}
      >
        &gt; &gt;
      </button>
    </div>
  );
}
