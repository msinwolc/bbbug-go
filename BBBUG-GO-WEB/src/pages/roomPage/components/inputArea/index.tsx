import * as React from "react";
import "./index.less";
import { useEffect, useState } from "react";
import TextArea from "antd/es/input/TextArea";
import { RootState } from "../../../../store/store";
import { useSelector } from "react-redux";
import { sendMsg } from "../../../../api/global";
import { sendSocketMsg } from "../../../../common/utils/tools";

interface MenuBannerProps {}

const InputArea: React.FunctionComponent<MenuBannerProps> = (props) => {
  const {} = props;
  const ref = React.useRef(null);
  const [value, setValue] = useState("");
  const userMsg = useSelector((state: RootState) => state.userMsg);
  let socket = useSelector((state: RootState) => state.socket);

  const handlePressEnter = (e: any) => {
    console.log(e);
  };

  const handleKeyDown = (e: any) => {
    if (e.keyCode === 13) {
      e.preventDefault();
      sendSocketMsg({ type: "text", msg: value, roomId: 888 });
      setValue("");
    }
    if ((e.altKey || e.metaKey) && e.key === "Enter") {
      // 阻止默认行为，可能包括表单提交等
      e.preventDefault();

      // 在当前内容后添加换行符
      setValue(value + "\n");

      // 保持或重置光标位置
      // maintainCursorAfterAction();
    }
  };

  const maintainCursorAfterAction = () => {
    // 获取当前可编辑元素
    const editElement = ref.current as any;

    // 清除现有选区
    const selection = window.getSelection();
    if (selection) {
      selection.removeAllRanges();
    }

    // 创建新的Range对象，定位到内容的末尾
    const range = document.createRange();
    range.selectNodeContents(editElement);
    range.collapse(false); // 收缩到范围的末尾

    // 设置新的选区
    if (selection) {
      selection.addRange(range);
    }
  };

  return (
    <div className={"input-area-panel"}>
      <TextArea
        ref={ref}
        rows={4}
        value={value}
        placeholder="maxLength is 6"
        maxLength={300}
        onChange={(e: any) => setValue(e.target.value)}
        onPressEnter={handlePressEnter}
        onKeyDown={handleKeyDown}
      />
    </div>
  );
};

export default InputArea;
