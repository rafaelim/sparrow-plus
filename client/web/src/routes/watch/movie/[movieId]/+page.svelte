<script lang="ts">
	import 'video.js/dist/video-js.css';
	import Video from '@components/video.svelte';

	export let data: { movieId: string };
	const baseUrl = 'http://localhost:3000/api/stream/master';

	export let options = {
		source: `${baseUrl}/master.m3u8?type=movie&id=${data.movieId}`
	};
</script>

<Video {options} />
