import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface Room {
  room_id: number;
  room_user: number;
  room_addsongcd: number;
  room_addcount: number;
  room_pushdaycount: number;
  room_pushsongcd: number;
  room_online: number;
  room_realonline: number;
  room_hide: number;
  room_name: string;
  room_type: number;
  room_public: number;
  room_password: string;
  room_notice: string;
  room_addsong: number;
  room_sendmsg: number;
  room_robot: number;
  room_order: number;
  room_reason: string;
  room_playone: number;
  room_votepass: number;
  room_votepercent: number;
  room_background: string;
  room_app: string;
  room_fullpage: number;
  room_status: number;
  room_createtime: number;
  room_updatetime: number;
  admin: any;
}
interface UserMsg {
  pass_count: number;
  push_count: number;
  user_admin: boolean;
  user_id: number;
  user_icon: number;
  user_sex: number;
  user_name: string;
  user_head: string;
  user_remark: string;
  user_extra: string;
  user_device: string;
  user_touchtip: string;
  user_vip: string;
  user_group: number;
  myRoom: Room;
  user_shutdown: boolean;
  user_songdown: boolean;
  user_guest: boolean;
}

interface SystemMsg {
  access_token: string;
  plat: string;
  version: number;
}

export interface GlobalState {
  userMsg: UserMsg;
  systemMsg: SystemMsg;
  socket: WebSocket;
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
  } as GlobalState,
  reducers: {
    setUserMsg: (state, action: PayloadAction<UserMsg>) => {
      state.userMsg = action.payload;
    },
    setSystemMsg: (state, action: PayloadAction<SystemMsg>) => {
      state.systemMsg = action.payload;
    },
    setSocket: (state, action: PayloadAction<WebSocket>) => {
      state.socket = action.payload;
    },
  },
});

export const { setUserMsg, setSystemMsg, setSocket } = gloableSlice.actions;
export default gloableSlice.reducer;
