import * as React from "react";
import "./index.less";
import { useEffect, useState } from "react";

interface MenuBannerProps {}

const Login: React.FunctionComponent<MenuBannerProps> = (props) => {
  const {} = props;
  
  useEffect(() => {
  }, []);

  return (
    <div className={"login-panel"}>
      login
    </div>
  );
};

export default Login;
