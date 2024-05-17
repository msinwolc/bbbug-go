import * as React from "react";
import "./index.less";
import { useEffect, useState } from "react";
import TopBar from "./components/topBar";
import Chatting from "./components/chatting";
import InputArea from "./components/inputArea";
import { getUserMsg } from "../../api/global";
import { setUserMsg } from "../../store/global";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../../store/store";

interface MainChatProps {}

const MainChat: React.FunctionComponent<MainChatProps> = (props) => {
  const {} = props;
  const dispatch = useDispatch();

  const systemMsg = useSelector((state: RootState) => state.systemMsg);

  const data = {
    roomId: 1234,
    roomName: "",
  };

  useEffect(() => {
    // 获取用户信息
    getUserMsg(systemMsg).then((res: any) => {
      dispatch(setUserMsg(res));
    });
  }, []);

  return (
    <div className={"mainChatPanel"}>
      <>房间</>
      <TopBar roomId={data.roomId} roomName={data.roomName} />
      <Chatting />
      <InputArea />
    </div>
  );
};

export default MainChat;
