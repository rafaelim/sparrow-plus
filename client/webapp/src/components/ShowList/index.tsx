import { useEffect } from "react";
import useFetch from "../../hooks/useFetch";
import Carousel from "../Carousel";
import { useNavigate } from "react-router-dom";

type TvShow = {
  showId: string;
  name: string;
};

type ShowListProps = {
  onLoad: () => void;
  active: boolean;
};

const ShowList: React.FC<ShowListProps> = ({ onLoad, active }) => {
  const navigate = useNavigate();
  const { data, error } = useFetch<TvShow[]>(
    "http://192.168.3.16:3000/api/shows"
  );

  useEffect(() => {
    if (error || !data?.length) return;
    onLoad();
  }, [error, data, onLoad]);

  if (error || !data?.length) return <></>;

  const onClick = (item: TvShow) => {
    navigate(`/shows/${item.showId}/episodes`);
  };

  return (
    <Carousel
      label="Shows"
      active={active}
      items={data ?? []}
      renderName={(item: TvShow) => item.name}
      onItemClick={onClick}
    />
  );
};

export default ShowList;
