import { useContext } from "react";
import useFetch from "../../hooks/useFetch";
import { NavigationContext } from "../../navigation/SamsungNavigation";
import Carousel from "../Carousel";

type Movies = {
  movieId: string;
  name: string;
  year: string;
  path: string;
};

function MoviesList() {
  const { position, positionHandler } = useContext(NavigationContext);
  const { data, error } = useFetch<Movies[]>(
    "http://192.168.3.16:3000/api/movies"
  );

  if (error || !data?.length) return <></>;

  const rowPosition = positionHandler.getNextY("movies");

  return (
    <Carousel
      label="Movies"
      items={data ?? []}
      isPositionActive={(colIdx) =>
        colIdx === position.x && position.y === rowPosition
      }
      renderName={(item: Movies) => item.name}
      onItemClick={() => console.log("dasdsadas")}
    />
  );
}

export default MoviesList;
