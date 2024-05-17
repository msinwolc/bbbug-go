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
        replacement: resolve(__dirname, "./src/components"),
      },
      {
        find: "@hooks",
        replacement: resolve(__dirname, "./src/hooks"),
      },
      {
        find: "@pages",
        replacement: resolve(__dirname, "./src/pages"),
      },
      {
        find: "@api",
        replacement: resolve(__dirname, "./src/api"),
      },
      {
        find: "@types",
        replacement: resolve(__dirname, "./src/common/types"),
      },
      {
        find: "@utils",
        replacement: resolve(__dirname, "./src/common/utils"),
      },
      {
        find: "@src",
        replacement: resolve(__dirname, "./src"),
      },
    ],
  },

  server: {
    host: "127.0.0.1",
    cors: true,
    port: 4000,
    open: false, //自动打开
    proxy: {
      "^/api": {
        target: "https://saysth.fun", // 后端基地址
        changeOrigin: true, // 允许跨域
        secure: true, // https需打开为true
        // rewrite: (path) => path.replace(/^\/api/, '')
      },
    },
  },
});
