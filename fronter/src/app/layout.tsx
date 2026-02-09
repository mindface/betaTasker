"use client";
import * as React from "react";
import { usePathname } from "next/navigation";
import { Provider } from "react-redux";
import { setupStore } from "../store";

import "../styles/style.sass";

export default function Layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const path = usePathname();
  return (
    <html lang="ja">
      <body>
        <Provider store={setupStore}>{children}</Provider>
      </body>
    </html>
  );
}
