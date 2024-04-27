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

	export let rowIndex = 0;

	export const options: CarouselOptions = {
		idKey: 'showId',
		nextRoute: '/shows',
		titleKey: 'name'
	};
</script>

{#if $shows.length}
	<section class="section" data-row-index={rowIndex}>
		<h1>TV Shows</h1>
		<Carousel {options} rows={$shows} />
	</section>
{/if}

<style lang="scss">
	.section {
		margin: 10px;
	}
</style>
