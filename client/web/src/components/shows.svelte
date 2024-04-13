<script lang="ts">
	import Carousel from './carousel.svelte';
	import { apiData, shows } from '$lib/showsStore';
	import { onMount } from 'svelte';
	import type { CarouselOptions } from '$lib/carousel';

	onMount(async () => {
		fetch('http://192.168.3.16:3000/api/shows')
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
		rows: $shows,
		idKey: 'showId',
		nextRoute: '/watch/show',
		titleKey: 'name'
	};
</script>

{#if $shows.length}
	<h1>Shows</h1>
	<div style="display: flex; flex-direction: column">
		<Carousel {options} />
	</div>
{/if}
