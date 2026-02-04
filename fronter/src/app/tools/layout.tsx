import * as React from "react";
import Head from "next/head";

export default function Layout({ children }: { children?: React.ReactNode }) {
  const title = "Task Tools";
  return (
    <>
      <Head>
        <title>{title}</title>
        <meta charSet="utf=8" />
      </Head>
      <div className="page-container-for-link">{children}</div>
    </>
  );
}
