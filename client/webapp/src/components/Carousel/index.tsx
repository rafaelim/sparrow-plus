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
};

const Carousel = <T extends Record<string, unknown>>({
  label,
  items,
  renderName,
  onItemClick,
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
              {/* <a
                href={getNextRoute(rows[col.index])}
                class="box"
                style="width: {rows[col.index].length}px" 
              >       */}
              <Item onClick={() => onItemClick(items[virtualColumn.index])}>
                <img src={placeholder} alt="img" />
                <div className="overlay">
                  <ItemTitle>
                    {renderName(items[virtualColumn.index])}
                  </ItemTitle>
                </div>
              </Item>
              {/* </a> */}
            </Card>
          ))}
        </div>
      </ListContainer>
    </Container>
  );
};

export default Carousel;

{
  /* <script lang="ts">
	import type { CarouselOptions } from '$lib/carousel';
	import { createVirtualizer } from '@tanstack/svelte-virtual';

	export let rows: Record<string, string>[] = [];
	export let options: CarouselOptions = {};
	let virtualListEl: HTMLDivElement;
	let virtualItemEls: HTMLDivElement[] = [];

	$: virtualizer = createVirtualizer<HTMLDivElement, HTMLDivElement>({
		horizontal: true,
		count: rows.length,
		getScrollElement: () => virtualListEl,
		estimateSize: () => 10
	});
	const getNextRoute = (row: Record<string, string>) => {
		const nextRoute =
			typeof options.nextRoute === 'function' ? options.nextRoute(row) : options.nextRoute;
		const type = typeof options.type === 'function' ? options.type(row) : options.type;
		const key = options.idKey ?? '';
		const rowId = row[key];

		let params = '';
		if (options.getParams) {
			params = options.getParams(row);
		}

		return `${nextRoute}?${type}=${rowId}?${params}`;
	};

	const getTitle = (row: Record<string, string>) => {
		const key = options.titleKey ?? '';
		return row[key];
	};
	let start;
	let end;
	$: {
		if (virtualItemEls.length) virtualItemEls.forEach((el) => $virtualizer.measureElement(el));
	}
</script>

<div class="scroll-container" {...options} bind:this={virtualListEl}>
	{#if options}
		<div
			class="container"
			style="position: relative; height: 100%; width: {$virtualizer.getTotalSize()}px;"
		>
			{#each $virtualizer.getVirtualItems() as col, idx (col.index)}
				<div
					bind:this={virtualItemEls[idx]}
					data-index={col.index}
					class="card"
					style="position: absolute; left: 0; transform: translateX({col.start}px);"
				>
					<a
						href={getNextRoute(rows[col.index])}
						class="box"
						style="width: {rows[col.index].length}px"
					>
						<img src="/images/placeholder.png" alt="img" />
						<div class="overlay">
							<h3 class="title">{getTitle(rows[col.index])}</h3>
						</div>
					</a>
				</div>
			{/each}
		</div>
	{/if}
</div>

<style lang="scss">
	
</style> */
}
