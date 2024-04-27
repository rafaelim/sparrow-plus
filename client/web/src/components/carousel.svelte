<script lang="ts">
	import type { CarouselOptions } from '$lib/carousel';
	import {
		createVirtualizer,
		elementScroll,
		type VirtualizerOptions
	} from '@tanstack/svelte-virtual';
	import { onMount } from 'svelte';

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
		const key = options.idKey ?? '';
		const rowId = row[key];

		let params = '';
		if (options.getParams) {
			params = options.getParams(row);
		}

		return `${nextRoute}/${rowId}?${params}`;
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
	.scroll-container {
		height: 300px;
		width: 100vw;
		overflow: scroll;
		transition: 500ms;
		scroll-behavior: smooth;

		scrollbar-width: none;
		&::-webkit-scrollbar {
			display: none;
		}
	}

	.card {
		height: 100%;
		&:first-child {
			.box {
				margin-left: 0;
			}
		}
		&:last-child {
			.box {
				margin-right: 0;
			}
		}
	}

	.box {
		position: relative;
		cursor: pointer;

		display: block;
		width: 250px;
		height: 300px;
		margin: 8px;
		cursor: pointer;
	}

	.box img {
		width: 100%;
		height: 100%;
		object-fit: cover;
	}

	.overlay {
		position: absolute;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		background: rgba(0, 0, 0, 0.7);
		padding: 20px;
		transition: all 0.5s ease;
		opacity: 0;
		pointer-events: none;
	}

	.box:hover .overlay {
		opacity: 1;
		pointer-events: all;
	}

	.overlay .title {
		color: #e9e9e9;
		font-size: 16px;
		text-transform: capitalize;
		font-family: sans-serif;
		cursor: pointer;
	}

	.overlay .title:hover {
		color: #fff;
	}
</style>
