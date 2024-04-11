import * as React from "react";
import "./index.less";
import { useEffect, useState } from "react";

interface MenuBannerProps {
  roomId: number;
  roomName: string;
}

const TopBar: React.FunctionComponent<MenuBannerProps> = (props) => {
  const { roomId, roomName } = props;

  return (
    <div className={"top-bar-panel"}>
      <span>{roomId}</span>
      <span>{roomName}</span>
    </div>
  );
};

export default TopBar;
