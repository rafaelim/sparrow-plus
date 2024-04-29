import { RouterProvider, createHashRouter } from "react-router-dom";
import Home from "./pages/Home";
import { lazy } from "react";

const Episodes = lazy(() => import("./pages/Episodes"));
const EpisodePlayer = lazy(() => import("./pages/EpisodePlayer"));
const MoviePlayer = lazy(() => import("./pages/MoviePlayer"));

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
    {
      path: "/watch/episode/:episodeId",
      element: <EpisodePlayer />,
    },
    {
      path: "/watch/movie/:movieId",
      element: <MoviePlayer />,
    },
  ],
  { basename: "/" }
);

function App() {
  return <RouterProvider router={router} />;
}

export default App;
