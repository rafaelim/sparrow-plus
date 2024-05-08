import useTizenClickable from "../../hooks/tizen-navigation/useTizenClickable";
import { Item, ItemTitle } from "./styled";

type CardProps = {
  onClick: () => void;
  active: boolean;
  placeholderImage: string;
  title: string;
};

function Card({ active, onClick, placeholderImage, title }: CardProps) {
  useTizenClickable({ onClick, isActive: active });

  return (
    <Item onClick={onClick} active={active}>
      <img src={placeholderImage} alt="img" />
      <div className="overlay">
        <ItemTitle>{title}</ItemTitle>
      </div>
    </Item>
  );
}

export default Card;
