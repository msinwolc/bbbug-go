import * as React from "react";
import "./index.less";
import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import Icon, { MessageOutlined, ProfileOutlined, SettingOutlined, StepBackwardOutlined } from "@ant-design/icons";

interface MenuBannerProps {}

const MenuBanner: React.FunctionComponent<MenuBannerProps> = (props) => {
  const {} = props;
  const [menuList, setMenuList] = useState<
    { menuName: string; menuIconUrl: string; icon?: any }[]
  >([]);
  useEffect(() => {
    // 根据用户权限获取菜单列表
    setMenuList([
      {
        menuName: "歌单",
        menuIconUrl: "/page/musicList",
        icon: <StepBackwardOutlined />,
      },
      {
        menuName: "已点",
        menuIconUrl: "/page/oderList",
        icon: <ProfileOutlined />,
      },
      {
        menuName: "房间",
        menuIconUrl: "/page",
        icon: <MessageOutlined />,
      },
      {
        menuName: "设置",
        menuIconUrl: "/page/setting",
        icon: <SettingOutlined />,
      },
    ]);
  }, []);

  return (
    <div className={"menu-bar"}>
      {/* 左侧菜单 */}
      <ul>
        {menuList.map((menu, index) => {
          return (
            <li key={index}>
              <Link to={menu.menuIconUrl}>
                <Icon component={()=>menu.icon} />
                <p>{menu.menuName}</p>
              </Link>
            </li>
          );
        })}
      </ul>
    </div>
  );
};

export default MenuBanner;
