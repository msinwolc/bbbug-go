import * as React from "react";
import "./index.less";
import { useEffect, useState } from "react";
import { Button, Checkbox, Form, Input, Space } from "antd";
import { useNavigate } from "react-router-dom";
import { doLogin } from "../../api/global";

interface MenuBannerProps {}
interface FieldType {
  user_account?: string;
  user_password?: string;
  remember?: string;
}

const Login: React.FunctionComponent<MenuBannerProps> = (props) => {
  const {} = props;
  const navigate = useNavigate();

  useEffect(() => {}, []);

  const onFinish = (values: any) => {
    doLogin(values).then((res: any) => {
      localStorage.setItem("access_token", res.access_token);
      localStorage.setItem("isLogin", "true");
      navigate("/");
    });
  };

  const onFinishFailed = (errorInfo: any) => {
    console.log("Failed:", errorInfo);
  };

  return (
    <div className={"login-panel"}>
      <Space direction="vertical" size="large">
        <Form
          name="basic"
          labelCol={{ span: 8 }}
          wrapperCol={{ span: 16 }}
          style={{ maxWidth: 600 }}
          initialValues={{ remember: true }}
          onFinish={onFinish}
          onFinishFailed={onFinishFailed}
          autoComplete="off"
        >
          <Form.Item<FieldType>
            label="账号"
            name="user_account"
            rules={[{ required: true, message: "请输入账号!" }]}
          >
            <Input />
          </Form.Item>

          <Form.Item<FieldType>
            label="密码"
            name="user_password"
            rules={[{ required: true, message: "请输入密码!" }]}
          >
            <Input.Password />
          </Form.Item>

          <Form.Item<FieldType>
            name="remember"
            valuePropName="checked"
            wrapperCol={{ offset: 8, span: 16 }}
          >
            <Checkbox>记住密码</Checkbox>
          </Form.Item>

          <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
            <Button type="primary" htmlType="submit">
              登录
            </Button>
          </Form.Item>
        </Form>
      </Space>
    </div>
  );
};

export default Login;
