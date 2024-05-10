//路由配置
import { lazy } from "react";
const Login = lazy(() => import("../pages/login"));
const MainPage = lazy(() => import("../pages/mainPage"));
const MusicListPage = lazy(() => import("../pages/musicListPage"));
const OderListPage = lazy(() => import("../pages/oderListPage"));
const RoomPage = lazy(() => import("../pages/roomPage"));
const SettingPage = lazy(() => import("../pages/settingPage"));

const router = [
  { path: "/login", element: <Login /> },
  {
    path: "/page",
    element: <MainPage />,
    children: [
      { path: "/page/musicList", element: <MusicListPage />},
      { path: "/page/oderList", element: <OderListPage />},
      { path: "", element: <RoomPage />, index: true },
      { path: "/page/setting", element: <SettingPage />},
    ],
  },
];

export default router;
