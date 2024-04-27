<script lang="ts">
	import Carousel from './carousel.svelte';
	import { apiData, watchNext } from '$lib/watchNextStore';
	import { onMount } from 'svelte';
	import type { CarouselOptions } from '$lib/carousel';

	export let rowIndex = 0;

	onMount(async () => {
		fetch('http://192.168.3.16:3000/api/watchStatus')
			.then((response) => response.json())
			.then((data) => {
				console.log(data);
				apiData.set(data);
			})
			.catch((error) => {
				console.log(error);
				return [];
			});
	});

	const getWatchNextURL = (row: Record<string, unknown>) => {
		if (row.relationType === 'EPISODE') return `/watch/episode`;
		if (row.relationType === 'MOVIE') return `/watch/movie`;
		throw new Error(`unknown type => ${row.relationType}`);
	};

	const getParams = (row: Record<string, unknown>) => {
		return `timestamp=${row.timestamp}`;
	};

	export const options: CarouselOptions = {
		idKey: 'relationId',
		nextRoute: getWatchNextURL,
		getParams,
		titleKey: ''
	};
</script>

{#if $watchNext.length}
	<section class="section" data-row-index={rowIndex}>
		<h1>Watch Next</h1>
		<div style="display: flex; flex-direction: column">
			<Carousel {options} rows={$watchNext} />
		</div>
	</section>
{/if}

<style lang="scss">
	.section {
		margin: 10px;
	}
</style>
