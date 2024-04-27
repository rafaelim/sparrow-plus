import useFetch from "../../hooks/useFetch";
import Carousel from "../Carousel";

type Movies = {
  movieId: string;
  name: string;
  year: string;
  path: string;
};

const MoviesList: React.FC = () => {
  const { data, error } = useFetch<Movies[]>(
    "http://192.168.3.16:3000/api/movies"
  );

  if (error || !data?.length) return <></>;
  return (
    <Carousel
      label="Movies"
      items={data ?? []}
      renderName={(item: Movies) => item.name}
      onItemClick={() => console.log("dasdsadas")}
    />
  );
};

export default MoviesList;
