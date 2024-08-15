import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { ChatMsgObj, SystemMsg, UserMsg } from "../types/global";

// 全局状态
export interface GlobalState {
  userMsg: UserMsg;
  systemMsg: SystemMsg;
  socket: WebSocket;
  chatMsg: ChatMsgObj[];
}

const gloableSlice = createSlice({
  name: "global",
  initialState: {
    userMsg: {},
    systemMsg: {
      access_token:
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIwMCwiVXNlck5hbWUiOiJ0ZXN0QHFxLmNvbSIsIlVzZXJQbGF0IjoidGVzdCIsImV4cCI6MTcxNjI3MTcxMSwiaWF0IjoxNzE1NjY2OTExfQ.6wIZZ-J6iV7LWOgzCYHBNMcifBQO5vR9XwH3YY4kd04",
      plat: "vue",
      version: 10000,
    },
    socket: {},
    chatMsg: [] as ChatMsgObj[],
  } as GlobalState,
  reducers: {
    setUserMsg: (state, action: PayloadAction<UserMsg>) => {
      state.userMsg = action.payload;
    },
    setSystemMsg: (state, action: PayloadAction<string>) => {
      state.systemMsg = { ...state.systemMsg, access_token: action.payload };
    },
    setSocket: (state, action: PayloadAction<WebSocket>) => {
      state.socket = action.payload;
    },
    setChatMsg: (state, action: PayloadAction<ChatMsgObj>) => {
      state.chatMsg.push(action.payload);
    },
  },
});

export const { setUserMsg, setSystemMsg, setSocket, setChatMsg } =
  gloableSlice.actions;
export default gloableSlice.reducer;
