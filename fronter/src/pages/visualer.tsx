"use client"
import * as React from 'react'
import { NextPage } from 'next'
import Head from 'next/head'
import ContainersBase from '../containers/ContainersBase'
import SectionViewer from '../components/SectionViewer'

const Visualer: NextPage = () => {
  return (
    <div className="about">
      <Head>
        <title key="title">Visualer page</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <ContainersBase>
        <SectionViewer />
      </ContainersBase>
    </div>
  )
}

export default Visualer
