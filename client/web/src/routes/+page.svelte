<script lang="ts">
	import 'video.js/dist/video-js.css';
	import Movies from '@components/movies.svelte';
	import Shows from '@components/shows.svelte';
	import WatchNext from '@components/watchNext.svelte';
</script>

<main class="main">
	<Movies rowIndex={0} />
	<Shows rowIndex={1} />
	<WatchNext />
</main>

<style lang="scss">
	.main {
		display: flex;
		flex-direction: column;
		gap: 15px;
	}
</style>