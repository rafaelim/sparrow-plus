import { useEffect } from "react";
import useFetch from "../../hooks/useFetch";
import Carousel from "../Carousel";

type Movies = {
  movieId: string;
  name: string;
  year: string;
  path: string;
};

type MoviesListProps = {
  onLoad: () => void;
  active: boolean;
};

const MoviesList: React.FC<MoviesListProps> = ({ onLoad, active }) => {
  const { data, error } = useFetch<Movies[]>(
    "http://192.168.3.16:3000/api/movies"
  );

  useEffect(() => {
    if (error || !data?.length) return;
    onLoad();
  }, [error, data, onLoad]);

  if (error || !data?.length) return <></>;
  return (
    <Carousel
      label="Movies"
      items={data ?? []}
      active={active}
      renderName={(item: Movies) => item.name}
      onItemClick={() => console.log("dasdsadas")}
    />
  );
};

export default MoviesList;
