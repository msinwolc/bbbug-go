//路由配置
import { createBrowserRouter } from "react-router-dom";
import { lazy } from "react";
const Login = lazy(() => import("../pages/login"));
const MainChat = lazy(() => import("../pages/mainChat"));
const MainPage = lazy(() => import("../pages/mainPage"));

const router = [
  { path: "/login", element: <Login /> },
  {
    path: "/page",
    element: <MainPage />,
    children: [
      { path: "", element: <MainChat />, index: true },
      { path: "", element: <MainChat />, index: true },
    ],
  },
];

export default router;
