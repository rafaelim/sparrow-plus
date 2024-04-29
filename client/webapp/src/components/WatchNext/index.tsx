import { useNavigate } from "react-router-dom";
import useFetch from "../../hooks/useFetch";
import Carousel from "../Carousel";

type WatchStatus = {
  watchStatusId: string;
  relationId: string;
  relationType: string;
  timestamp: string;
};

const WatchNext: React.FC = () => {
  const navigate = useNavigate();
  const { data, error } = useFetch<WatchStatus[]>(
    "http://192.168.3.16:3000/api/watchStatus"
  );

  if (error || !data?.length) return <></>;

  const handleOnClick = (row: WatchStatus) => {
    if (row.relationType === "EPISODE") {
      navigate(`/watch/episode/${row.relationId}?timestamp=${row.timestamp}`);
    }
    if (row.relationType === "MOVIE") {
      navigate(`/watch/movieId/${row.relationId}?timestamp=${row.timestamp}`);
    }
  };

  return (
    <Carousel
      label="Watch Next"
      items={data ?? []}
      renderName={(item: WatchStatus) => item.watchStatusId}
      onItemClick={handleOnClick}
    />
  );
};

export default WatchNext;
