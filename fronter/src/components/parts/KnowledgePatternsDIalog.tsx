"use client";
import React, { useState } from "react";
import { KnowledgePattern, ConversionPath } from "../../model/knowledgePattern";
import CommonDialog from "./CommonDialog";

interface knowledgePatternsDialogProps {
  knowledgePatterns: KnowledgePattern[]
}

const KnowledgePatternsDialog = (props: knowledgePatternsDialogProps) => {
  const [isOpenDialog, setIsOpenDialog] = useState(false);

  const handleClose = () => {
    setIsOpenDialog(!isOpenDialog)
  }

  const RenderConversionPath = (item: ConversionPath) => {
    return <div className="pattern-box bg-gray">
      <h3 className="titile">conversion path</h3>
      <div className="p-b-8">level | {item.level}</div>
      <div className="p-b-8">task type: { item.pattern_type}</div>
      <div className="p-b-8">pattern type: { item.pattern_type}</div>
      <div className="p-b-8">{item.characteristics.join(" | ")}</div>
      <div className="p-b-8">{item.coefficient}</div>
      <div className="p-b-8">{item.outcomes.join(" | ")}</div>
      <div className="p-b-8">{item.triggers.join(" | ")}</div>
      {/* {patternList.map((item) => <></>)} */}
    </div>
  }

  return (
    <>
      <div className="knowledge-patterns-dialog">
        <button className="btn btn-secondary" onClick={handleClose}>knowledge patternsを確認</button>
        <CommonDialog
          isOpen={isOpenDialog}
          onClose={() => {
            handleClose()
          }}
          title="knowledge patterns"
          children={
            <div className="knowledge-patterns-item p-8">
              {props!.knowledgePatterns && <>表示する情報がありません。</>}
              {props.knowledgePatterns.map((item,index) => <div key={`changeh-${item.id}-${index}`} className="pattern-box">
                <div className="p-b-8">{item.domain}</div>
                <div className="p-b-8">
                  <span className="d-inline-block bg-gray border-radius p-4">task type:</span>
                  {item.type}
                </div>
                <div className="p-b-8">
                  <span className="d-inline-block bg-gray border-radius p-4">accuracy:</span>
                  {item.accuracy}
                </div>
                <div className="p-b-8">
                  <span className="d-inline-block bg-gray border-radius p-4">coverage:</span>
                  {item.accuracy}
                </div>
                <div className="p-b-8">
                  <span className="d-inline-block bg-gray border-radius p-4">abstract_level:</span>
                    {item.abstract_level}
                </div>
                <div className="p-b-8">
                  {RenderConversionPath(item.conversion_path)}
                </div>
              </div>)}
            </div>
          }
        />
      </div>
    </>
  );
};

export default KnowledgePatternsDialog;
