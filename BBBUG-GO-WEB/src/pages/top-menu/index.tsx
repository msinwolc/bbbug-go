import * as React from "react";
import { IconButton, Input } from "@chakra-ui/react";
import { SearchIcon } from "@chakra-ui/icons";
import "./index.less";
import { post, get } from "../../common/utils/httpRequest";

interface TopMenuBarProps {
  menuData: string[];
  isFix: boolean;
}

const TopMenuBar: React.FunctionComponent<TopMenuBarProps> = (props) => {
  const { menuData, isFix = false } = props;
  return (
    <div className={isFix ? "top-menu-bar top-menu-bar-fix" : "top-menu-bar"}>
      {/* 左侧菜单 */}
      <ul className={"menu-left"}>
        {menuData.map((menu, index) => {
          return (
            <li
              key={index}
              onClick={() => {
                post("api/song/songList", {
                  Rid: "",
                  Plat: "",
                  Version: "",
                  AccessToken: "",
                }).then((res) => {
                  console.log(res);
                });
              }}
            >
              {menu}
            </li>
          );
        })}
      </ul>
      {/* 中间搜索框 */}
      <div className={"menu-search"}>
        <Input />
        <IconButton aria-label="Search database" icon={<SearchIcon />} />
      </div>
      {/* 右侧菜单 */}
      <ul className="menu-right">
        {menuData.map((menu, index) => {
          return <li key={index}>{menu}</li>;
        })}
      </ul>
    </div>
  );
};

export default TopMenuBar;
