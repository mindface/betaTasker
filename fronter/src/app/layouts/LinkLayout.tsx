import * as React from 'react'
import Head from 'next/head'

type Props = {
  children?: React.ReactNode
  title?: string,
  header_category?: string,
}

export default function Layout ({
  children,
  title = 'Link maker',
  header_category = 'Link Maker',
}: Props) {
  return (
      <>
        <Head>
          <title>{title}</title>
          <meta charSet="utf=8" />
        </Head>
        <div className="page-container-for-link">
          {children}
        </div>
      </>
    )
}
