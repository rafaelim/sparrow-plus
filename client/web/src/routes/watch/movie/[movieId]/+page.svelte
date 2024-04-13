<script lang="ts">
	import 'video.js/dist/video-js.css';
	import Video from '@components/video.svelte';

	export let data: { movieId: string };
	const baseUrl = 'http://localhost:3000/api/stream/movie';

	export let options = {
		source: `${baseUrl}/${data.movieId}/master.m3u8`
	};
</script>

<Video {options} />
