<script lang="ts">
	import Carousel from '../components/carousel.svelte';
	import { apiData, movies } from '$lib/moviesStore';
	import { onMount } from 'svelte';
	import type { CarouselOptions } from '$lib/carousel';

	onMount(async () => {
		fetch('http://192.168.3.16:3000/api/movies')
			.then((response) => response.json())
			.then((data) => {
				apiData.set(data);
			})
			.catch((error) => {
				console.log(error);
				return [];
			});
	});

	export const options: CarouselOptions = {
		idKey: 'movieId',
		nextRoute: '/watch/movie',
		titleKey: 'name'
	};
</script>

{#if $movies.length}
	<h1>Movies</h1>
	<div style="display: flex; flex-direction: column">
		<Carousel {options} rows={$movies} />
	</div>
{/if}
