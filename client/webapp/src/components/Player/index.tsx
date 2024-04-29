import { useCallback, useEffect, useRef } from "react";
import videoJS from "video.js";
import type Player from "video.js/dist/types/player";
import styled from "styled-components";

const Container = styled.div`
  width: 100vw;
  height: 100vh;
`;

type PlayerComponentProps = {
  onTimestampUpdates: (timestamp: number) => void;
  source: string;
};

const PlayerComponent: React.FC<PlayerComponentProps> = ({
  onTimestampUpdates,
  source,
}) => {
  const videoRef = useRef<HTMLVideoElement>(null);
  const playerRef = useRef<Player>();

  useEffect(() => {
    if (!videoRef?.current) return;
    playerRef.current = videoJS(videoRef?.current as Element, {
      controls: true,
      autoplay: true,
      fill: true,
      html5: {
        nativeTextTracks: false,
        nativeAudioTracks: false,
        hls: {
          overrideNative: true,
        },
      },
      sources: [
        {
          src: source,
          type: "application/x-mpegURL",
        },
      ],
    });
  }, [source]);

  const update = useCallback(
    (player: Player) => {
      const timestamp = Math.ceil(player.currentTime() ?? 0);
      onTimestampUpdates(timestamp);
    },
    [onTimestampUpdates]
  );

  useEffect(() => {
    if (!playerRef?.current) return;

    let intervalId: number;
    const urlParams = new URLSearchParams(window.location.search);

    const timestamp = urlParams.get("timestamp");
    if (timestamp) {
      playerRef.current?.currentTime(timestamp ?? 0);
    }
    playerRef.current.on("play", () => {
      intervalId = setInterval(() => {
        update(playerRef.current as Player);
      }, 5000);
    });
    playerRef.current.on("pause", () => {
      clearInterval(intervalId);
      update(playerRef.current as Player);
    });
  }, [playerRef, update]);

  return (
    <Container>
      <video
        ref={videoRef}
        className="video-js vjs-default-skin vjs-16-9 vjs-big-play-centered"
      />
      <div id="audioTrackControl"></div>
    </Container>
  );
};

export default PlayerComponent;
