import useFetch from "../../hooks/useFetch";
import Carousel from "../Carousel";

type WatchStatus = {
  watchStatusId: string;
  relationId: string;
  relationType: string;
};

const WatchNext: React.FC = () => {
  const { data, error } = useFetch<WatchStatus[]>(
    "http://192.168.3.16:3000/api/watchStatus"
  );

  if (error || !data?.length) return <></>;
  return (
    <Carousel
      label="Watch Next"
      items={data ?? []}
      renderName={(item: WatchStatus) => item.watchStatusId}
      onItemClick={() => console.log("dasdsadas")}
    />
  );
};

export default WatchNext;
