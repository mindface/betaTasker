"use client";
import React, { useEffect, useState } from "react";
import { Task } from "../../model/task";
import { HeuristicsAnalysis, HeuristicsPattern } from "../../model/heuristics";
import { useDispatch, useSelector } from "react-redux";
import { fetchAnalysisLimit } from "../../features/heuristics/heuristicsSlice";
import ListDialog from "./ListDialog";

import { RootState } from "../../store";

interface HeuristicsAnalysisDialogProps {
  task: Task;
}

const HeuristicsAnalysisDialog = ({ task }: HeuristicsAnalysisDialogProps) => {
  const dispath = useDispatch();
  const { analyses } = useSelector(
    (state: RootState) => state.heuristics,
  );

  const reData = (data: string, intKey: string) => {
    return JSON.parse(data)[intKey];
  };

  const viewAnyalze = (taskId: number) => {
    dispath(fetchAnalysisLimit({page: 1, limit: 20, task_id: taskId, include: "pattern," }));
  }

  const changehAeuristicsPatterns = (index: number, item: HeuristicsPattern) => {
    const patternList = JSON.parse(item.pattern);
    return <div key={`changehAeuristicsPatterns-${item.id}-${index}`} className="pattern-box">
      <div className="p-b-8">{item.name}</div>
      <div className="p-b-8">task type: {patternList.task_type}</div>
      <div className="p-b-8">{patternList.characteristics.join(" | ")}</div>
      <div className="p-b-8">{patternList.coefficient}</div>
      <div className="p-b-8">{patternList.outcomes.join(" | ")}</div>
      <div className="p-b-8">{patternList.triggers.join(" | ")}</div>
      {/* {patternList.map((item) => <></>)} */}
    </div>
  }

  return (
    <>
      <div className="heuristics-analysis-dialog">
        <ListDialog<HeuristicsAnalysis>
          initializationActionSet={() => viewAnyalze(task.id)}
          viewData={analyses}
          title="Heuristics Analysis"
          btnText="Heuristics Analysisを確認"
          indexType="HeuristicsAnalysis"
          renderItem={(item, index) => (
            <div key={`${index}`} className="language_optimizations-item p-8">
              <div className="p-b-8">analysis_type | {item.analysis_type}</div>
              <div className="p-b-8">confidence | {item.confidence}</div>
              <div className="p-b-8">
                difficulty_score | {item.difficulty_score}
              </div>
              <div className="p-b-8">
                efficiency_score | {item.efficiency_score}
              </div>
              <p>{reData(item.result, "outcomes")}</p>
              <p>{reData(item.result, "weaknesses")}</p>
              <ul className="list p-8">
                {item.heuristics_patterns &&
                  item.heuristics_patterns.map((item,index) => changehAeuristicsPatterns(index, item))}
              </ul>
            </div>
          )}
        />
      </div>
    </>
  );
};

export default HeuristicsAnalysisDialog;
