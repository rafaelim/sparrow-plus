<script lang="ts">
	import 'video.js/dist/video-js.css';
	import Video from '@components/video.svelte';

	export let data: { episodeId: string };
	const baseUrl = 'http://192.168.3.16:3000/api';

	export let options = {
		source: `${baseUrl}/stream/master.m3u8?watch=episode&id=${data.episodeId}`
	};

	const updateWatchStatus = (options: { timestamp: number }) => {
		const watchStatus = {
			timestamp: options.timestamp,
			relationId: data.episodeId,
			relationType: 'EPISODE'
		};
		fetch(`${baseUrl}/watchStatus`, {
			method: 'POST',
			headers: [['Content-Type', 'application/json']],
			body: JSON.stringify(watchStatus)
		});
	};
</script>

<Video {options} {updateWatchStatus} />
