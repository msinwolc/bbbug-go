import * as React from "react";
import "./index.less";
import MenuBanner from "../menuBanner";
import { Outlet } from "react-router-dom";

interface MainPageProps {}

const MainPage: React.FunctionComponent<MainPageProps> = (props) => {
  const {} = props;

  return (
    <div className={"mainPage-baner"}>
      <MenuBanner />
      <div style={{ display: "inline-block", flex: 1 }}>
        <Outlet />
      </div>
    </div>
  );
};

export default MainPage;
