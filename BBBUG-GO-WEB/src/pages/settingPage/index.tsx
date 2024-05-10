import * as React from "react";
import "./index.less";
import { useEffect, useState } from "react";
import TopBar from "./components/topBar";
import Chatting from "./components/chatting";
import InputArea from "./components/inputArea";

interface MainChatProps {}

const MainChat: React.FunctionComponent<MainChatProps> = (props) => {
  const {} = props;

  const data = {
    roomId: 1234,
    roomName: "",
  };

  return (
    <div className={"mainChatPanel"}>
      <>设置</>
    </div>
  );
};

export default MainChat;
