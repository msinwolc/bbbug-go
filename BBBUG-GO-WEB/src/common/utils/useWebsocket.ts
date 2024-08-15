import {
  ActionCreatorWithPayload,
  AnyAction,
  Dispatch,
} from "@reduxjs/toolkit";
import { getWebsocketUrl } from "../../api/global";
import { ChatMsgObj, UserMsg } from "../../types/global";

export const useWebsocket = async (
  chatMsg: any[],
  dispatch: Dispatch<AnyAction>,
  setChatMsg: ActionCreatorWithPayload<ChatMsgObj, "global/setChatMsg">,
  userMsg: UserMsg,
  roomId?: number
) => {
  console.log("@@@@@@@@@@@@@@@@@@@@@@@");
  const access_token = localStorage.getItem("access_token");
  const res: any = await getWebsocketUrl({
    access_token:
      access_token ||
      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIwMCwiVXNlck5hbWUiOiJ0ZXN0QHFxLmNvbSIsIlVzZXJQbGF0IjoidGVzdCIsImV4cCI6MTcxNjQ1NzA1NiwiaWF0IjoxNzE1ODUyMjU2fQ.siLCsJv6zi78t2Lkht8pXwLMhiEttwsllAGlRhDVMuM",
    channel: roomId || 888,
    plat: "vue",
    referer: false,
    version: 10000,
  });
  const socketNew = new WebSocket(
    "wss://saysth.fun/socket?account=" +
      res.account +
      "&channel=" +
      res.channel +
      "&ticket=" +
      res.ticket
  );
  socketNew.onopen = () => {
    console.log("连接成功");
  };
  socketNew.onmessage = (event: any) => {
    const data = JSON.parse(event.data);
    if (data?.type === "text" && data?.user.user_id !== userMsg.user_id) {
      dispatch(
        setChatMsg({
          msg: data?.content,
          user_id: data?.user.user_id,
          user_name: data?.user.user_name,
          user_head: data?.user.user_head,
          message_time: data?.message_time,
        })
      );
    }
  };
  socketNew.onclose = () => {
    console.log("连接关闭");
  };
  return socketNew;
};
