import { Navigate, Route, Routes } from "react-router-dom";
import { useIsAuth } from "../hooks/useIsAuth";
import router from "./router";
import Login from "../pages/login";

const RootComponent = () => {
  const isAuth = useIsAuth({});
  return isAuth ? (
    <Routes>
      {router.map((route) => {
        return route?.children ? (
          <Route {...route}>
            {route.children.map((children: any) => {
              return <Route {...children} />;
            })}
          </Route>
        ) : (
          <Route {...route} />
        );
      })}
    </Routes>
  ) : (
    <Navigate replace to="/login" />
  );
};

export default RootComponent;
