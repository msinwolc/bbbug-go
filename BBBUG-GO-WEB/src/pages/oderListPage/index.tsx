import * as React from "react";
import "./index.less";


interface MainChatProps {}

const MainChat: React.FunctionComponent<MainChatProps> = (props) => {
  const {} = props;

  const data = {
    roomId: 1234,
    roomName: "",
  };

  return (
    <div className={"mainChatPanel"}>
      <>已点</>
    </div>
  );
};

export default MainChat;
