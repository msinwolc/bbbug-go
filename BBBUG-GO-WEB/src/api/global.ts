/**
 * 通用接口
 */
import { post } from "../common/utils/httpRequest";

export const getUserMsg = (params: any) => {
  return post("/api/user/getmyinfo", params);
};
