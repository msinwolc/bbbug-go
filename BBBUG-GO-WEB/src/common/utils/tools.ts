import { sendMsg } from "../../api/global";

interface sendSocketMsgProps {
  type: "text" | "img";
  msg: string;
  roomId: number;
}
export const sendSocketMsg = (params: sendSocketMsgProps) => {
  const { type, msg, roomId } = params;
  const msgObj = {
    at: false,
    where: "channel",
    to: roomId,
    type: type,
    msg: encodeURIComponent(msg),
    access_token: localStorage.getItem("access_token"),
    plat: "vue",
    version: 10000,
  };
  sendMsg(msgObj);
};
