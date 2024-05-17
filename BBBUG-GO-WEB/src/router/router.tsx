//路由配置
import { lazy } from "react";
import { Navigate } from "react-router-dom";
const MainPage = lazy(() => import("../pages/mainPage"));
const MusicListPage = lazy(() => import("../pages/musicListPage"));
const OderListPage = lazy(() => import("../pages/oderListPage"));
const RoomPage = lazy(() => import("../pages/roomPage"));
const SettingPage = lazy(() => import("../pages/settingPage"));

const router = [
  // { path: "/login", element: <Login /> },
  {
    path: "/page",
    element: <MainPage />,
    children: [
      { path: "/page/musicList", element: <MusicListPage /> },
      { path: "/page/oderList", element: <OderListPage /> },
      { path: "", element: <RoomPage />, index: true },
      { path: "/page/setting", element: <SettingPage /> },
    ],
  },
  { path: "/", element: <Navigate replace to="/page" /> },
];

export default router;
