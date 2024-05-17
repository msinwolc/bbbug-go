import { configureStore } from "@reduxjs/toolkit";
import logger from "redux-logger";

import globalReducer from "./global"; // 引入合并后的reducer
import { GlobalState } from "./global";

export type RootState = GlobalState;

const store = configureStore({
  reducer: globalReducer, // 使用你的rootReducer
  middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(logger),
});

export default store;
