<script lang="ts">
	import { onMount } from 'svelte';
	import type { CarouselOptions } from '$lib/carousel';
	import Carousel from '@components/carousel.svelte';
	import { apiData, episodes } from '$lib/episodeStore';

	export let data: { showId: string };
	onMount(async () => {
		fetch(`http://192.168.3.16:3000/api/episodes/show/${data.showId}`)
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

	export const options: CarouselOptions = {
		idKey: 'episodeId',
		nextRoute: '/watch/episode',
		titleKey: 'name'
	};
</script>

{#if $episodes.length}
	<h1>Episodes</h1>
	<div style="display: flex; flex-direction: column">
		<Carousel {options} rows={$episodes} />
	</div>
{/if}
