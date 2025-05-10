import * as React from 'react'
import Head from 'next/head'
import IndexHeader from '../components/IndexHeader'
import IndexFooter from '../components/IndexFooter'

type Props = {
  children?: React.ReactNode
  title?: string,
  header_category?: string,
}

export default function IndexLayout ({
  children,
  title = 'Task Link',
  header_category = 'Task Link',
}:Props) {
  return (
      <>
        <Head>
          <title>{title}</title>
          <meta charSet="utf=8" />
        </Head>
        <div>
          <div className="page-container">
            <IndexHeader title={header_category}/>
            {children}
            <IndexFooter title="&copy; realize" />
          </div>
        </div>
      </>
    )
}
