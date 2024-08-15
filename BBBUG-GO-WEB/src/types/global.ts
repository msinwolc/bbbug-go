//聊天消息对象
export interface ChatMsgObj {
  msg: any;
  user_id: number;
  user_name: string;
  user_head: string;
  message_time: number;
}

// 房间对象
export interface Room {
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

// 用户信息
export interface UserMsg {
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

// 系统信息
export interface SystemMsg {
  access_token: string;
  plat: string;
  version: number;
}
