import * as React from "react";
import "./index.less";
import { useEffect, useState } from "react";
import { Icon } from "@chakra-ui/react";
import { CopyIcon } from "@chakra-ui/icons";

interface MenuBannerProps {}

const MenuBanner: React.FunctionComponent<MenuBannerProps> = (props) => {
  const {} = props;
  const [menuList, setMenuList] = useState<
    { menuName: string; menuIconUrl: string }[]
  >([]);
  useEffect(() => {
    // 根据用户权限获取菜单列表
    setMenuList([
      { menuName: "歌单", menuIconUrl: "" },
      { menuName: "已点", menuIconUrl: "" },
      { menuName: "房间", menuIconUrl: "" },
      { menuName: "设置", menuIconUrl: "" },
    ]);
  }, []);

  return (
    <div className={"menu-bar"}>
      {/* 左侧菜单 */}
      <ul>
        {menuList.map((menu, index) => {
          return (
            <li key={index}>
              <Icon as={CopyIcon} />
              <p>{menu.menuName}</p>
            </li>
          );
        })}
      </ul>
    </div>
  );
};

export default MenuBanner;
