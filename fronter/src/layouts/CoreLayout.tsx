import * as React from 'react'
import { Provider } from 'react-redux'
import { setupStore } from '../../../store'
import LinkLayout from './LinkLayout'
import RootLayout from './RootLayout'

type Props = {
  children?: React.ReactNode
}

export default function CoreLayout ({
  children,
}: Props) {

  return (
    <Provider store={setupStore}>
        <LinkLayout>{children}</LinkLayout>
      {/* {path.startsWith('/tools') ? (
        <LinkLayout>{children}</LinkLayout>
      ) : (
        <RootLayout>{children}</RootLayout>
      )} */}
    </Provider>
  );
}
