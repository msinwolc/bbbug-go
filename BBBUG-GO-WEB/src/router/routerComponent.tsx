import { Navigate, Route, Routes, useLocation } from "react-router-dom";
import { useIsAuth } from "../hooks/useIsAuth";
import router from "./router";
import Login from "../pages/login";
import { JSX } from "react/jsx-runtime";

const RootComponent = () => {
  const isAuth = useIsAuth({});

  const withAuth = (Component: any) => {
    const location = useLocation();
    if (!isAuth) {
      return <Navigate replace to="/login" state={{ from: location }} />;
    }

    return Component;
  };

  return (
    <Routes>
      {router.map((route) => {
        return route?.children ? (
          <Route path={route.path} element={withAuth(route.element)}>
            {route.children.map((children: any) => {
              return <Route {...children} />;
            })}
          </Route>
        ) : (
          <Route path={route.path} element={withAuth(route.element)} />
        );
      })}
      <Route path={"/login"} element={<Login />} />
    </Routes>
  );
};

export default RootComponent;
