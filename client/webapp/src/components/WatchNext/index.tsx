import { createSearchParams, useNavigate } from "react-router-dom";
import useFetch from "../../hooks/useFetch";
import Carousel from "../Carousel";
import { NavigationContext } from "../../navigation/SamsungNavigation";
import { useContext } from "react";

type WatchStatus = {
  watchStatusId: string;
  relationId: string;
  relationType: string;
  timestamp: string;
};

function WatchNext() {
  const navigate = useNavigate();
  const { position, positionHandler } = useContext(NavigationContext);
  const { data, error } = useFetch<WatchStatus[]>(
    "http://192.168.3.16:3000/api/watchStatus"
  );

  if (error || !data?.length) return <></>;
  const rowPosition = positionHandler.getNextY("watch-next");

  const handleOnClick = (row: WatchStatus) => {
    if (row.relationType === "EPISODE") {
      navigate({
        pathname: `/watch/episode/${row.relationId}`,
        search: createSearchParams({
          timestamp: row.timestamp,
        }).toString(),
      });
    }
    if (row.relationType === "MOVIE") {
      navigate({
        pathname: `/watch/movieId/${row.relationId}`,
        search: createSearchParams({ timestamp: row.timestamp }).toString(),
      });
    }
  };

  return (
    <Carousel
      label="Watch Next"
      items={data ?? []}
      isPositionActive={(colIdx) =>
        colIdx === position.x && position.y === rowPosition
      }
      renderName={(item: WatchStatus) => item.watchStatusId}
      onItemClick={handleOnClick}
    />
  );
}

export default WatchNext;
