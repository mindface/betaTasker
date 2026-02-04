"use client";
import React, { useState, useRef } from "react";
import SvgRect from "./partsSvg/SvgRect";

export default function SectionBetaTools() {
  const svgElement = useRef(null);
  const [items, setItems] = useState<{ id: number; text: string }[]>([]);

  const addRect = () => {
    setItems([...items, { id: items.length, text: "text" }]);
  };

  return (
    <div className="section__inner section--tools">
      <div className="section-continer">
        <div className="tools-header">
          <button
            onClick={(e) => {
              addRect();
            }}
          >
            Rect追加
          </button>
        </div>
        <div className="tools__body" ref={svgElement}>
          <svg
            id="svg"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            width="500"
            height="500"
            viewBox="0 0 500 500"
          >
            {items.map((item) => {
              return <SvgRect key={item.id} />;
            })}
          </svg>
        </div>
      </div>
    </div>
  );
}
