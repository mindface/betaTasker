import React from "react";
import Link from "next/link";
import menu from "../json/menu.json";

interface Props {
  title: string;
}

function IndexHeader(props: Props) {
  const { title } = props;

  return (
    <header className="index-header">
      <div className="header--body _flex_s_b_">
        <h1 className="header__title">
          <Link href="/tools">{title}</Link>
        </h1>
        <nav className="g-nav">
          <ul className="list _flex_">
            {menu.map((item: { title: string; path: string }, idx: number) => (
              <li className="item" key={`indexmenu${idx}`}>
                <Link href={item.path} className="link">
                  {item.title}
                </Link>
              </li>
            ))}
          </ul>
        </nav>
      </div>
    </header>
  );
}

export default IndexHeader;
