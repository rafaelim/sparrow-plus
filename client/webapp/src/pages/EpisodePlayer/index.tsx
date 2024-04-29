import { useParams } from "react-router-dom";
import PlayerComponent from "../../components/Player";
import { BASE_URL } from "../../utils/constants";

const EpisodePlayer = () => {
  const { episodeId } = useParams<{ episodeId: string }>();

  const updateWatchStatus = (timestamp: number) => {
    const watchStatus = {
      timestamp: timestamp,
      relationId: episodeId,
      relationType: "EPISODE",
    };
    fetch(`${BASE_URL}/watchStatus`, {
      method: "POST",
      headers: [["Content-Type", "application/json"]],
      body: JSON.stringify(watchStatus),
    });
  };

  return (
    <PlayerComponent
      onTimestampUpdates={updateWatchStatus}
      source={`${BASE_URL}/stream/master.m3u8?watch=episode&id=${episodeId}`}
    />
  );
};

export default EpisodePlayer;
