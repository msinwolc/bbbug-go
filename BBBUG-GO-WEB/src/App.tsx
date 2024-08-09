import "./App.css";
import { Suspense, useEffect, useRef } from "react";
import * as React from "react";
import RootComponent from "./router/routerComponent";
import { BrowserRouter } from "react-router-dom";
import Provider from "react-redux/es/components/Provider";
import store, { RootState } from "./store/store";
import { getWebsocketUrl } from "./api/global";
import { useSelector } from "react-redux";

function App() {
  const mainPageRef = useRef(null);

  return (
    <div className={"mainPage"} ref={mainPageRef}>
      <Suspense>
          <Provider store={store}>
            <BrowserRouter>
              <RootComponent />
            </BrowserRouter>
          </Provider>
        </Suspense>
    </div>
  );
}

export default App;
