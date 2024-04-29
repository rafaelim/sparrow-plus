import { useParams } from "react-router-dom";
import PlayerComponent from "../../components/Player";
import { BASE_URL } from "../../utils/constants";

const MoviePlayer = () => {
  const { movieId } = useParams<{ movieId: string }>();

  const updateWatchStatus = (timestamp: number) => {
    const watchStatus = {
      timestamp: timestamp,
      relationId: movieId,
      relationType: "MOVIE",
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
      source={`${BASE_URL}/stream/master.m3u8?watch=movie&id=${movieId}`}
    />
  );
};

export default MoviePlayer;
