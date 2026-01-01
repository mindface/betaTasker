import * as React from 'react'
import { Provider } from 'react-redux'
import { setupStore } from '../../store'
import { usePathname } from 'next/navigation'
import LinkLayout from './LinkLayout'
import RootLayout from './RootLayout'

type Props = {
  children?: React.ReactNode
  title?: string,
  header_category?: string,
}

export default function CoreLayout ({
  children,
}: {
  children: React.ReactNode;
}) {
  const path = usePathname();

  return (
    <Provider store={setupStore}>
      {path.startsWith('/tools') ? (
        <LinkLayout>{children}</LinkLayout>
      ) : (
        <RootLayout>{children}</RootLayout>
      )}
    </Provider>
  );
}
