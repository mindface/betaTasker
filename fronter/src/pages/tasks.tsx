"use client"
import * as React from 'react'
import { NextPage } from 'next'
import Head from 'next/head'
import ContainersBase from '../containers/ContainersBase'
import SectionTaskView from '../components/SectionTaskView'

const Tasks: NextPage = () => {
  return (
    <div className="about">
      <Head>
        <title key="title">Visualer page</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <ContainersBase>
        <SectionTaskView />
      </ContainersBase>
    </div>
  )
}

export default Tasks
