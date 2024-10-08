import * as React from "react";
import "./index.less";
import { useEffect, useState } from "react";
import { List } from "antd";
import Dialog from "../dialog";

interface MenuBannerProps {
  msgList: any[];
}

const Chatting: React.FunctionComponent<MenuBannerProps> = (props) => {
  const { msgList } = props;

  return (
    <div className={"chatting-panel"}>
      {msgList.length > 0 && (
        <List
          itemLayout="horizontal"
          dataSource={msgList}
          renderItem={(item, index) => (
            <List.Item key={index}>
              <Dialog msgObj={item} />
            </List.Item>
          )}
        />
      )}
    </div>
  );
};

export default Chatting;
