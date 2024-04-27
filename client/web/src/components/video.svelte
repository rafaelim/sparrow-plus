<script lang="ts">
	import { onDestroy } from 'svelte';
	import videoJS from 'video.js';
	import type Player from 'video.js/dist/types/player';

	type VideoOptions = {
		source: string;
	};

	export let updateWatchStatus: (options: { timestamp: number }) => void;
	export let options: VideoOptions;

	let scrollbox: Element;
	let intervalId: number;
	$: if (scrollbox) {
		const player = videoJS(scrollbox, {
			controls: true,
			autoplay: true,
			fill: true,
			html5: {
				nativeTextTracks: false,
				nativeAudioTracks: false,
				hls: {
					overrideNative: true
				}
			},
			sources: [
				{
					src: options.source,
					type: 'application/x-mpegURL'
				}
			]
		});
		const urlParams = new URLSearchParams(window.location.search);

		const timestamp = urlParams.get('timestamp');
		player.on('play', () => {
			player.currentTime(timestamp ?? 0);
			intervalId = setInterval(() => {
				update(player);
			}, 5000);
		});
		player.on('pause', () => {
			clearInterval(intervalId);
			update(player);
		});
	}

	const update = (player: Player) => {
		const timestamp = Math.ceil(player.currentTime() ?? 0);
		updateWatchStatus({ timestamp });
	};

	onDestroy(() => {
		clearInterval(intervalId);
		if (scrollbox) videoJS(scrollbox).dispose();
	});
</script>

<div class="player">
	<!-- svelte-ignore a11y-media-has-caption -->
	<video bind:this={scrollbox} controls={true} class="vjs-matrix video-js" />
	<div id="audioTrackControl"></div>
</div>

<style lang="scss">
	.player {
		width: 100vw;
		height: 100vh;
	}
</style>
