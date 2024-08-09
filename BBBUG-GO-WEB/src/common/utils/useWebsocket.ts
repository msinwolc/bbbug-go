import { getWebsocketUrl } from "../../api/global";

export const useWebsocket = async (roomId?: number) => {
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
    console.log(JSON.parse(event.data));
  };
  socketNew.onclose = () => {
    console.log("连接关闭");
  };
  return socketNew;
};
