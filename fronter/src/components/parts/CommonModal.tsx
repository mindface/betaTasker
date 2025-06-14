import React from 'react';

interface CommonModalProps {
  isOpen: boolean;
  onClose: () => void;
  title: string;
  children: React.ReactNode;
}

const CommonModal: React.FC<CommonModalProps> = ({
  isOpen,
  onClose,
  title,
  children,
}) => {

  return (
    <>
      {isOpen && <div className="modal-overlay">
        <div className="modal-content">
          <div className="modal-header">
            <h2>{title}</h2>
            <button onClick={onClose} className="modal-close">
              ×
            </button>
          </div>
          <div className="modal-body">
            {children}
          </div>
        </div>
      </div>}
    </>
  );
};

export default CommonModal; 