import * as React from "react";
import "./index.less";
import { useEffect, useState } from "react";

interface MenuBannerProps {
  roomId: number;
  roomName: string;
  inlineUserCount: number;
}

const TopBar: React.FunctionComponent<MenuBannerProps> = (props) => {
  const { roomId, roomName, inlineUserCount } = props;

  return (
    <div className={"top-bar-panel"}>
      <span>ID：{roomId || ""}</span>
      <span>{roomName || "聊天大厅"}</span>
      <span style={{ display: "inline-block", minWidth: "50px" }}>
        在线：{inlineUserCount || 0}
      </span>
    </div>
  );
};

export default TopBar;
