<script lang="ts">
	import 'video.js/dist/video-js.css';
	import Video from '@components/video.svelte';

	export let data: { showId: string };
	const baseUrl = 'http://192.168.3.16:3000/api/stream/show';

	export let options = {
		source: `${baseUrl}/${data.showId}/master.m3u8`
	};
</script>

<Video {options} />
