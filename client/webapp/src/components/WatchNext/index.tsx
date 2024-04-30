import { createSearchParams, useNavigate } from "react-router-dom";
import useFetch from "../../hooks/useFetch";
import Carousel from "../Carousel";
import { useEffect } from "react";

type WatchStatus = {
  watchStatusId: string;
  relationId: string;
  relationType: string;
  timestamp: string;
};

type WatchNextProps = {
  onLoad: () => void;
  active: boolean;
};

const WatchNext: React.FC<WatchNextProps> = ({ onLoad, active }) => {
  const navigate = useNavigate();
  const { data, error } = useFetch<WatchStatus[]>(
    "http://192.168.3.16:3000/api/watchStatus"
  );

  useEffect(() => {
    if (error || !data?.length) return;
    onLoad();
  }, [error, data, onLoad]);

  if (error || !data?.length) return <></>;

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
      active={active}
      renderName={(item: WatchStatus) => item.watchStatusId}
      onItemClick={handleOnClick}
    />
  );
};

export default WatchNext;
