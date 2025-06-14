"use client"
import * as React from 'react'
import { Provider } from 'react-redux'
import { setupStore } from '../store'

// const store = setupStore()
import IndexLayout from '../layout/IndexLayout'

import '../styles/style.sass'


export default function BaseApp ({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body>
        <Provider store={setupStore}>
          <IndexLayout>
          <div>
            {children}
          </div>
          </IndexLayout>
        </Provider>
      </body>
    </html>
  )
}
