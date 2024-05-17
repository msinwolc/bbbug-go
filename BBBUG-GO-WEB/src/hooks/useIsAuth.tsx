interface useIsAuthProps {}

export const useIsAuth = (props: useIsAuthProps) => {
  const {} = props;
  const isLogin = localStorage.getItem("isLogin");
  return isLogin;
};
