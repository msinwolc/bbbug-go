import * as React from "react";
import "./index.less";
import { useEffect, useState } from "react";
import { Button, Checkbox, Form, Input, Space } from "antd";
import { useNavigate } from "react-router-dom";

interface MenuBannerProps {}
interface FieldType {
  username?: string;
  password?: string;
  remember?: string;
}

const Login: React.FunctionComponent<MenuBannerProps> = (props) => {
  const {} = props;
  const navigate = useNavigate();

  useEffect(() => {}, []);

  const onFinish = (values: any) => {
    console.log("Success:", values);
  };
  const onFinishFailed = (errorInfo: any) => {
    console.log("Failed:", errorInfo);
  };
  const onLogin = () => {
    localStorage.setItem("isLogin", "true");
    navigate("/");
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
            name="username"
            rules={[{ required: true, message: "请输入账号!" }]}
          >
            <Input />
          </Form.Item>

          <Form.Item<FieldType>
            label="密码"
            name="password"
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
            <Button type="primary" htmlType="submit" onClick={onLogin}>
              登录
            </Button>
          </Form.Item>
        </Form>
      </Space>
    </div>
  );
};

export default Login;
