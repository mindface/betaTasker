"use clinet"
import React, {} from 'react'

type Props = {
  children?: React.ReactNode
}

function ContentModal(info:HTMLDivElement)  {
  const openBtn = document.createElement('button')
  openBtn.className = 'btn btn--primary open'
  openBtn.innerHTML = 'open'
  const closeBtn = document.createElement('button')
  closeBtn.className = 'btn btn--primary close'
  closeBtn.innerHTML = 'close'

  const modalOuter = document.createElement('div')
  modalOuter.className = 'modal-transfer';

  const modal = document.createElement('div')
  modal.id = 'modal-root';
  modal.className = 'modal-box';
  modal.appendChild(info)
  // if (false) {
  //   modal.innerHTML = String(info)
  // }
  // modal.appendChild(openBtn)
  modal.appendChild(closeBtn)
  modalOuter.appendChild(modal)

  openBtn.addEventListener('click', () => {
  })

  closeBtn.addEventListener('click', () => {
    modalOuter.remove()
  })

  return modalOuter
}

export default ContentModal