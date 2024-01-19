import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { resolve } from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: [
    { 
      find: "@components", 
      replacement: resolve(__dirname, "./src/components")
    },
    { 
      find: "@hooks", 
      replacement: resolve(__dirname, "./src/hooks")
    },
    { 
      find: "@pages", 
      replacement: resolve(__dirname, "./src/pages")
    },
    { 
      find: "@api", 
      replacement: resolve(__dirname, "./src/api")
    },
    { 
      find: "@types", 
      replacement: resolve(__dirname, "./src/common/types")
    },
    { 
      find: "@utils", 
      replacement: resolve(__dirname, "./src/common/utils")
    },
    { 
      find: "@src", 
      replacement: resolve(__dirname, "./src")
    },
  ],
  },

server: {
  host: "127.0.0.1",
  cors: true,
  port: 4000,
  open: false, //自动打开
  proxy: {
    // 这里的ccc可乱写, 是拼接在url后面的地址 如果接口中没有统一的后缀可自定义
    // 如果有统一后缀, 如api, 直接写api即可, 也不用rewrite了
    "^/api": {
      target: "http://127.0.0.1:8080", // 真实接口地址, 后端给的基地址
      changeOrigin: true, // 允许跨域
      secure: false,  //关键参数，不懂的自己去查
      rewrite: (path) => path.replace(/^\/api/, '')
    },
  },
},
});
