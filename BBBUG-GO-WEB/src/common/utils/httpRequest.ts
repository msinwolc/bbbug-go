/**
 * http请求封装
 */

import axios, { AxiosRequestConfig } from "axios";

type ajaxProps = AxiosRequestConfig;

//默认配置
const instance = axios.create({
  // baseURL: "http://127.0.0.1:4000/api",
  baseURL: "",
  timeout: 1000,
  // headers: {'X-Custom-Header': 'foobar'}
});

instance.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem("token");
    if (token) {
      config.headers.Authorization = `Bearer ${token}`; // 将 token 添加到请求头中
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export const ajax = (config: ajaxProps) => {
  return new Promise((resolve, reject) => {
    instance({ ...config })
      .then((res) => {
        const { data } = res;
        switch (data?.code) {
          case 0:
            resolve(data?.data || {});
            break;

          default:
            reject(data.msg);
            break;
        }
      })
      .catch((err) => {
        reject(err);
      });
  });
};

export function get(url: string, params?:any, config?: ajaxProps) {
  return ajax({
    url: url,
    method: "get",
    params: params,
    responseType: config?.responseType || "json",
  });
}

export function post(url: string, data: any, config?: ajaxProps) {
  return ajax({
    url: url,
    method: "post",
    data: data,
    responseType: config?.responseType || "json",
  });
}
