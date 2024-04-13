<script lang="ts">
	import 'video.js/dist/video-js.css';
	import videoJS from 'video.js';
	import { apiData, movies } from '$lib/moviesStore';
	import { onMount } from 'svelte';

	onMount(async () => {
		fetch('http://localhost:3000/api/movies')
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

	let scrollbox: Element;
	$: if (scrollbox) {
		const player = videoJS(scrollbox, {
			controls: true,
			autoplay: true,
			aspectRatio: '16:9',
			html5: {
				nativeTextTracks: false,
				nativeAudioTracks: false,
				hls: {
					overrideNative: true
				}
			},
			sources: [
				{
					src: 'http://localhost:3000/api/stream/<filename>/master',
					type: 'application/x-mpegURL'
				}
			]
		});
	}

	const watchMovie = (movieId: string) => {
		console.log(movieId);
	};
</script>

<main>
	<h1>Movies</h1>
	<div style="display: flex; flex-direction: column">
		{#if movies}
			{#each $movies as movie}
				<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-static-element-interactions -->
				<div on:click={() => watchMovie(movie.movieId)}>{movie.name}</div>
			{/each}
		{/if}
	</div>
</main>
<!-- svelte-ignore a11y-media-has-caption -->
<video bind:this={scrollbox} controls={true} class="vjs-matrix video-js" />
<div id="audioTrackControl"></div>
