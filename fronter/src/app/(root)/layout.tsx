import * as React from 'react'
import Head from 'next/head'
import IndexHeader from '../../components/IndexHeader'
import IndexFooter from '../../components/IndexFooter'

export default function RootLayout ({
  children,
}: {
  children: React.ReactNode
}) {
  const title = 'Task Link';
  const header_category = 'Task Flow';

  return (
      <>
        <Head>
          <title>{title}</title>
          <meta charSet="utf=8" />
        </Head>
        <IndexHeader title={header_category}/>
        {children}
        <IndexFooter title="&copy; realize" />
      </>
    )
}
