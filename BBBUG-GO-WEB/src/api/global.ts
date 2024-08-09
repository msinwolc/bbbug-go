/**
 * 通用接口
 */
import { post } from "../common/utils/httpRequest";

export const getUserMsg = (params: any) => {
  return post("/api/user/getmyinfo", params);
};

export const getWebsocketUrl = (params: any) => {
  return post("/api/room/getWebsocketUrl", params);
};

export const doLogin = (params: any) => {
  return post("/api/user/login", params);
};

export const sendMsg = (params: any) => {
  return post("/api/message/send", params);
};