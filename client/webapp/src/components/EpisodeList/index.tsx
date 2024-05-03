import { useNavigate, useParams } from "react-router-dom";
import useFetch from "../../hooks/useFetch";
import Carousel from "../Carousel";
import { NavigationContext } from "../../navigation/SamsungNavigation";
import { useContext } from "react";

type Episode = {
  episodeId: string;
  episodeNumber: string;
  name: string;
  season: string;
  path: string;
};

const EpisodeList: React.FC = () => {
  const { position, positionHandler } = useContext(NavigationContext);
  const { showId } = useParams<{ showId: string }>();
  const navigate = useNavigate();
  const { data, error } = useFetch<Episode[]>(
    `http://192.168.3.16:3000/api/shows/${showId}/episodes`
  );

  if (error || !data?.length) return <></>;

  const rowPosition = positionHandler.getNextY("episodes");
  const getName = (item: Episode) => {
    return `${item.season}x${item.episodeNumber} - ${item.name}`;
  };
  return (
    <Carousel
      label="Episodes"
      items={data ?? []}
      isPositionActive={(colIdx) =>
        colIdx === position.x && position.y === rowPosition
      }
      renderName={getName}
      onItemClick={(row) => navigate(`/watch/episode/${row.episodeId}`)}
    />
  );
};

export default EpisodeList;
