import { useCallback, useState } from "react";
import MoviesList from "../../components/MoviesList";
import ShowList from "../../components/ShowList";
import WatchNext from "../../components/WatchNext";

const Home = () => {
  const [sections, setSections] = useState<Set<string>>(new Set());
  const [row, setRow] = useState(0);

  const handleOnLoad = useCallback((section: string) => {
    setSections((oldSections) => {
      const newSet = oldSections.add(section);
      return new Set(newSet);
    });
  }, []);

  const watchNextSection = useCallback(
    () => handleOnLoad("watch-next"),
    [handleOnLoad]
  );
  const moviesSection = useCallback(
    () => handleOnLoad("movies"),
    [handleOnLoad]
  );
  const tvShowsSection = useCallback(
    () => handleOnLoad("tv-shows"),
    [handleOnLoad]
  );

  const getPosition = (section: string) => {
    const list = [...sections.values()];

    return list.indexOf(section);
  };

  return (
    <>
      <WatchNext
        onLoad={watchNextSection}
        active={row === getPosition("watch-next")}
      />
      <MoviesList
        onLoad={moviesSection}
        active={row === getPosition("movies")}
      />
      <ShowList
        onLoad={tvShowsSection}
        active={row === getPosition("tv-shows")}
      />
    </>
  );
};

export default Home;
