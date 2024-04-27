import { RouterProvider, createHashRouter } from "react-router-dom";
import Home from "./pages/home";
import { lazy } from "react";

const Episodes = lazy(() => import("./pages/episodes"));

const router = createHashRouter(
  [
    {
      path: "/",
      element: <Home />,
    },
    {
      path: "/shows/:showId/episodes",
      element: <Episodes />,
    },
  ],
  { basename: "/" }
);

function App() {
  return <RouterProvider router={router} />;
}

export default App;
