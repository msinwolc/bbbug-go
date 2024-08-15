import { Popover } from "antd";
import * as React from "react";
import "./index.less";

interface DialogProps {
  msgObj: any;
}

const Dialog: React.FunctionComponent<DialogProps> = (props) => {
  const { msgObj } = props;
  return (
    <>
      <div className={"dialog-panel"}>
        <div className="user_pic">
          {msgObj.picUrl ? (
            <img src="src\assets\user.png" alt="" />
          ) : (
            <img src="src\assets\user.png" alt="" />
          )}
        </div>
        <div className="user_msg_panel">
          <p className="user_name">{msgObj.user_name}</p>
          <p className="user_msg">{msgObj.msg}</p>
        </div>
      </div>
    </>
  );
};

export default Dialog;
