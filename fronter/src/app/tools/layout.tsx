"use client"
import * as React from 'react'
import { Provider } from 'react-redux'
import { setupStore } from '../../store'

// const store = setupStore()
import LinkLayout from '../layouts/LinkLayout'

import '../../styles/style.sass'

export default function layout ({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body>
        <Provider store={setupStore}>
          <LinkLayout>
          <div>
            {children}
          </div>
          </LinkLayout>
        </Provider>
      </body>
    </html>
  )
}
