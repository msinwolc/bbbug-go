import "./App.css";
import { Suspense, useEffect, useRef } from "react";
import * as React from "react";
import RootComponent from "./router/routerComponent";
import { BrowserRouter } from "react-router-dom";
import Provider from "react-redux/es/components/Provider";
import store, { RootState } from "./store/store";

function App() {
  const mainPageRef = useRef(null);
  // const dispatch = useDispatch();

  return (
    <div className={"mainPage"} ref={mainPageRef}>
      <React.StrictMode>
        <Suspense>
          <Provider store={store}>
            <BrowserRouter>
              <RootComponent />
            </BrowserRouter>
          </Provider>
        </Suspense>
      </React.StrictMode>
    </div>
  );
}

export default App;
