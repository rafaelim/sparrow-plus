import { useVirtualizer } from "@tanstack/react-virtual";
import placeholder from "../../assets/png/placeholder.png";
import { useRef } from "react";
import {
  Card,
  Container,
  ListContainer,
  Item,
  ItemTitle,
  Title,
} from "./styled";

type CarouselProps<T> = {
  label: string;
  items: T[];
  renderName: (item: T) => string;
  onItemClick: (item: T) => void;
  isPositionActive: (colIdx: number) => boolean;
};

const Carousel = <T extends Record<string, unknown>>({
  label,
  items,
  renderName,
  onItemClick,
  isPositionActive,
}: CarouselProps<T>) => {
  const parentRef = useRef<HTMLDivElement | null>(null);

  const virtualizer = useVirtualizer({
    horizontal: true,
    count: items.length,
    getScrollElement: () => parentRef.current,
    estimateSize: () => 45,
  });

  return (
    <Container>
      <Title>{label}</Title>
      <ListContainer ref={parentRef} style={{ height: 320, overflowY: "auto" }}>
        <div
          style={{
            width: virtualizer.getTotalSize(),
            height: "100%",
            position: "relative",
          }}
        >
          {virtualizer.getVirtualItems().map((virtualColumn) => (
            <Card
              key={virtualColumn.key}
              data-index={virtualColumn.index}
              ref={virtualizer.measureElement}
              style={{
                position: "absolute",
                top: 0,
                left: 0,
                height: "100%",
                transform: `translateX(${virtualColumn.start}px)`,
              }}
            >
              <Item
                onClick={() => onItemClick(items[virtualColumn.index])}
                active={isPositionActive(virtualColumn.index)}
              >
                <img src={placeholder} alt="img" />
                <div className="overlay">
                  <ItemTitle>
                    {renderName(items[virtualColumn.index])}
                  </ItemTitle>
                </div>
              </Item>
            </Card>
          ))}
        </div>
      </ListContainer>
    </Container>
  );
};

export default Carousel;
