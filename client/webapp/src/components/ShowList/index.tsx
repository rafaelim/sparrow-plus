import { useContext } from "react";
import useFetch from "../../hooks/useFetch";
import Carousel from "../Carousel";
import { useNavigate } from "react-router-dom";
import { NavigationContext } from "../../navigation/SamsungNavigation";

type TvShow = {
  showId: string;
  name: string;
};

function ShowList() {
  const navigate = useNavigate();
  const { position, positionHandler } = useContext(NavigationContext);
  const { data, error } = useFetch<TvShow[]>(
    "http://192.168.3.16:3000/api/shows"
  );

  if (error || !data?.length) return <></>;

  const rowPosition = positionHandler.getNextY("tv-shows");

  const onClick = (item: TvShow) => {
    navigate(`/shows/${item.showId}/episodes`);
  };

  return (
    <Carousel
      label="Shows"
      items={data ?? []}
      isPositionActive={(colIdx) =>
        colIdx === position.x && position.y === rowPosition
      }
      renderName={(item: TvShow) => item.name}
      onItemClick={onClick}
    />
  );
}

export default ShowList;
