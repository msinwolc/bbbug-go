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
      <Outlet />
    </div>
  );
};

export default MainPage;
