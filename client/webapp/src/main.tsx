import React, { Suspense } from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "video.js/dist/video-js.css";
import GlobalStyle from "./GlobalStyles.tsx";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <Suspense fallback={<div>loading...</div>}>
      <GlobalStyle />
      <App />
    </Suspense>
  </React.StrictMode>
);
