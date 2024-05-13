import "./App.css";
import { Suspense, useEffect, useRef } from "react";
import router from "./router/router";
import * as React from "react";
import RootComponent from "./router/routerComponent";
import { BrowserRouter } from "react-router-dom";

function App() {
  const mainPageRef = useRef(null);
  useEffect(() => {
    console.log(router);
  }, [router]);

  return (
    <div className={"mainPage"} ref={mainPageRef}>
      <React.StrictMode>
        <Suspense>
          <BrowserRouter>
            <RootComponent />
          </BrowserRouter>
        </Suspense>
      </React.StrictMode>
    </div>
  );
}

export default App;
