<script lang="ts">
	import 'video.js/dist/video-js.css';
	import Video from '@components/video.svelte';

	export let data: { movieId: string };
	const baseUrl = 'http://192.168.3.16:3000/api/stream';

	export let options = {
		source: `${baseUrl}/master.m3u8?watch=movies&id=${data.movieId}`
	};
</script>

<Video {options} />
