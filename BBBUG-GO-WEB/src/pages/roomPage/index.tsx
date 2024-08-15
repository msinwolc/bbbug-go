import * as React from "react";
import "./index.less";
import { useEffect, useState } from "react";
import TopBar from "./components/topBar";
import Chatting from "./components/chatting";
import InputArea from "./components/inputArea";
import { getUserMsg } from "../../api/global";
import { setSocket, setUserMsg, setChatMsg } from "../../store/global";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../../store/store";
import { isEmpty } from "lodash";
import { useWebsocket } from "../../common/utils/useWebsocket";

interface MainChatProps {}

const MainChat: React.FunctionComponent<MainChatProps> = (props) => {
  const {} = props;
  const dispatch = useDispatch();

  const systemMsg = useSelector((state: RootState) => state.systemMsg);
  const userMsg = useSelector((state: RootState) => state.userMsg);
  const chatMsg = useSelector((state: RootState) => state.chatMsg);

  const [data, setData] = useState({
    roomId: 1234,
    roomName: "",
    inlineUserCount: 0,
    msgList: [],
  });

  let socket = useSelector((state: RootState) => state.socket);

  // useEffect(() => {
  //   if (isEmpty(socket) || socket?.readyState === WebSocket.CLOSED) {
  //     useWebsocket(chatMsg, dispatch, setChatMsg, userMsg).then((socketNew) => {
  //       dispatch(setSocket(socketNew));
  //     });
  //     return () => {
  //       console.log("组件卸载");
  //       if (
  //         socket.readyState === WebSocket.OPEN ||
  //         socket.readyState === WebSocket.CONNECTING
  //       ) {
  //         socket.close();
  //       }
  //     };
  //   }
  // }, []);

  useEffect(() => {
    // 获取用户信息
    const access_token = localStorage.getItem("access_token");
    getUserMsg({ ...systemMsg, access_token: access_token }).then(
      (res: any) => {
        // 变更全局用户信息
        dispatch(setUserMsg(res));
        // 创建socket
        if (isEmpty(socket) || socket?.readyState === WebSocket.CLOSED) {
          useWebsocket(chatMsg, dispatch, setChatMsg, res).then((socketNew) => {
            dispatch(setSocket(socketNew));
          });
        }
      }
    );

    return () => {
      console.log("组件卸载");
      if (
        socket.readyState === WebSocket.OPEN ||
        socket.readyState === WebSocket.CONNECTING
      ) {
        socket.close();
      }
    };
  }, []);

  return (
    <div className={"mainChatPanel"}>
      <div className={"mainChatPanel-topBar"}>
        <TopBar
          roomId={data.roomId}
          roomName={data.roomName}
          inlineUserCount={data.inlineUserCount}
        />
      </div>
      <div className={"mainChatPanel-chatting"}>
        <Chatting msgList={chatMsg} />
      </div>
      <div className={"mainChatPanel-inputArea"}>
        <InputArea />
      </div>
    </div>
  );
};

export default MainChat;
