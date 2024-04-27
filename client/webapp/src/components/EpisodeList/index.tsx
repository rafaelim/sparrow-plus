import { useParams } from "react-router-dom";
import useFetch from "../../hooks/useFetch";
import Carousel from "../Carousel";

type Episode = {
  episodeId: string;
  episodeNumber: string;
  name: string;
  season: string;
  path: string;
};

const EpisodeList: React.FC = () => {
  const { showId } = useParams<{ showId: string }>();
  const { data, error } = useFetch<Episode[]>(
    `http://192.168.3.16:3000/api/shows/${showId}/episodes`
  );

  if (error || !data?.length) return <></>;

  const getName = (item: Episode) => {
    return `${item.season}x${item.episodeNumber} - ${item.name}`;
  };
  return (
    <Carousel
      label="Episodes"
      items={data ?? []}
      renderName={getName}
      onItemClick={() => console.log("dasdsadas")}
    />
  );
};

export default EpisodeList;
