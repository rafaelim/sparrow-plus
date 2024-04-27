<script lang="ts">
	import 'video.js/dist/video-js.css';
	import Movies from '@components/movies.svelte';
	import Shows from '@components/shows.svelte';
	import WatchNext from '@components/watchNext.svelte';
</script>

<main class="main">
	<WatchNext rowIndex={0} />
	<Shows rowIndex={1} />
	<Movies rowIndex={2} />
</main>

<style lang="scss">
	.main {
		display: flex;
		flex-direction: column;
		gap: 15px;
	}
</style>
