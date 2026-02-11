"use client";
import * as React from "react";
import { Provider } from "react-redux";
import { setupStore } from "../store";

import "../styles/style.sass";

export default function Layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body>
        <Provider store={setupStore}>{children}</Provider>
      </body>
    </html>
  );
}
