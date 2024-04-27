<script lang="ts">
	import 'video.js/dist/video-js.css';
	import Video from '@components/video.svelte';

	export let data: { movieId: string };
	const baseUrl = 'http://192.168.3.16:3000/api/stream';

	export let options = {
		source: `${baseUrl}/master.m3u8?watch=movies&id=${data.movieId}`
	};

	const updateWatchStatus = (options: { timestamp: number }) => {
		const watchStatus = {
			timestamp: options.timestamp,
			relationId: data.movieId,
			relationType: 'MOVIE'
		};
		fetch(`${baseUrl}/watchStatus`, {
			method: 'POST',
			headers: [['Content-Type', 'application/json']],
			body: JSON.stringify(watchStatus)
		});
	};
</script>

<Video {options} {updateWatchStatus} />
